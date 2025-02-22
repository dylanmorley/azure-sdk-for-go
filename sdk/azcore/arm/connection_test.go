//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package arm

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/pipeline"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/log"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/mock"
)

type mockTokenCred struct{}

func (mockTokenCred) NewAuthenticationPolicy(azruntime.AuthenticationOptions) policy.Policy {
	return pipeline.PolicyFunc(func(req *policy.Request) (*http.Response, error) {
		return req.Next()
	})
}

func (mockTokenCred) GetToken(context.Context, policy.TokenRequestOptions) (*azcore.AccessToken, error) {
	return &azcore.AccessToken{
		Token:     "abc123",
		ExpiresOn: time.Now().Add(1 * time.Hour),
	}, nil
}

const rpUnregisteredResp = `{
	"error":{
		"code":"MissingSubscriptionRegistration",
		"message":"The subscription registration is in 'Unregistered' state. The subscription must be registered to use namespace 'Microsoft.Storage'. See https://aka.ms/rps-not-found for how to register subscriptions.",
		"details":[{
				"code":"MissingSubscriptionRegistration",
				"target":"Microsoft.Storage",
				"message":"The subscription registration is in 'Unregistered' state. The subscription must be registered to use namespace 'Microsoft.Storage'. See https://aka.ms/rps-not-found for how to register subscriptions."
			}
		]
	}
}`

func TestNewDefaultConnection(t *testing.T) {
	opt := ConnectionOptions{}
	con := NewDefaultConnection(mockTokenCred{}, &opt)
	if ep := con.Endpoint(); ep != AzurePublicCloud {
		t.Fatalf("unexpected endpoint %s", ep)
	}
}

func TestNewConnection(t *testing.T) {
	const customEndpoint = "https://contoso.com/fake/endpoint"
	con := NewConnection(customEndpoint, mockTokenCred{}, nil)
	if ep := con.Endpoint(); ep != customEndpoint {
		t.Fatalf("unexpected endpoint %s", ep)
	}
}

func TestNewConnectionWithOptions(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.AppendResponse()
	opt := ConnectionOptions{}
	opt.HTTPClient = srv
	con := NewConnection(Endpoint(srv.URL()), mockTokenCred{}, &opt)
	if ep := con.Endpoint(); ep != Endpoint(srv.URL()) {
		t.Fatalf("unexpected endpoint %s", ep)
	}
	req, err := azruntime.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	resp, err := con.NewPipeline("armtest", "v1.2.3").Do(req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
	if ua := resp.Request.Header.Get("User-Agent"); !strings.HasPrefix(ua, "azsdk-go-armtest/v1.2.3") {
		t.Fatalf("unexpected User-Agent %s", ua)
	}
}

func TestNewConnectionWithCustomTelemetry(t *testing.T) {
	const myTelemetry = "something"
	srv, close := mock.NewServer()
	defer close()
	srv.AppendResponse()
	opt := ConnectionOptions{}
	opt.HTTPClient = srv
	opt.Telemetry.ApplicationID = myTelemetry
	con := NewConnection(Endpoint(srv.URL()), mockTokenCred{}, &opt)
	if ep := con.Endpoint(); ep != Endpoint(srv.URL()) {
		t.Fatalf("unexpected endpoint %s", ep)
	}
	if opt.Telemetry.ApplicationID != myTelemetry {
		t.Fatalf("telemetry was modified: %s", opt.Telemetry.ApplicationID)
	}
	req, err := azruntime.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	resp, err := con.NewPipeline("armtest", "v1.2.3").Do(req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
	if ua := resp.Request.Header.Get("User-Agent"); !strings.HasPrefix(ua, myTelemetry+" "+"azsdk-go-armtest/v1.2.3") {
		t.Fatalf("unexpected User-Agent %s", ua)
	}
}

func TestDisableAutoRPRegistration(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	// initial response that RP is unregistered
	srv.SetResponse(mock.WithStatusCode(http.StatusConflict), mock.WithBody([]byte(rpUnregisteredResp)))
	con := NewConnection(Endpoint(srv.URL()), mockTokenCred{}, &ConnectionOptions{DisableRPRegistration: true})
	if ep := con.Endpoint(); ep != Endpoint(srv.URL()) {
		t.Fatalf("unexpected endpoint %s", ep)
	}
	req, err := azruntime.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	// log only RP registration
	log.SetClassifications(armruntime.LogRPRegistration)
	defer func() {
		// reset logging
		log.SetClassifications()
	}()
	logEntries := 0
	log.SetListener(func(cls log.Classification, msg string) {
		logEntries++
	})
	resp, err := con.NewPipeline("armtest", "v1.2.3").Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusConflict {
		t.Fatalf("unexpected status code %d:", resp.StatusCode)
	}
	// shouldn't be any log entries
	if logEntries != 0 {
		t.Fatalf("expected 0 log entries, got %d", logEntries)
	}
}

// policy that tracks the number of times it was invoked
type countingPolicy struct {
	count int
}

func (p *countingPolicy) Do(req *policy.Request) (*http.Response, error) {
	p.count++
	return req.Next()
}

func TestConnectionWithCustomPolicies(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	// initial response is a failure to trigger retry
	srv.AppendResponse(mock.WithStatusCode(http.StatusInternalServerError))
	srv.AppendResponse(mock.WithStatusCode(http.StatusOK))
	perCallPolicy := countingPolicy{}
	perRetryPolicy := countingPolicy{}
	con := NewConnection(Endpoint(srv.URL()), mockTokenCred{}, &ConnectionOptions{
		DisableRPRegistration: true,
		PerCallPolicies:       []policy.Policy{&perCallPolicy},
		PerRetryPolicies:      []policy.Policy{&perRetryPolicy},
	})
	req, err := azruntime.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatal(err)
	}
	resp, err := con.NewPipeline("armtest", "v1.2.3").Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code %d", resp.StatusCode)
	}
	if perCallPolicy.count != 1 {
		t.Fatalf("unexpected per call policy count %d", perCallPolicy.count)
	}
	if perRetryPolicy.count != 2 {
		t.Fatalf("unexpected per retry policy count %d", perRetryPolicy.count)
	}
}
