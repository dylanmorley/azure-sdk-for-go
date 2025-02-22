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

// ServerAdvisorsClient contains the methods for the ServerAdvisors group.
// Don't use this type directly, use NewServerAdvisorsClient() instead.
type ServerAdvisorsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewServerAdvisorsClient creates a new instance of ServerAdvisorsClient with the specified values.
func NewServerAdvisorsClient(con *arm.Connection, subscriptionID string) *ServerAdvisorsClient {
	return &ServerAdvisorsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Gets a server advisor.
// If the operation fails it returns a generic error.
func (client *ServerAdvisorsClient) Get(ctx context.Context, resourceGroupName string, serverName string, advisorName string, options *ServerAdvisorsGetOptions) (ServerAdvisorsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, advisorName, options)
	if err != nil {
		return ServerAdvisorsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerAdvisorsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAdvisorsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerAdvisorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, advisorName string, options *ServerAdvisorsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advisors/{advisorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
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
func (client *ServerAdvisorsClient) getHandleResponse(resp *http.Response) (ServerAdvisorsGetResponse, error) {
	result := ServerAdvisorsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Advisor); err != nil {
		return ServerAdvisorsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServerAdvisorsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByServer - Gets a list of server advisors.
// If the operation fails it returns a generic error.
func (client *ServerAdvisorsClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdvisorsListByServerOptions) (ServerAdvisorsListByServerResponse, error) {
	req, err := client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerAdvisorsListByServerResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerAdvisorsListByServerResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAdvisorsListByServerResponse{}, client.listByServerHandleError(resp)
	}
	return client.listByServerHandleResponse(resp)
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerAdvisorsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdvisorsListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advisors"
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerAdvisorsClient) listByServerHandleResponse(resp *http.Response) (ServerAdvisorsListByServerResponse, error) {
	result := ServerAdvisorsListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdvisorArray); err != nil {
		return ServerAdvisorsListByServerResponse{}, err
	}
	return result, nil
}

// listByServerHandleError handles the ListByServer error response.
func (client *ServerAdvisorsClient) listByServerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Update - Updates a server advisor.
// If the operation fails it returns a generic error.
func (client *ServerAdvisorsClient) Update(ctx context.Context, resourceGroupName string, serverName string, advisorName string, parameters Advisor, options *ServerAdvisorsUpdateOptions) (ServerAdvisorsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, advisorName, parameters, options)
	if err != nil {
		return ServerAdvisorsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerAdvisorsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAdvisorsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ServerAdvisorsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, advisorName string, parameters Advisor, options *ServerAdvisorsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advisors/{advisorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
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
func (client *ServerAdvisorsClient) updateHandleResponse(resp *http.Response) (ServerAdvisorsUpdateResponse, error) {
	result := ServerAdvisorsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Advisor); err != nil {
		return ServerAdvisorsUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *ServerAdvisorsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
