//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// RecoverableServersClient contains the methods for the RecoverableServers group.
// Don't use this type directly, use NewRecoverableServersClient() instead.
type RecoverableServersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRecoverableServersClient creates a new instance of RecoverableServersClient with the specified values.
func NewRecoverableServersClient(con *arm.Connection, subscriptionID string) *RecoverableServersClient {
	return &RecoverableServersClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Gets a recoverable PostgreSQL Server.
// If the operation fails it returns the *CloudError error type.
func (client *RecoverableServersClient) Get(ctx context.Context, resourceGroupName string, serverName string, options *RecoverableServersGetOptions) (RecoverableServersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return RecoverableServersGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecoverableServersGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecoverableServersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RecoverableServersClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *RecoverableServersGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/recoverableServers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecoverableServersClient) getHandleResponse(resp *http.Response) (RecoverableServersGetResponse, error) {
	result := RecoverableServersGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoverableServerResource); err != nil {
		return RecoverableServersGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RecoverableServersClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
