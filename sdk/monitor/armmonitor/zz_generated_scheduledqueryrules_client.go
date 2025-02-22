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
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ScheduledQueryRulesClient contains the methods for the ScheduledQueryRules group.
// Don't use this type directly, use NewScheduledQueryRulesClient() instead.
type ScheduledQueryRulesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewScheduledQueryRulesClient creates a new instance of ScheduledQueryRulesClient with the specified values.
func NewScheduledQueryRulesClient(con *arm.Connection, subscriptionID string) *ScheduledQueryRulesClient {
	return &ScheduledQueryRulesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates an log search rule.
// If the operation fails it returns the *ErrorContract error type.
func (client *ScheduledQueryRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters LogSearchRuleResource, options *ScheduledQueryRulesCreateOrUpdateOptions) (ScheduledQueryRulesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ruleName, parameters, options)
	if err != nil {
		return ScheduledQueryRulesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScheduledQueryRulesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ScheduledQueryRulesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ScheduledQueryRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, parameters LogSearchRuleResource, options *ScheduledQueryRulesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Insights/scheduledQueryRules/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-04-16")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ScheduledQueryRulesClient) createOrUpdateHandleResponse(resp *http.Response) (ScheduledQueryRulesCreateOrUpdateResponse, error) {
	result := ScheduledQueryRulesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogSearchRuleResource); err != nil {
		return ScheduledQueryRulesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ScheduledQueryRulesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorContract{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes a Log Search rule
// If the operation fails it returns the *ErrorContract error type.
func (client *ScheduledQueryRulesClient) Delete(ctx context.Context, resourceGroupName string, ruleName string, options *ScheduledQueryRulesDeleteOptions) (ScheduledQueryRulesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return ScheduledQueryRulesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScheduledQueryRulesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ScheduledQueryRulesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ScheduledQueryRulesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ScheduledQueryRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *ScheduledQueryRulesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Insights/scheduledQueryRules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-04-16")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ScheduledQueryRulesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorContract{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets an Log Search rule
// If the operation fails it returns the *ErrorContract error type.
func (client *ScheduledQueryRulesClient) Get(ctx context.Context, resourceGroupName string, ruleName string, options *ScheduledQueryRulesGetOptions) (ScheduledQueryRulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return ScheduledQueryRulesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScheduledQueryRulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScheduledQueryRulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ScheduledQueryRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *ScheduledQueryRulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Insights/scheduledQueryRules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-04-16")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ScheduledQueryRulesClient) getHandleResponse(resp *http.Response) (ScheduledQueryRulesGetResponse, error) {
	result := ScheduledQueryRulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogSearchRuleResource); err != nil {
		return ScheduledQueryRulesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ScheduledQueryRulesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorContract{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - List the Log Search rules within a resource group.
// If the operation fails it returns the *ErrorContract error type.
func (client *ScheduledQueryRulesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *ScheduledQueryRulesListByResourceGroupOptions) (ScheduledQueryRulesListByResourceGroupResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return ScheduledQueryRulesListByResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScheduledQueryRulesListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScheduledQueryRulesListByResourceGroupResponse{}, client.listByResourceGroupHandleError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ScheduledQueryRulesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ScheduledQueryRulesListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Insights/scheduledQueryRules"
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
	reqQP.Set("api-version", "2018-04-16")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ScheduledQueryRulesClient) listByResourceGroupHandleResponse(resp *http.Response) (ScheduledQueryRulesListByResourceGroupResponse, error) {
	result := ScheduledQueryRulesListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogSearchRuleResourceCollection); err != nil {
		return ScheduledQueryRulesListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ScheduledQueryRulesClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorContract{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - List the Log Search rules within a subscription group.
// If the operation fails it returns the *ErrorContract error type.
func (client *ScheduledQueryRulesClient) ListBySubscription(ctx context.Context, options *ScheduledQueryRulesListBySubscriptionOptions) (ScheduledQueryRulesListBySubscriptionResponse, error) {
	req, err := client.listBySubscriptionCreateRequest(ctx, options)
	if err != nil {
		return ScheduledQueryRulesListBySubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScheduledQueryRulesListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScheduledQueryRulesListBySubscriptionResponse{}, client.listBySubscriptionHandleError(resp)
	}
	return client.listBySubscriptionHandleResponse(resp)
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ScheduledQueryRulesClient) listBySubscriptionCreateRequest(ctx context.Context, options *ScheduledQueryRulesListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/scheduledQueryRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-04-16")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ScheduledQueryRulesClient) listBySubscriptionHandleResponse(resp *http.Response) (ScheduledQueryRulesListBySubscriptionResponse, error) {
	result := ScheduledQueryRulesListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogSearchRuleResourceCollection); err != nil {
		return ScheduledQueryRulesListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *ScheduledQueryRulesClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorContract{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Update log search Rule.
// If the operation fails it returns the *ErrorContract error type.
func (client *ScheduledQueryRulesClient) Update(ctx context.Context, resourceGroupName string, ruleName string, parameters LogSearchRuleResourcePatch, options *ScheduledQueryRulesUpdateOptions) (ScheduledQueryRulesUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, ruleName, parameters, options)
	if err != nil {
		return ScheduledQueryRulesUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScheduledQueryRulesUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScheduledQueryRulesUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ScheduledQueryRulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, parameters LogSearchRuleResourcePatch, options *ScheduledQueryRulesUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Insights/scheduledQueryRules/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-04-16")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ScheduledQueryRulesClient) updateHandleResponse(resp *http.Response) (ScheduledQueryRulesUpdateResponse, error) {
	result := ScheduledQueryRulesUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogSearchRuleResource); err != nil {
		return ScheduledQueryRulesUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *ScheduledQueryRulesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorContract{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
