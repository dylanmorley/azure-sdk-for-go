//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagedInstanceTdeCertificatesClient contains the methods for the ManagedInstanceTdeCertificates group.
// Don't use this type directly, use NewManagedInstanceTdeCertificatesClient() instead.
type ManagedInstanceTdeCertificatesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewManagedInstanceTdeCertificatesClient creates a new instance of ManagedInstanceTdeCertificatesClient with the specified values.
func NewManagedInstanceTdeCertificatesClient(con *arm.Connection, subscriptionID string) *ManagedInstanceTdeCertificatesClient {
	return &ManagedInstanceTdeCertificatesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreate - Creates a TDE certificate for a given server.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceTdeCertificatesClient) BeginCreate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters TdeCertificate, options *ManagedInstanceTdeCertificatesBeginCreateOptions) (ManagedInstanceTdeCertificatesCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, managedInstanceName, parameters, options)
	if err != nil {
		return ManagedInstanceTdeCertificatesCreatePollerResponse{}, err
	}
	result := ManagedInstanceTdeCertificatesCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ManagedInstanceTdeCertificatesClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return ManagedInstanceTdeCertificatesCreatePollerResponse{}, err
	}
	result.Poller = &ManagedInstanceTdeCertificatesCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a TDE certificate for a given server.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceTdeCertificatesClient) create(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters TdeCertificate, options *ManagedInstanceTdeCertificatesBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, managedInstanceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ManagedInstanceTdeCertificatesClient) createCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters TdeCertificate, options *ManagedInstanceTdeCertificatesBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/tdeCertificates"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleError handles the Create error response.
func (client *ManagedInstanceTdeCertificatesClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
