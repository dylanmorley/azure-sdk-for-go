//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// VirtualMachineSizesClient contains the methods for the VirtualMachineSizes group.
// Don't use this type directly, use NewVirtualMachineSizesClient() instead.
type VirtualMachineSizesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVirtualMachineSizesClient creates a new instance of VirtualMachineSizesClient with the specified values.
func NewVirtualMachineSizesClient(con *arm.Connection, subscriptionID string) *VirtualMachineSizesClient {
	return &VirtualMachineSizesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - This API is deprecated. Use Resources Skus [https://docs.microsoft.com/rest/api/compute/resourceskus/list]
// If the operation fails it returns a generic error.
func (client *VirtualMachineSizesClient) List(ctx context.Context, location string, options *VirtualMachineSizesListOptions) (VirtualMachineSizesListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, options)
	if err != nil {
		return VirtualMachineSizesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineSizesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineSizesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *VirtualMachineSizesClient) listCreateRequest(ctx context.Context, location string, options *VirtualMachineSizesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/vmSizes"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineSizesClient) listHandleResponse(resp *http.Response) (VirtualMachineSizesListResponse, error) {
	result := VirtualMachineSizesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineSizeListResult); err != nil {
		return VirtualMachineSizesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VirtualMachineSizesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
