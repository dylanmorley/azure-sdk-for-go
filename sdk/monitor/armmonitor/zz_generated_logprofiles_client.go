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

// LogProfilesClient contains the methods for the LogProfiles group.
// Don't use this type directly, use NewLogProfilesClient() instead.
type LogProfilesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewLogProfilesClient creates a new instance of LogProfilesClient with the specified values.
func NewLogProfilesClient(con *arm.Connection, subscriptionID string) *LogProfilesClient {
	return &LogProfilesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create or update a log profile in Azure Monitoring REST API.
// If the operation fails it returns a generic error.
func (client *LogProfilesClient) CreateOrUpdate(ctx context.Context, logProfileName string, parameters LogProfileResource, options *LogProfilesCreateOrUpdateOptions) (LogProfilesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, logProfileName, parameters, options)
	if err != nil {
		return LogProfilesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LogProfilesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogProfilesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LogProfilesClient) createOrUpdateCreateRequest(ctx context.Context, logProfileName string, parameters LogProfileResource, options *LogProfilesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/logprofiles/{logProfileName}"
	if logProfileName == "" {
		return nil, errors.New("parameter logProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logProfileName}", url.PathEscape(logProfileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LogProfilesClient) createOrUpdateHandleResponse(resp *http.Response) (LogProfilesCreateOrUpdateResponse, error) {
	result := LogProfilesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogProfileResource); err != nil {
		return LogProfilesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *LogProfilesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Delete - Deletes the log profile.
// If the operation fails it returns a generic error.
func (client *LogProfilesClient) Delete(ctx context.Context, logProfileName string, options *LogProfilesDeleteOptions) (LogProfilesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, logProfileName, options)
	if err != nil {
		return LogProfilesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LogProfilesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogProfilesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return LogProfilesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LogProfilesClient) deleteCreateRequest(ctx context.Context, logProfileName string, options *LogProfilesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/logprofiles/{logProfileName}"
	if logProfileName == "" {
		return nil, errors.New("parameter logProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logProfileName}", url.PathEscape(logProfileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *LogProfilesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets the log profile.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LogProfilesClient) Get(ctx context.Context, logProfileName string, options *LogProfilesGetOptions) (LogProfilesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, logProfileName, options)
	if err != nil {
		return LogProfilesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LogProfilesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogProfilesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LogProfilesClient) getCreateRequest(ctx context.Context, logProfileName string, options *LogProfilesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/logprofiles/{logProfileName}"
	if logProfileName == "" {
		return nil, errors.New("parameter logProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logProfileName}", url.PathEscape(logProfileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LogProfilesClient) getHandleResponse(resp *http.Response) (LogProfilesGetResponse, error) {
	result := LogProfilesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogProfileResource); err != nil {
		return LogProfilesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *LogProfilesClient) getHandleError(resp *http.Response) error {
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

// List - List the log profiles.
// If the operation fails it returns a generic error.
func (client *LogProfilesClient) List(ctx context.Context, options *LogProfilesListOptions) (LogProfilesListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return LogProfilesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LogProfilesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogProfilesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *LogProfilesClient) listCreateRequest(ctx context.Context, options *LogProfilesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/logprofiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LogProfilesClient) listHandleResponse(resp *http.Response) (LogProfilesListResponse, error) {
	result := LogProfilesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogProfileCollection); err != nil {
		return LogProfilesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *LogProfilesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Update - Updates an existing LogProfilesResource. To update other fields use the CreateOrUpdate method.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LogProfilesClient) Update(ctx context.Context, logProfileName string, logProfilesResource LogProfileResourcePatch, options *LogProfilesUpdateOptions) (LogProfilesUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, logProfileName, logProfilesResource, options)
	if err != nil {
		return LogProfilesUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LogProfilesUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogProfilesUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *LogProfilesClient) updateCreateRequest(ctx context.Context, logProfileName string, logProfilesResource LogProfileResourcePatch, options *LogProfilesUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/logprofiles/{logProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if logProfileName == "" {
		return nil, errors.New("parameter logProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logProfileName}", url.PathEscape(logProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, logProfilesResource)
}

// updateHandleResponse handles the Update response.
func (client *LogProfilesClient) updateHandleResponse(resp *http.Response) (LogProfilesUpdateResponse, error) {
	result := LogProfilesUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogProfileResource); err != nil {
		return LogProfilesUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *LogProfilesClient) updateHandleError(resp *http.Response) error {
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
