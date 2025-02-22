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
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// DiskAccessesClient contains the methods for the DiskAccesses group.
// Don't use this type directly, use NewDiskAccessesClient() instead.
type DiskAccessesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDiskAccessesClient creates a new instance of DiskAccessesClient with the specified values.
func NewDiskAccessesClient(con *arm.Connection, subscriptionID string) *DiskAccessesClient {
	return &DiskAccessesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a disk access resource
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskAccessName string, diskAccess DiskAccess, options *DiskAccessesBeginCreateOrUpdateOptions) (DiskAccessesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, diskAccessName, diskAccess, options)
	if err != nil {
		return DiskAccessesCreateOrUpdatePollerResponse{}, err
	}
	result := DiskAccessesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DiskAccessesClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return DiskAccessesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DiskAccessesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a disk access resource
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) createOrUpdate(ctx context.Context, resourceGroupName string, diskAccessName string, diskAccess DiskAccess, options *DiskAccessesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, diskAccessName, diskAccess, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DiskAccessesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, diskAccessName string, diskAccess DiskAccess, options *DiskAccessesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskAccessName == "" {
		return nil, errors.New("parameter diskAccessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskAccessName}", url.PathEscape(diskAccessName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, diskAccess)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DiskAccessesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes a disk access resource.
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) BeginDelete(ctx context.Context, resourceGroupName string, diskAccessName string, options *DiskAccessesBeginDeleteOptions) (DiskAccessesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, diskAccessName, options)
	if err != nil {
		return DiskAccessesDeletePollerResponse{}, err
	}
	result := DiskAccessesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DiskAccessesClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return DiskAccessesDeletePollerResponse{}, err
	}
	result.Poller = &DiskAccessesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a disk access resource.
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) deleteOperation(ctx context.Context, resourceGroupName string, diskAccessName string, options *DiskAccessesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, diskAccessName, options)
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
func (client *DiskAccessesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, diskAccessName string, options *DiskAccessesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskAccessName == "" {
		return nil, errors.New("parameter diskAccessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskAccessName}", url.PathEscape(diskAccessName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DiskAccessesClient) deleteHandleError(resp *http.Response) error {
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

// BeginDeleteAPrivateEndpointConnection - Deletes a private endpoint connection under a disk access resource.
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) BeginDeleteAPrivateEndpointConnection(ctx context.Context, resourceGroupName string, diskAccessName string, privateEndpointConnectionName string, options *DiskAccessesBeginDeleteAPrivateEndpointConnectionOptions) (DiskAccessesDeleteAPrivateEndpointConnectionPollerResponse, error) {
	resp, err := client.deleteAPrivateEndpointConnection(ctx, resourceGroupName, diskAccessName, privateEndpointConnectionName, options)
	if err != nil {
		return DiskAccessesDeleteAPrivateEndpointConnectionPollerResponse{}, err
	}
	result := DiskAccessesDeleteAPrivateEndpointConnectionPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DiskAccessesClient.DeleteAPrivateEndpointConnection", "", resp, client.pl, client.deleteAPrivateEndpointConnectionHandleError)
	if err != nil {
		return DiskAccessesDeleteAPrivateEndpointConnectionPollerResponse{}, err
	}
	result.Poller = &DiskAccessesDeleteAPrivateEndpointConnectionPoller{
		pt: pt,
	}
	return result, nil
}

// DeleteAPrivateEndpointConnection - Deletes a private endpoint connection under a disk access resource.
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) deleteAPrivateEndpointConnection(ctx context.Context, resourceGroupName string, diskAccessName string, privateEndpointConnectionName string, options *DiskAccessesBeginDeleteAPrivateEndpointConnectionOptions) (*http.Response, error) {
	req, err := client.deleteAPrivateEndpointConnectionCreateRequest(ctx, resourceGroupName, diskAccessName, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteAPrivateEndpointConnectionHandleError(resp)
	}
	return resp, nil
}

// deleteAPrivateEndpointConnectionCreateRequest creates the DeleteAPrivateEndpointConnection request.
func (client *DiskAccessesClient) deleteAPrivateEndpointConnectionCreateRequest(ctx context.Context, resourceGroupName string, diskAccessName string, privateEndpointConnectionName string, options *DiskAccessesBeginDeleteAPrivateEndpointConnectionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskAccessName == "" {
		return nil, errors.New("parameter diskAccessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskAccessName}", url.PathEscape(diskAccessName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteAPrivateEndpointConnectionHandleError handles the DeleteAPrivateEndpointConnection error response.
func (client *DiskAccessesClient) deleteAPrivateEndpointConnectionHandleError(resp *http.Response) error {
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

// Get - Gets information about a disk access resource.
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) Get(ctx context.Context, resourceGroupName string, diskAccessName string, options *DiskAccessesGetOptions) (DiskAccessesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, diskAccessName, options)
	if err != nil {
		return DiskAccessesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiskAccessesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiskAccessesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiskAccessesClient) getCreateRequest(ctx context.Context, resourceGroupName string, diskAccessName string, options *DiskAccessesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskAccessName == "" {
		return nil, errors.New("parameter diskAccessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskAccessName}", url.PathEscape(diskAccessName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiskAccessesClient) getHandleResponse(resp *http.Response) (DiskAccessesGetResponse, error) {
	result := DiskAccessesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskAccess); err != nil {
		return DiskAccessesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DiskAccessesClient) getHandleError(resp *http.Response) error {
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

// GetAPrivateEndpointConnection - Gets information about a private endpoint connection under a disk access resource.
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) GetAPrivateEndpointConnection(ctx context.Context, resourceGroupName string, diskAccessName string, privateEndpointConnectionName string, options *DiskAccessesGetAPrivateEndpointConnectionOptions) (DiskAccessesGetAPrivateEndpointConnectionResponse, error) {
	req, err := client.getAPrivateEndpointConnectionCreateRequest(ctx, resourceGroupName, diskAccessName, privateEndpointConnectionName, options)
	if err != nil {
		return DiskAccessesGetAPrivateEndpointConnectionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiskAccessesGetAPrivateEndpointConnectionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiskAccessesGetAPrivateEndpointConnectionResponse{}, client.getAPrivateEndpointConnectionHandleError(resp)
	}
	return client.getAPrivateEndpointConnectionHandleResponse(resp)
}

// getAPrivateEndpointConnectionCreateRequest creates the GetAPrivateEndpointConnection request.
func (client *DiskAccessesClient) getAPrivateEndpointConnectionCreateRequest(ctx context.Context, resourceGroupName string, diskAccessName string, privateEndpointConnectionName string, options *DiskAccessesGetAPrivateEndpointConnectionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskAccessName == "" {
		return nil, errors.New("parameter diskAccessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskAccessName}", url.PathEscape(diskAccessName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAPrivateEndpointConnectionHandleResponse handles the GetAPrivateEndpointConnection response.
func (client *DiskAccessesClient) getAPrivateEndpointConnectionHandleResponse(resp *http.Response) (DiskAccessesGetAPrivateEndpointConnectionResponse, error) {
	result := DiskAccessesGetAPrivateEndpointConnectionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return DiskAccessesGetAPrivateEndpointConnectionResponse{}, err
	}
	return result, nil
}

// getAPrivateEndpointConnectionHandleError handles the GetAPrivateEndpointConnection error response.
func (client *DiskAccessesClient) getAPrivateEndpointConnectionHandleError(resp *http.Response) error {
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

// GetPrivateLinkResources - Gets the private link resources possible under disk access resource
// If the operation fails it returns a generic error.
func (client *DiskAccessesClient) GetPrivateLinkResources(ctx context.Context, resourceGroupName string, diskAccessName string, options *DiskAccessesGetPrivateLinkResourcesOptions) (DiskAccessesGetPrivateLinkResourcesResponse, error) {
	req, err := client.getPrivateLinkResourcesCreateRequest(ctx, resourceGroupName, diskAccessName, options)
	if err != nil {
		return DiskAccessesGetPrivateLinkResourcesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiskAccessesGetPrivateLinkResourcesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiskAccessesGetPrivateLinkResourcesResponse{}, client.getPrivateLinkResourcesHandleError(resp)
	}
	return client.getPrivateLinkResourcesHandleResponse(resp)
}

// getPrivateLinkResourcesCreateRequest creates the GetPrivateLinkResources request.
func (client *DiskAccessesClient) getPrivateLinkResourcesCreateRequest(ctx context.Context, resourceGroupName string, diskAccessName string, options *DiskAccessesGetPrivateLinkResourcesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskAccessName == "" {
		return nil, errors.New("parameter diskAccessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskAccessName}", url.PathEscape(diskAccessName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getPrivateLinkResourcesHandleResponse handles the GetPrivateLinkResources response.
func (client *DiskAccessesClient) getPrivateLinkResourcesHandleResponse(resp *http.Response) (DiskAccessesGetPrivateLinkResourcesResponse, error) {
	result := DiskAccessesGetPrivateLinkResourcesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return DiskAccessesGetPrivateLinkResourcesResponse{}, err
	}
	return result, nil
}

// getPrivateLinkResourcesHandleError handles the GetPrivateLinkResources error response.
func (client *DiskAccessesClient) getPrivateLinkResourcesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Lists all the disk access resources under a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) List(options *DiskAccessesListOptions) *DiskAccessesListPager {
	return &DiskAccessesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DiskAccessesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DiskAccessList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DiskAccessesClient) listCreateRequest(ctx context.Context, options *DiskAccessesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/diskAccesses"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DiskAccessesClient) listHandleResponse(resp *http.Response) (DiskAccessesListResponse, error) {
	result := DiskAccessesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskAccessList); err != nil {
		return DiskAccessesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *DiskAccessesClient) listHandleError(resp *http.Response) error {
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

// ListByResourceGroup - Lists all the disk access resources under a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) ListByResourceGroup(resourceGroupName string, options *DiskAccessesListByResourceGroupOptions) *DiskAccessesListByResourceGroupPager {
	return &DiskAccessesListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp DiskAccessesListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DiskAccessList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DiskAccessesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DiskAccessesListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses"
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
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DiskAccessesClient) listByResourceGroupHandleResponse(resp *http.Response) (DiskAccessesListByResourceGroupResponse, error) {
	result := DiskAccessesListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskAccessList); err != nil {
		return DiskAccessesListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *DiskAccessesClient) listByResourceGroupHandleError(resp *http.Response) error {
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

// ListPrivateEndpointConnections - List information about private endpoint connections under a disk access resource
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) ListPrivateEndpointConnections(resourceGroupName string, diskAccessName string, options *DiskAccessesListPrivateEndpointConnectionsOptions) *DiskAccessesListPrivateEndpointConnectionsPager {
	return &DiskAccessesListPrivateEndpointConnectionsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listPrivateEndpointConnectionsCreateRequest(ctx, resourceGroupName, diskAccessName, options)
		},
		advancer: func(ctx context.Context, resp DiskAccessesListPrivateEndpointConnectionsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PrivateEndpointConnectionListResult.NextLink)
		},
	}
}

// listPrivateEndpointConnectionsCreateRequest creates the ListPrivateEndpointConnections request.
func (client *DiskAccessesClient) listPrivateEndpointConnectionsCreateRequest(ctx context.Context, resourceGroupName string, diskAccessName string, options *DiskAccessesListPrivateEndpointConnectionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskAccessName == "" {
		return nil, errors.New("parameter diskAccessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskAccessName}", url.PathEscape(diskAccessName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listPrivateEndpointConnectionsHandleResponse handles the ListPrivateEndpointConnections response.
func (client *DiskAccessesClient) listPrivateEndpointConnectionsHandleResponse(resp *http.Response) (DiskAccessesListPrivateEndpointConnectionsResponse, error) {
	result := DiskAccessesListPrivateEndpointConnectionsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResult); err != nil {
		return DiskAccessesListPrivateEndpointConnectionsResponse{}, err
	}
	return result, nil
}

// listPrivateEndpointConnectionsHandleError handles the ListPrivateEndpointConnections error response.
func (client *DiskAccessesClient) listPrivateEndpointConnectionsHandleError(resp *http.Response) error {
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

// BeginUpdate - Updates (patches) a disk access resource.
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) BeginUpdate(ctx context.Context, resourceGroupName string, diskAccessName string, diskAccess DiskAccessUpdate, options *DiskAccessesBeginUpdateOptions) (DiskAccessesUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, diskAccessName, diskAccess, options)
	if err != nil {
		return DiskAccessesUpdatePollerResponse{}, err
	}
	result := DiskAccessesUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DiskAccessesClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return DiskAccessesUpdatePollerResponse{}, err
	}
	result.Poller = &DiskAccessesUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates (patches) a disk access resource.
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) update(ctx context.Context, resourceGroupName string, diskAccessName string, diskAccess DiskAccessUpdate, options *DiskAccessesBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, diskAccessName, diskAccess, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DiskAccessesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, diskAccessName string, diskAccess DiskAccessUpdate, options *DiskAccessesBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskAccessName == "" {
		return nil, errors.New("parameter diskAccessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskAccessName}", url.PathEscape(diskAccessName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, diskAccess)
}

// updateHandleError handles the Update error response.
func (client *DiskAccessesClient) updateHandleError(resp *http.Response) error {
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

// BeginUpdateAPrivateEndpointConnection - Approve or reject a private endpoint connection under disk access resource, this can't be used to create a new
// private endpoint connection.
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) BeginUpdateAPrivateEndpointConnection(ctx context.Context, resourceGroupName string, diskAccessName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection, options *DiskAccessesBeginUpdateAPrivateEndpointConnectionOptions) (DiskAccessesUpdateAPrivateEndpointConnectionPollerResponse, error) {
	resp, err := client.updateAPrivateEndpointConnection(ctx, resourceGroupName, diskAccessName, privateEndpointConnectionName, privateEndpointConnection, options)
	if err != nil {
		return DiskAccessesUpdateAPrivateEndpointConnectionPollerResponse{}, err
	}
	result := DiskAccessesUpdateAPrivateEndpointConnectionPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DiskAccessesClient.UpdateAPrivateEndpointConnection", "", resp, client.pl, client.updateAPrivateEndpointConnectionHandleError)
	if err != nil {
		return DiskAccessesUpdateAPrivateEndpointConnectionPollerResponse{}, err
	}
	result.Poller = &DiskAccessesUpdateAPrivateEndpointConnectionPoller{
		pt: pt,
	}
	return result, nil
}

// UpdateAPrivateEndpointConnection - Approve or reject a private endpoint connection under disk access resource, this can't be used to create a new private
// endpoint connection.
// If the operation fails it returns the *CloudError error type.
func (client *DiskAccessesClient) updateAPrivateEndpointConnection(ctx context.Context, resourceGroupName string, diskAccessName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection, options *DiskAccessesBeginUpdateAPrivateEndpointConnectionOptions) (*http.Response, error) {
	req, err := client.updateAPrivateEndpointConnectionCreateRequest(ctx, resourceGroupName, diskAccessName, privateEndpointConnectionName, privateEndpointConnection, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateAPrivateEndpointConnectionHandleError(resp)
	}
	return resp, nil
}

// updateAPrivateEndpointConnectionCreateRequest creates the UpdateAPrivateEndpointConnection request.
func (client *DiskAccessesClient) updateAPrivateEndpointConnectionCreateRequest(ctx context.Context, resourceGroupName string, diskAccessName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection, options *DiskAccessesBeginUpdateAPrivateEndpointConnectionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskAccessName == "" {
		return nil, errors.New("parameter diskAccessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskAccessName}", url.PathEscape(diskAccessName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, privateEndpointConnection)
}

// updateAPrivateEndpointConnectionHandleError handles the UpdateAPrivateEndpointConnection error response.
func (client *DiskAccessesClient) updateAPrivateEndpointConnectionHandleError(resp *http.Response) error {
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
