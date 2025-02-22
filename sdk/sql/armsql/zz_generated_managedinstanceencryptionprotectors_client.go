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

// ManagedInstanceEncryptionProtectorsClient contains the methods for the ManagedInstanceEncryptionProtectors group.
// Don't use this type directly, use NewManagedInstanceEncryptionProtectorsClient() instead.
type ManagedInstanceEncryptionProtectorsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewManagedInstanceEncryptionProtectorsClient creates a new instance of ManagedInstanceEncryptionProtectorsClient with the specified values.
func NewManagedInstanceEncryptionProtectorsClient(con *arm.Connection, subscriptionID string) *ManagedInstanceEncryptionProtectorsClient {
	return &ManagedInstanceEncryptionProtectorsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Updates an existing encryption protector.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceEncryptionProtectorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, parameters ManagedInstanceEncryptionProtector, options *ManagedInstanceEncryptionProtectorsBeginCreateOrUpdateOptions) (ManagedInstanceEncryptionProtectorsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, encryptionProtectorName, parameters, options)
	if err != nil {
		return ManagedInstanceEncryptionProtectorsCreateOrUpdatePollerResponse{}, err
	}
	result := ManagedInstanceEncryptionProtectorsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ManagedInstanceEncryptionProtectorsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ManagedInstanceEncryptionProtectorsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ManagedInstanceEncryptionProtectorsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Updates an existing encryption protector.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceEncryptionProtectorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, parameters ManagedInstanceEncryptionProtector, options *ManagedInstanceEncryptionProtectorsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, encryptionProtectorName, parameters, options)
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
func (client *ManagedInstanceEncryptionProtectorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, parameters ManagedInstanceEncryptionProtector, options *ManagedInstanceEncryptionProtectorsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/encryptionProtector/{encryptionProtectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
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
func (client *ManagedInstanceEncryptionProtectorsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a managed instance encryption protector.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceEncryptionProtectorsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, options *ManagedInstanceEncryptionProtectorsGetOptions) (ManagedInstanceEncryptionProtectorsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, encryptionProtectorName, options)
	if err != nil {
		return ManagedInstanceEncryptionProtectorsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedInstanceEncryptionProtectorsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedInstanceEncryptionProtectorsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedInstanceEncryptionProtectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, options *ManagedInstanceEncryptionProtectorsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/encryptionProtector/{encryptionProtectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
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
func (client *ManagedInstanceEncryptionProtectorsClient) getHandleResponse(resp *http.Response) (ManagedInstanceEncryptionProtectorsGetResponse, error) {
	result := ManagedInstanceEncryptionProtectorsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceEncryptionProtector); err != nil {
		return ManagedInstanceEncryptionProtectorsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ManagedInstanceEncryptionProtectorsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByInstance - Gets a list of managed instance encryption protectors
// If the operation fails it returns a generic error.
func (client *ManagedInstanceEncryptionProtectorsClient) ListByInstance(resourceGroupName string, managedInstanceName string, options *ManagedInstanceEncryptionProtectorsListByInstanceOptions) *ManagedInstanceEncryptionProtectorsListByInstancePager {
	return &ManagedInstanceEncryptionProtectorsListByInstancePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
		},
		advancer: func(ctx context.Context, resp ManagedInstanceEncryptionProtectorsListByInstanceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ManagedInstanceEncryptionProtectorListResult.NextLink)
		},
	}
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *ManagedInstanceEncryptionProtectorsClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstanceEncryptionProtectorsListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/encryptionProtector"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
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

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *ManagedInstanceEncryptionProtectorsClient) listByInstanceHandleResponse(resp *http.Response) (ManagedInstanceEncryptionProtectorsListByInstanceResponse, error) {
	result := ManagedInstanceEncryptionProtectorsListByInstanceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceEncryptionProtectorListResult); err != nil {
		return ManagedInstanceEncryptionProtectorsListByInstanceResponse{}, err
	}
	return result, nil
}

// listByInstanceHandleError handles the ListByInstance error response.
func (client *ManagedInstanceEncryptionProtectorsClient) listByInstanceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginRevalidate - Revalidates an existing encryption protector.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceEncryptionProtectorsClient) BeginRevalidate(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, options *ManagedInstanceEncryptionProtectorsBeginRevalidateOptions) (ManagedInstanceEncryptionProtectorsRevalidatePollerResponse, error) {
	resp, err := client.revalidate(ctx, resourceGroupName, managedInstanceName, encryptionProtectorName, options)
	if err != nil {
		return ManagedInstanceEncryptionProtectorsRevalidatePollerResponse{}, err
	}
	result := ManagedInstanceEncryptionProtectorsRevalidatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ManagedInstanceEncryptionProtectorsClient.Revalidate", "", resp, client.pl, client.revalidateHandleError)
	if err != nil {
		return ManagedInstanceEncryptionProtectorsRevalidatePollerResponse{}, err
	}
	result.Poller = &ManagedInstanceEncryptionProtectorsRevalidatePoller{
		pt: pt,
	}
	return result, nil
}

// Revalidate - Revalidates an existing encryption protector.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceEncryptionProtectorsClient) revalidate(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, options *ManagedInstanceEncryptionProtectorsBeginRevalidateOptions) (*http.Response, error) {
	req, err := client.revalidateCreateRequest(ctx, resourceGroupName, managedInstanceName, encryptionProtectorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.revalidateHandleError(resp)
	}
	return resp, nil
}

// revalidateCreateRequest creates the Revalidate request.
func (client *ManagedInstanceEncryptionProtectorsClient) revalidateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, encryptionProtectorName EncryptionProtectorName, options *ManagedInstanceEncryptionProtectorsBeginRevalidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/encryptionProtector/{encryptionProtectorName}/revalidate"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
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

// revalidateHandleError handles the Revalidate error response.
func (client *ManagedInstanceEncryptionProtectorsClient) revalidateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
