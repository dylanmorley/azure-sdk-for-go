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

// InstanceFailoverGroupsClient contains the methods for the InstanceFailoverGroups group.
// Don't use this type directly, use NewInstanceFailoverGroupsClient() instead.
type InstanceFailoverGroupsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewInstanceFailoverGroupsClient creates a new instance of InstanceFailoverGroupsClient with the specified values.
func NewInstanceFailoverGroupsClient(con *arm.Connection, subscriptionID string) *InstanceFailoverGroupsClient {
	return &InstanceFailoverGroupsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a failover group.
// If the operation fails it returns a generic error.
func (client *InstanceFailoverGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, parameters InstanceFailoverGroup, options *InstanceFailoverGroupsBeginCreateOrUpdateOptions) (InstanceFailoverGroupsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, locationName, failoverGroupName, parameters, options)
	if err != nil {
		return InstanceFailoverGroupsCreateOrUpdatePollerResponse{}, err
	}
	result := InstanceFailoverGroupsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("InstanceFailoverGroupsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return InstanceFailoverGroupsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &InstanceFailoverGroupsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a failover group.
// If the operation fails it returns a generic error.
func (client *InstanceFailoverGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, parameters InstanceFailoverGroup, options *InstanceFailoverGroupsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, locationName, failoverGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *InstanceFailoverGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, parameters InstanceFailoverGroup, options *InstanceFailoverGroupsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *InstanceFailoverGroupsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Deletes a failover group.
// If the operation fails it returns a generic error.
func (client *InstanceFailoverGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsBeginDeleteOptions) (InstanceFailoverGroupsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, locationName, failoverGroupName, options)
	if err != nil {
		return InstanceFailoverGroupsDeletePollerResponse{}, err
	}
	result := InstanceFailoverGroupsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("InstanceFailoverGroupsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return InstanceFailoverGroupsDeletePollerResponse{}, err
	}
	result.Poller = &InstanceFailoverGroupsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a failover group.
// If the operation fails it returns a generic error.
func (client *InstanceFailoverGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, locationName, failoverGroupName, options)
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
func (client *InstanceFailoverGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *InstanceFailoverGroupsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginFailover - Fails over from the current primary managed instance to this managed instance.
// If the operation fails it returns a generic error.
func (client *InstanceFailoverGroupsClient) BeginFailover(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsBeginFailoverOptions) (InstanceFailoverGroupsFailoverPollerResponse, error) {
	resp, err := client.failover(ctx, resourceGroupName, locationName, failoverGroupName, options)
	if err != nil {
		return InstanceFailoverGroupsFailoverPollerResponse{}, err
	}
	result := InstanceFailoverGroupsFailoverPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("InstanceFailoverGroupsClient.Failover", "", resp, client.pl, client.failoverHandleError)
	if err != nil {
		return InstanceFailoverGroupsFailoverPollerResponse{}, err
	}
	result.Poller = &InstanceFailoverGroupsFailoverPoller{
		pt: pt,
	}
	return result, nil
}

// Failover - Fails over from the current primary managed instance to this managed instance.
// If the operation fails it returns a generic error.
func (client *InstanceFailoverGroupsClient) failover(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsBeginFailoverOptions) (*http.Response, error) {
	req, err := client.failoverCreateRequest(ctx, resourceGroupName, locationName, failoverGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.failoverHandleError(resp)
	}
	return resp, nil
}

// failoverCreateRequest creates the Failover request.
func (client *InstanceFailoverGroupsClient) failoverCreateRequest(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsBeginFailoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}/failover"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
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
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// failoverHandleError handles the Failover error response.
func (client *InstanceFailoverGroupsClient) failoverHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginForceFailoverAllowDataLoss - Fails over from the current primary managed instance to this managed instance. This operation might result in data
// loss.
// If the operation fails it returns a generic error.
func (client *InstanceFailoverGroupsClient) BeginForceFailoverAllowDataLoss(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsBeginForceFailoverAllowDataLossOptions) (InstanceFailoverGroupsForceFailoverAllowDataLossPollerResponse, error) {
	resp, err := client.forceFailoverAllowDataLoss(ctx, resourceGroupName, locationName, failoverGroupName, options)
	if err != nil {
		return InstanceFailoverGroupsForceFailoverAllowDataLossPollerResponse{}, err
	}
	result := InstanceFailoverGroupsForceFailoverAllowDataLossPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("InstanceFailoverGroupsClient.ForceFailoverAllowDataLoss", "", resp, client.pl, client.forceFailoverAllowDataLossHandleError)
	if err != nil {
		return InstanceFailoverGroupsForceFailoverAllowDataLossPollerResponse{}, err
	}
	result.Poller = &InstanceFailoverGroupsForceFailoverAllowDataLossPoller{
		pt: pt,
	}
	return result, nil
}

// ForceFailoverAllowDataLoss - Fails over from the current primary managed instance to this managed instance. This operation might result in data loss.
// If the operation fails it returns a generic error.
func (client *InstanceFailoverGroupsClient) forceFailoverAllowDataLoss(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsBeginForceFailoverAllowDataLossOptions) (*http.Response, error) {
	req, err := client.forceFailoverAllowDataLossCreateRequest(ctx, resourceGroupName, locationName, failoverGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.forceFailoverAllowDataLossHandleError(resp)
	}
	return resp, nil
}

// forceFailoverAllowDataLossCreateRequest creates the ForceFailoverAllowDataLoss request.
func (client *InstanceFailoverGroupsClient) forceFailoverAllowDataLossCreateRequest(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsBeginForceFailoverAllowDataLossOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}/forceFailoverAllowDataLoss"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
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
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// forceFailoverAllowDataLossHandleError handles the ForceFailoverAllowDataLoss error response.
func (client *InstanceFailoverGroupsClient) forceFailoverAllowDataLossHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a failover group.
// If the operation fails it returns a generic error.
func (client *InstanceFailoverGroupsClient) Get(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsGetOptions) (InstanceFailoverGroupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, locationName, failoverGroupName, options)
	if err != nil {
		return InstanceFailoverGroupsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InstanceFailoverGroupsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InstanceFailoverGroupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InstanceFailoverGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, locationName string, failoverGroupName string, options *InstanceFailoverGroupsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups/{failoverGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
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
func (client *InstanceFailoverGroupsClient) getHandleResponse(resp *http.Response) (InstanceFailoverGroupsGetResponse, error) {
	result := InstanceFailoverGroupsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.InstanceFailoverGroup); err != nil {
		return InstanceFailoverGroupsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *InstanceFailoverGroupsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByLocation - Lists the failover groups in a location.
// If the operation fails it returns a generic error.
func (client *InstanceFailoverGroupsClient) ListByLocation(resourceGroupName string, locationName string, options *InstanceFailoverGroupsListByLocationOptions) *InstanceFailoverGroupsListByLocationPager {
	return &InstanceFailoverGroupsListByLocationPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByLocationCreateRequest(ctx, resourceGroupName, locationName, options)
		},
		advancer: func(ctx context.Context, resp InstanceFailoverGroupsListByLocationResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.InstanceFailoverGroupListResult.NextLink)
		},
	}
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *InstanceFailoverGroupsClient) listByLocationCreateRequest(ctx context.Context, resourceGroupName string, locationName string, options *InstanceFailoverGroupsListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/locations/{locationName}/instanceFailoverGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
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

// listByLocationHandleResponse handles the ListByLocation response.
func (client *InstanceFailoverGroupsClient) listByLocationHandleResponse(resp *http.Response) (InstanceFailoverGroupsListByLocationResponse, error) {
	result := InstanceFailoverGroupsListByLocationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.InstanceFailoverGroupListResult); err != nil {
		return InstanceFailoverGroupsListByLocationResponse{}, err
	}
	return result, nil
}

// listByLocationHandleError handles the ListByLocation error response.
func (client *InstanceFailoverGroupsClient) listByLocationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
