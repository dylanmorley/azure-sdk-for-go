//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/internal/mock"
)

func TestRequestIdPolicy(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.SetResponse(mock.WithStatusCode(http.StatusOK))
	pl := NewPipeline(srv, NewRequestIdPolicy())
	req, err := NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	body := newRewindTrackingBody("stuff")
	if err := req.SetBody(body, "text/plain"); err != nil {
		t.Fatal(err)
	}
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
	if resp.Request.Header.Get("x-ms-client-request-id") == "" {
		t.Fatalf("client request id header was not set")
	}
}
