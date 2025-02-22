//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredis

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// RedisClient contains the methods for the Redis group.
// Don't use this type directly, use NewRedisClient() instead.
type RedisClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRedisClient creates a new instance of RedisClient with the specified values.
func NewRedisClient(con *arm.Connection, subscriptionID string) *RedisClient {
	return &RedisClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CheckNameAvailability - Checks that the redis cache name is valid and is not already in use.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) CheckNameAvailability(ctx context.Context, parameters CheckNameAvailabilityParameters, options *RedisCheckNameAvailabilityOptions) (RedisCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, parameters, options)
	if err != nil {
		return RedisCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RedisCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RedisCheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return RedisCheckNameAvailabilityResponse{RawResponse: resp}, nil
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *RedisClient) checkNameAvailabilityCreateRequest(ctx context.Context, parameters CheckNameAvailabilityParameters, options *RedisCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Cache/CheckNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *RedisClient) checkNameAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginCreate - Create or replace (overwrite/recreate, with potential downtime) an existing Redis cache.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) BeginCreate(ctx context.Context, resourceGroupName string, name string, parameters RedisCreateParameters, options *RedisBeginCreateOptions) (RedisCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return RedisCreatePollerResponse{}, err
	}
	result := RedisCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RedisClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return RedisCreatePollerResponse{}, err
	}
	result.Poller = &RedisCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Create or replace (overwrite/recreate, with potential downtime) an existing Redis cache.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) create(ctx context.Context, resourceGroupName string, name string, parameters RedisCreateParameters, options *RedisBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *RedisClient) createCreateRequest(ctx context.Context, resourceGroupName string, name string, parameters RedisCreateParameters, options *RedisBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleError handles the Create error response.
func (client *RedisClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a Redis cache.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) BeginDelete(ctx context.Context, resourceGroupName string, name string, options *RedisBeginDeleteOptions) (RedisDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, name, options)
	if err != nil {
		return RedisDeletePollerResponse{}, err
	}
	result := RedisDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RedisClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return RedisDeletePollerResponse{}, err
	}
	result.Poller = &RedisDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a Redis cache.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) deleteOperation(ctx context.Context, resourceGroupName string, name string, options *RedisBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, options)
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
func (client *RedisClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, options *RedisBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *RedisClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginExportData - Export data from the redis cache to blobs in a container.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) BeginExportData(ctx context.Context, resourceGroupName string, name string, parameters ExportRDBParameters, options *RedisBeginExportDataOptions) (RedisExportDataPollerResponse, error) {
	resp, err := client.exportData(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return RedisExportDataPollerResponse{}, err
	}
	result := RedisExportDataPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RedisClient.ExportData", "", resp, client.pl, client.exportDataHandleError)
	if err != nil {
		return RedisExportDataPollerResponse{}, err
	}
	result.Poller = &RedisExportDataPoller{
		pt: pt,
	}
	return result, nil
}

// ExportData - Export data from the redis cache to blobs in a container.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) exportData(ctx context.Context, resourceGroupName string, name string, parameters ExportRDBParameters, options *RedisBeginExportDataOptions) (*http.Response, error) {
	req, err := client.exportDataCreateRequest(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.exportDataHandleError(resp)
	}
	return resp, nil
}

// exportDataCreateRequest creates the ExportData request.
func (client *RedisClient) exportDataCreateRequest(ctx context.Context, resourceGroupName string, name string, parameters ExportRDBParameters, options *RedisBeginExportDataOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/export"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// exportDataHandleError handles the ExportData error response.
func (client *RedisClient) exportDataHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ForceReboot - Reboot specified Redis node(s). This operation requires write permission to the cache resource. There can be potential data loss.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) ForceReboot(ctx context.Context, resourceGroupName string, name string, parameters RedisRebootParameters, options *RedisForceRebootOptions) (RedisForceRebootResponseEnvelope, error) {
	req, err := client.forceRebootCreateRequest(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return RedisForceRebootResponseEnvelope{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RedisForceRebootResponseEnvelope{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RedisForceRebootResponseEnvelope{}, client.forceRebootHandleError(resp)
	}
	return client.forceRebootHandleResponse(resp)
}

// forceRebootCreateRequest creates the ForceReboot request.
func (client *RedisClient) forceRebootCreateRequest(ctx context.Context, resourceGroupName string, name string, parameters RedisRebootParameters, options *RedisForceRebootOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/forceReboot"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// forceRebootHandleResponse handles the ForceReboot response.
func (client *RedisClient) forceRebootHandleResponse(resp *http.Response) (RedisForceRebootResponseEnvelope, error) {
	result := RedisForceRebootResponseEnvelope{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisForceRebootResponse); err != nil {
		return RedisForceRebootResponseEnvelope{}, err
	}
	return result, nil
}

// forceRebootHandleError handles the ForceReboot error response.
func (client *RedisClient) forceRebootHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a Redis cache (resource description).
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) Get(ctx context.Context, resourceGroupName string, name string, options *RedisGetOptions) (RedisGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return RedisGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RedisGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RedisGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RedisClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, options *RedisGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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

// getHandleResponse handles the Get response.
func (client *RedisClient) getHandleResponse(resp *http.Response) (RedisGetResponse, error) {
	result := RedisGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisResource); err != nil {
		return RedisGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RedisClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginImportData - Import data into Redis cache.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) BeginImportData(ctx context.Context, resourceGroupName string, name string, parameters ImportRDBParameters, options *RedisBeginImportDataOptions) (RedisImportDataPollerResponse, error) {
	resp, err := client.importData(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return RedisImportDataPollerResponse{}, err
	}
	result := RedisImportDataPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RedisClient.ImportData", "", resp, client.pl, client.importDataHandleError)
	if err != nil {
		return RedisImportDataPollerResponse{}, err
	}
	result.Poller = &RedisImportDataPoller{
		pt: pt,
	}
	return result, nil
}

// ImportData - Import data into Redis cache.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) importData(ctx context.Context, resourceGroupName string, name string, parameters ImportRDBParameters, options *RedisBeginImportDataOptions) (*http.Response, error) {
	req, err := client.importDataCreateRequest(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.importDataHandleError(resp)
	}
	return resp, nil
}

// importDataCreateRequest creates the ImportData request.
func (client *RedisClient) importDataCreateRequest(ctx context.Context, resourceGroupName string, name string, parameters ImportRDBParameters, options *RedisBeginImportDataOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/import"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// importDataHandleError handles the ImportData error response.
func (client *RedisClient) importDataHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Lists all Redis caches in a resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) ListByResourceGroup(resourceGroupName string, options *RedisListByResourceGroupOptions) *RedisListByResourceGroupPager {
	return &RedisListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp RedisListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RedisListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *RedisClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *RedisListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis"
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
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *RedisClient) listByResourceGroupHandleResponse(resp *http.Response) (RedisListByResourceGroupResponse, error) {
	result := RedisListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisListResult); err != nil {
		return RedisListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *RedisClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Gets all Redis caches in the specified subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) ListBySubscription(options *RedisListBySubscriptionOptions) *RedisListBySubscriptionPager {
	return &RedisListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp RedisListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RedisListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *RedisClient) listBySubscriptionCreateRequest(ctx context.Context, options *RedisListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Cache/redis"
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

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *RedisClient) listBySubscriptionHandleResponse(resp *http.Response) (RedisListBySubscriptionResponse, error) {
	result := RedisListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisListResult); err != nil {
		return RedisListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *RedisClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListKeys - Retrieve a Redis cache's access keys. This operation requires write permission to the cache resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) ListKeys(ctx context.Context, resourceGroupName string, name string, options *RedisListKeysOptions) (RedisListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return RedisListKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RedisListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RedisListKeysResponse{}, client.listKeysHandleError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *RedisClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, name string, options *RedisListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/listKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *RedisClient) listKeysHandleResponse(resp *http.Response) (RedisListKeysResponse, error) {
	result := RedisListKeysResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisAccessKeys); err != nil {
		return RedisListKeysResponse{}, err
	}
	return result, nil
}

// listKeysHandleError handles the ListKeys error response.
func (client *RedisClient) listKeysHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListUpgradeNotifications - Gets any upgrade notifications for a Redis cache.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) ListUpgradeNotifications(resourceGroupName string, name string, history float64, options *RedisListUpgradeNotificationsOptions) *RedisListUpgradeNotificationsPager {
	return &RedisListUpgradeNotificationsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listUpgradeNotificationsCreateRequest(ctx, resourceGroupName, name, history, options)
		},
		advancer: func(ctx context.Context, resp RedisListUpgradeNotificationsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.NotificationListResponse.NextLink)
		},
	}
}

// listUpgradeNotificationsCreateRequest creates the ListUpgradeNotifications request.
func (client *RedisClient) listUpgradeNotificationsCreateRequest(ctx context.Context, resourceGroupName string, name string, history float64, options *RedisListUpgradeNotificationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/listUpgradeNotifications"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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
	reqQP.Set("history", strconv.FormatFloat(history, 'f', -1, 64))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listUpgradeNotificationsHandleResponse handles the ListUpgradeNotifications response.
func (client *RedisClient) listUpgradeNotificationsHandleResponse(resp *http.Response) (RedisListUpgradeNotificationsResponse, error) {
	result := RedisListUpgradeNotificationsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotificationListResponse); err != nil {
		return RedisListUpgradeNotificationsResponse{}, err
	}
	return result, nil
}

// listUpgradeNotificationsHandleError handles the ListUpgradeNotifications error response.
func (client *RedisClient) listUpgradeNotificationsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// RegenerateKey - Regenerate Redis cache's access keys. This operation requires write permission to the cache resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) RegenerateKey(ctx context.Context, resourceGroupName string, name string, parameters RedisRegenerateKeyParameters, options *RedisRegenerateKeyOptions) (RedisRegenerateKeyResponse, error) {
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return RedisRegenerateKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RedisRegenerateKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RedisRegenerateKeyResponse{}, client.regenerateKeyHandleError(resp)
	}
	return client.regenerateKeyHandleResponse(resp)
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *RedisClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, name string, parameters RedisRegenerateKeyParameters, options *RedisRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/regenerateKey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// regenerateKeyHandleResponse handles the RegenerateKey response.
func (client *RedisClient) regenerateKeyHandleResponse(resp *http.Response) (RedisRegenerateKeyResponse, error) {
	result := RedisRegenerateKeyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisAccessKeys); err != nil {
		return RedisRegenerateKeyResponse{}, err
	}
	return result, nil
}

// regenerateKeyHandleError handles the RegenerateKey error response.
func (client *RedisClient) regenerateKeyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Update an existing Redis cache.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RedisClient) Update(ctx context.Context, resourceGroupName string, name string, parameters RedisUpdateParameters, options *RedisUpdateOptions) (RedisUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, name, parameters, options)
	if err != nil {
		return RedisUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RedisUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RedisUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *RedisClient) updateCreateRequest(ctx context.Context, resourceGroupName string, name string, parameters RedisUpdateParameters, options *RedisUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *RedisClient) updateHandleResponse(resp *http.Response) (RedisUpdateResponse, error) {
	result := RedisUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RedisResource); err != nil {
		return RedisUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *RedisClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
