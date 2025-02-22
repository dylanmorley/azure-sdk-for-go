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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ServerAutomaticTuningClient contains the methods for the ServerAutomaticTuning group.
// Don't use this type directly, use NewServerAutomaticTuningClient() instead.
type ServerAutomaticTuningClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewServerAutomaticTuningClient creates a new instance of ServerAutomaticTuningClient with the specified values.
func NewServerAutomaticTuningClient(con *arm.Connection, subscriptionID string) *ServerAutomaticTuningClient {
	return &ServerAutomaticTuningClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Retrieves server automatic tuning options.
// If the operation fails it returns a generic error.
func (client *ServerAutomaticTuningClient) Get(ctx context.Context, resourceGroupName string, serverName string, options *ServerAutomaticTuningGetOptions) (ServerAutomaticTuningGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerAutomaticTuningGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerAutomaticTuningGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAutomaticTuningGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerAutomaticTuningClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAutomaticTuningGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/automaticTuning/current"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerAutomaticTuningClient) getHandleResponse(resp *http.Response) (ServerAutomaticTuningGetResponse, error) {
	result := ServerAutomaticTuningGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerAutomaticTuning); err != nil {
		return ServerAutomaticTuningGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServerAutomaticTuningClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Update - Update automatic tuning options on server.
// If the operation fails it returns a generic error.
func (client *ServerAutomaticTuningClient) Update(ctx context.Context, resourceGroupName string, serverName string, parameters ServerAutomaticTuning, options *ServerAutomaticTuningUpdateOptions) (ServerAutomaticTuningUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, parameters, options)
	if err != nil {
		return ServerAutomaticTuningUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerAutomaticTuningUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAutomaticTuningUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ServerAutomaticTuningClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, parameters ServerAutomaticTuning, options *ServerAutomaticTuningUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/automaticTuning/current"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ServerAutomaticTuningClient) updateHandleResponse(resp *http.Response) (ServerAutomaticTuningUpdateResponse, error) {
	result := ServerAutomaticTuningUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerAutomaticTuning); err != nil {
		return ServerAutomaticTuningUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *ServerAutomaticTuningClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
