//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// VirtualHubsClient contains the methods for the VirtualHubs group.
// Don't use this type directly, use NewVirtualHubsClient() instead.
type VirtualHubsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVirtualHubsClient creates a new instance of VirtualHubsClient with the specified values.
func NewVirtualHubsClient(con *arm.Connection, subscriptionID string) *VirtualHubsClient {
	return &VirtualHubsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a VirtualHub resource if it doesn't exist else updates the existing VirtualHub.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, virtualHubParameters VirtualHub, options *VirtualHubsBeginCreateOrUpdateOptions) (VirtualHubsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualHubName, virtualHubParameters, options)
	if err != nil {
		return VirtualHubsCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualHubsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualHubsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualHubsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualHubsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a VirtualHub resource if it doesn't exist else updates the existing VirtualHub.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubsClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, virtualHubParameters VirtualHub, options *VirtualHubsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, virtualHubParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualHubsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, virtualHubParameters VirtualHub, options *VirtualHubsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, virtualHubParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualHubsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes a VirtualHub.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubsBeginDeleteOptions) (VirtualHubsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, virtualHubName, options)
	if err != nil {
		return VirtualHubsDeletePollerResponse{}, err
	}
	result := VirtualHubsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualHubsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return VirtualHubsDeletePollerResponse{}, err
	}
	result.Poller = &VirtualHubsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a VirtualHub.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubsClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualHubName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualHubsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualHubsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Retrieves the details of a VirtualHub.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubsClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubsGetOptions) (VirtualHubsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, options)
	if err != nil {
		return VirtualHubsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualHubsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualHubsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualHubsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualHubsClient) getHandleResponse(resp *http.Response) (VirtualHubsGetResponse, error) {
	result := VirtualHubsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualHub); err != nil {
		return VirtualHubsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualHubsClient) getHandleError(resp *http.Response) error {
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

// BeginGetEffectiveVirtualHubRoutes - Gets the effective routes configured for the Virtual Hub resource or the specified resource .
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubsClient) BeginGetEffectiveVirtualHubRoutes(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubsBeginGetEffectiveVirtualHubRoutesOptions) (VirtualHubsGetEffectiveVirtualHubRoutesPollerResponse, error) {
	resp, err := client.getEffectiveVirtualHubRoutes(ctx, resourceGroupName, virtualHubName, options)
	if err != nil {
		return VirtualHubsGetEffectiveVirtualHubRoutesPollerResponse{}, err
	}
	result := VirtualHubsGetEffectiveVirtualHubRoutesPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualHubsClient.GetEffectiveVirtualHubRoutes", "location", resp, client.pl, client.getEffectiveVirtualHubRoutesHandleError)
	if err != nil {
		return VirtualHubsGetEffectiveVirtualHubRoutesPollerResponse{}, err
	}
	result.Poller = &VirtualHubsGetEffectiveVirtualHubRoutesPoller{
		pt: pt,
	}
	return result, nil
}

// GetEffectiveVirtualHubRoutes - Gets the effective routes configured for the Virtual Hub resource or the specified resource .
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubsClient) getEffectiveVirtualHubRoutes(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubsBeginGetEffectiveVirtualHubRoutesOptions) (*http.Response, error) {
	req, err := client.getEffectiveVirtualHubRoutesCreateRequest(ctx, resourceGroupName, virtualHubName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.getEffectiveVirtualHubRoutesHandleError(resp)
	}
	return resp, nil
}

// getEffectiveVirtualHubRoutesCreateRequest creates the GetEffectiveVirtualHubRoutes request.
func (client *VirtualHubsClient) getEffectiveVirtualHubRoutesCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubsBeginGetEffectiveVirtualHubRoutesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/effectiveRoutes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.EffectiveRoutesParameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.EffectiveRoutesParameters)
	}
	return req, nil
}

// getEffectiveVirtualHubRoutesHandleError handles the GetEffectiveVirtualHubRoutes error response.
func (client *VirtualHubsClient) getEffectiveVirtualHubRoutesHandleError(resp *http.Response) error {
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

// List - Lists all the VirtualHubs in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubsClient) List(options *VirtualHubsListOptions) *VirtualHubsListPager {
	return &VirtualHubsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp VirtualHubsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVirtualHubsResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualHubsClient) listCreateRequest(ctx context.Context, options *VirtualHubsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualHubs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualHubsClient) listHandleResponse(resp *http.Response) (VirtualHubsListResponse, error) {
	result := VirtualHubsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVirtualHubsResult); err != nil {
		return VirtualHubsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VirtualHubsClient) listHandleError(resp *http.Response) error {
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

// ListByResourceGroup - Lists all the VirtualHubs in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubsClient) ListByResourceGroup(resourceGroupName string, options *VirtualHubsListByResourceGroupOptions) *VirtualHubsListByResourceGroupPager {
	return &VirtualHubsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp VirtualHubsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVirtualHubsResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualHubsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualHubsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualHubsClient) listByResourceGroupHandleResponse(resp *http.Response) (VirtualHubsListByResourceGroupResponse, error) {
	result := VirtualHubsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVirtualHubsResult); err != nil {
		return VirtualHubsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *VirtualHubsClient) listByResourceGroupHandleError(resp *http.Response) error {
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

// UpdateTags - Updates VirtualHub tags.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubsClient) UpdateTags(ctx context.Context, resourceGroupName string, virtualHubName string, virtualHubParameters TagsObject, options *VirtualHubsUpdateTagsOptions) (VirtualHubsUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, virtualHubName, virtualHubParameters, options)
	if err != nil {
		return VirtualHubsUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualHubsUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualHubsUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VirtualHubsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, virtualHubParameters TagsObject, options *VirtualHubsUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, virtualHubParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VirtualHubsClient) updateTagsHandleResponse(resp *http.Response) (VirtualHubsUpdateTagsResponse, error) {
	result := VirtualHubsUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualHub); err != nil {
		return VirtualHubsUpdateTagsResponse{}, err
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *VirtualHubsClient) updateTagsHandleError(resp *http.Response) error {
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
