//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

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

// ProviderResourceTypesClient contains the methods for the ProviderResourceTypes group.
// Don't use this type directly, use NewProviderResourceTypesClient() instead.
type ProviderResourceTypesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewProviderResourceTypesClient creates a new instance of ProviderResourceTypesClient with the specified values.
func NewProviderResourceTypesClient(con *arm.Connection, subscriptionID string) *ProviderResourceTypesClient {
	return &ProviderResourceTypesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - List the resource types for a specified resource provider.
// If the operation fails it returns the *CloudError error type.
func (client *ProviderResourceTypesClient) List(ctx context.Context, resourceProviderNamespace string, options *ProviderResourceTypesListOptions) (ProviderResourceTypesListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceProviderNamespace, options)
	if err != nil {
		return ProviderResourceTypesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProviderResourceTypesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProviderResourceTypesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ProviderResourceTypesClient) listCreateRequest(ctx context.Context, resourceProviderNamespace string, options *ProviderResourceTypesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/resourceTypes"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
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
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProviderResourceTypesClient) listHandleResponse(resp *http.Response) (ProviderResourceTypesListResponse, error) {
	result := ProviderResourceTypesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderResourceTypeListResult); err != nil {
		return ProviderResourceTypesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ProviderResourceTypesClient) listHandleError(resp *http.Response) error {
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
