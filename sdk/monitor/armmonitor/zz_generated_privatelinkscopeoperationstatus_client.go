//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// PrivateLinkScopeOperationStatusClient contains the methods for the PrivateLinkScopeOperationStatus group.
// Don't use this type directly, use NewPrivateLinkScopeOperationStatusClient() instead.
type PrivateLinkScopeOperationStatusClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPrivateLinkScopeOperationStatusClient creates a new instance of PrivateLinkScopeOperationStatusClient with the specified values.
func NewPrivateLinkScopeOperationStatusClient(con *arm.Connection, subscriptionID string) *PrivateLinkScopeOperationStatusClient {
	return &PrivateLinkScopeOperationStatusClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Get the status of an azure asynchronous operation associated with a private link scope operation.
// If the operation fails it returns a generic error.
func (client *PrivateLinkScopeOperationStatusClient) Get(ctx context.Context, asyncOperationID string, resourceGroupName string, options *PrivateLinkScopeOperationStatusGetOptions) (PrivateLinkScopeOperationStatusGetResponse, error) {
	req, err := client.getCreateRequest(ctx, asyncOperationID, resourceGroupName, options)
	if err != nil {
		return PrivateLinkScopeOperationStatusGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkScopeOperationStatusGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkScopeOperationStatusGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkScopeOperationStatusClient) getCreateRequest(ctx context.Context, asyncOperationID string, resourceGroupName string, options *PrivateLinkScopeOperationStatusGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/privateLinkScopeOperationStatuses/{asyncOperationId}"
	if asyncOperationID == "" {
		return nil, errors.New("parameter asyncOperationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{asyncOperationId}", url.PathEscape(asyncOperationID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkScopeOperationStatusClient) getHandleResponse(resp *http.Response) (PrivateLinkScopeOperationStatusGetResponse, error) {
	result := PrivateLinkScopeOperationStatusGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatus); err != nil {
		return PrivateLinkScopeOperationStatusGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PrivateLinkScopeOperationStatusClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
