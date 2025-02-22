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

// ManagedInstanceLongTermRetentionPoliciesClient contains the methods for the ManagedInstanceLongTermRetentionPolicies group.
// Don't use this type directly, use NewManagedInstanceLongTermRetentionPoliciesClient() instead.
type ManagedInstanceLongTermRetentionPoliciesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewManagedInstanceLongTermRetentionPoliciesClient creates a new instance of ManagedInstanceLongTermRetentionPoliciesClient with the specified values.
func NewManagedInstanceLongTermRetentionPoliciesClient(con *arm.Connection, subscriptionID string) *ManagedInstanceLongTermRetentionPoliciesClient {
	return &ManagedInstanceLongTermRetentionPoliciesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Sets a managed database's long term retention policy.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedInstanceLongTermRetentionPolicyName, parameters ManagedInstanceLongTermRetentionPolicy, options *ManagedInstanceLongTermRetentionPoliciesBeginCreateOrUpdateOptions) (ManagedInstanceLongTermRetentionPoliciesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, parameters, options)
	if err != nil {
		return ManagedInstanceLongTermRetentionPoliciesCreateOrUpdatePollerResponse{}, err
	}
	result := ManagedInstanceLongTermRetentionPoliciesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ManagedInstanceLongTermRetentionPoliciesClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ManagedInstanceLongTermRetentionPoliciesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ManagedInstanceLongTermRetentionPoliciesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Sets a managed database's long term retention policy.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedInstanceLongTermRetentionPolicyName, parameters ManagedInstanceLongTermRetentionPolicy, options *ManagedInstanceLongTermRetentionPoliciesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, parameters, options)
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
func (client *ManagedInstanceLongTermRetentionPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedInstanceLongTermRetentionPolicyName, parameters ManagedInstanceLongTermRetentionPolicy, options *ManagedInstanceLongTermRetentionPoliciesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/backupLongTermRetentionPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(string(policyName)))
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
func (client *ManagedInstanceLongTermRetentionPoliciesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a managed database's long term retention policy.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedInstanceLongTermRetentionPolicyName, options *ManagedInstanceLongTermRetentionPoliciesGetOptions) (ManagedInstanceLongTermRetentionPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, policyName, options)
	if err != nil {
		return ManagedInstanceLongTermRetentionPoliciesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedInstanceLongTermRetentionPoliciesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedInstanceLongTermRetentionPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName ManagedInstanceLongTermRetentionPolicyName, options *ManagedInstanceLongTermRetentionPoliciesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/backupLongTermRetentionPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(string(policyName)))
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
func (client *ManagedInstanceLongTermRetentionPoliciesClient) getHandleResponse(resp *http.Response) (ManagedInstanceLongTermRetentionPoliciesGetResponse, error) {
	result := ManagedInstanceLongTermRetentionPoliciesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceLongTermRetentionPolicy); err != nil {
		return ManagedInstanceLongTermRetentionPoliciesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByDatabase - Gets a database's long term retention policy.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) ListByDatabase(resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedInstanceLongTermRetentionPoliciesListByDatabaseOptions) *ManagedInstanceLongTermRetentionPoliciesListByDatabasePager {
	return &ManagedInstanceLongTermRetentionPoliciesListByDatabasePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDatabaseCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
		},
		advancer: func(ctx context.Context, resp ManagedInstanceLongTermRetentionPoliciesListByDatabaseResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ManagedInstanceLongTermRetentionPolicyListResult.NextLink)
		},
	}
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedInstanceLongTermRetentionPoliciesListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/backupLongTermRetentionPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
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

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) listByDatabaseHandleResponse(resp *http.Response) (ManagedInstanceLongTermRetentionPoliciesListByDatabaseResponse, error) {
	result := ManagedInstanceLongTermRetentionPoliciesListByDatabaseResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceLongTermRetentionPolicyListResult); err != nil {
		return ManagedInstanceLongTermRetentionPoliciesListByDatabaseResponse{}, err
	}
	return result, nil
}

// listByDatabaseHandleError handles the ListByDatabase error response.
func (client *ManagedInstanceLongTermRetentionPoliciesClient) listByDatabaseHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
