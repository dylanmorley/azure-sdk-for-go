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

// MetricAlertsStatusClient contains the methods for the MetricAlertsStatus group.
// Don't use this type directly, use NewMetricAlertsStatusClient() instead.
type MetricAlertsStatusClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewMetricAlertsStatusClient creates a new instance of MetricAlertsStatusClient with the specified values.
func NewMetricAlertsStatusClient(con *arm.Connection, subscriptionID string) *MetricAlertsStatusClient {
	return &MetricAlertsStatusClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - Retrieve an alert rule status.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MetricAlertsStatusClient) List(ctx context.Context, resourceGroupName string, ruleName string, options *MetricAlertsStatusListOptions) (MetricAlertsStatusListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return MetricAlertsStatusListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MetricAlertsStatusListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MetricAlertsStatusListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *MetricAlertsStatusClient) listCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *MetricAlertsStatusListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{ruleName}/status"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MetricAlertsStatusClient) listHandleResponse(resp *http.Response) (MetricAlertsStatusListResponse, error) {
	result := MetricAlertsStatusListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricAlertStatusCollection); err != nil {
		return MetricAlertsStatusListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *MetricAlertsStatusClient) listHandleError(resp *http.Response) error {
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

// ListByName - Retrieve an alert rule status.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MetricAlertsStatusClient) ListByName(ctx context.Context, resourceGroupName string, ruleName string, statusName string, options *MetricAlertsStatusListByNameOptions) (MetricAlertsStatusListByNameResponse, error) {
	req, err := client.listByNameCreateRequest(ctx, resourceGroupName, ruleName, statusName, options)
	if err != nil {
		return MetricAlertsStatusListByNameResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MetricAlertsStatusListByNameResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MetricAlertsStatusListByNameResponse{}, client.listByNameHandleError(resp)
	}
	return client.listByNameHandleResponse(resp)
}

// listByNameCreateRequest creates the ListByName request.
func (client *MetricAlertsStatusClient) listByNameCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, statusName string, options *MetricAlertsStatusListByNameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{ruleName}/status/{statusName}"
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
	if statusName == "" {
		return nil, errors.New("parameter statusName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{statusName}", url.PathEscape(statusName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByNameHandleResponse handles the ListByName response.
func (client *MetricAlertsStatusClient) listByNameHandleResponse(resp *http.Response) (MetricAlertsStatusListByNameResponse, error) {
	result := MetricAlertsStatusListByNameResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricAlertStatusCollection); err != nil {
		return MetricAlertsStatusListByNameResponse{}, err
	}
	return result, nil
}

// listByNameHandleError handles the ListByName error response.
func (client *MetricAlertsStatusClient) listByNameHandleError(resp *http.Response) error {
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
