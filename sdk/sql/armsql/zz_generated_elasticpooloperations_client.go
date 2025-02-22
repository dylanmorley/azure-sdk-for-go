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

// ElasticPoolOperationsClient contains the methods for the ElasticPoolOperations group.
// Don't use this type directly, use NewElasticPoolOperationsClient() instead.
type ElasticPoolOperationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewElasticPoolOperationsClient creates a new instance of ElasticPoolOperationsClient with the specified values.
func NewElasticPoolOperationsClient(con *arm.Connection, subscriptionID string) *ElasticPoolOperationsClient {
	return &ElasticPoolOperationsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Cancel - Cancels the asynchronous operation on the elastic pool.
// If the operation fails it returns a generic error.
func (client *ElasticPoolOperationsClient) Cancel(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, operationID string, options *ElasticPoolOperationsCancelOptions) (ElasticPoolOperationsCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, operationID, options)
	if err != nil {
		return ElasticPoolOperationsCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ElasticPoolOperationsCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ElasticPoolOperationsCancelResponse{}, client.cancelHandleError(resp)
	}
	return ElasticPoolOperationsCancelResponse{RawResponse: resp}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *ElasticPoolOperationsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, operationID string, options *ElasticPoolOperationsCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/operations/{operationId}/cancel"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if elasticPoolName == "" {
		return nil, errors.New("parameter elasticPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticPoolName}", url.PathEscape(elasticPoolName))
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
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
	return req, nil
}

// cancelHandleError handles the Cancel error response.
func (client *ElasticPoolOperationsClient) cancelHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByElasticPool - Gets a list of operations performed on the elastic pool.
// If the operation fails it returns a generic error.
func (client *ElasticPoolOperationsClient) ListByElasticPool(resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolOperationsListByElasticPoolOptions) *ElasticPoolOperationsListByElasticPoolPager {
	return &ElasticPoolOperationsListByElasticPoolPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByElasticPoolCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, options)
		},
		advancer: func(ctx context.Context, resp ElasticPoolOperationsListByElasticPoolResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ElasticPoolOperationListResult.NextLink)
		},
	}
}

// listByElasticPoolCreateRequest creates the ListByElasticPool request.
func (client *ElasticPoolOperationsClient) listByElasticPoolCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolOperationsListByElasticPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/operations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if elasticPoolName == "" {
		return nil, errors.New("parameter elasticPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticPoolName}", url.PathEscape(elasticPoolName))
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

// listByElasticPoolHandleResponse handles the ListByElasticPool response.
func (client *ElasticPoolOperationsClient) listByElasticPoolHandleResponse(resp *http.Response) (ElasticPoolOperationsListByElasticPoolResponse, error) {
	result := ElasticPoolOperationsListByElasticPoolResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ElasticPoolOperationListResult); err != nil {
		return ElasticPoolOperationsListByElasticPoolResponse{}, err
	}
	return result, nil
}

// listByElasticPoolHandleError handles the ListByElasticPool error response.
func (client *ElasticPoolOperationsClient) listByElasticPoolHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
