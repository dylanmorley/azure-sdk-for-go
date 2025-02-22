//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// RoleEligibilityScheduleRequestsClient contains the methods for the RoleEligibilityScheduleRequests group.
// Don't use this type directly, use NewRoleEligibilityScheduleRequestsClient() instead.
type RoleEligibilityScheduleRequestsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewRoleEligibilityScheduleRequestsClient creates a new instance of RoleEligibilityScheduleRequestsClient with the specified values.
func NewRoleEligibilityScheduleRequestsClient(con *arm.Connection) *RoleEligibilityScheduleRequestsClient {
	return &RoleEligibilityScheduleRequestsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// Cancel - Cancels a pending role eligibility schedule request.
// If the operation fails it returns the *CloudError error type.
func (client *RoleEligibilityScheduleRequestsClient) Cancel(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, options *RoleEligibilityScheduleRequestsCancelOptions) (RoleEligibilityScheduleRequestsCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, scope, roleEligibilityScheduleRequestName, options)
	if err != nil {
		return RoleEligibilityScheduleRequestsCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleEligibilityScheduleRequestsCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleEligibilityScheduleRequestsCancelResponse{}, client.cancelHandleError(resp)
	}
	return RoleEligibilityScheduleRequestsCancelResponse{RawResponse: resp}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *RoleEligibilityScheduleRequestsClient) cancelCreateRequest(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, options *RoleEligibilityScheduleRequestsCancelOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{roleEligibilityScheduleRequestName}/cancel"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleEligibilityScheduleRequestName == "" {
		return nil, errors.New("parameter roleEligibilityScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleEligibilityScheduleRequestName}", url.PathEscape(roleEligibilityScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// cancelHandleError handles the Cancel error response.
func (client *RoleEligibilityScheduleRequestsClient) cancelHandleError(resp *http.Response) error {
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

// Create - Creates a role eligibility schedule request.
// If the operation fails it returns the *CloudError error type.
func (client *RoleEligibilityScheduleRequestsClient) Create(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, parameters RoleEligibilityScheduleRequest, options *RoleEligibilityScheduleRequestsCreateOptions) (RoleEligibilityScheduleRequestsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, scope, roleEligibilityScheduleRequestName, parameters, options)
	if err != nil {
		return RoleEligibilityScheduleRequestsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleEligibilityScheduleRequestsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return RoleEligibilityScheduleRequestsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *RoleEligibilityScheduleRequestsClient) createCreateRequest(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, parameters RoleEligibilityScheduleRequest, options *RoleEligibilityScheduleRequestsCreateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{roleEligibilityScheduleRequestName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleEligibilityScheduleRequestName == "" {
		return nil, errors.New("parameter roleEligibilityScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleEligibilityScheduleRequestName}", url.PathEscape(roleEligibilityScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *RoleEligibilityScheduleRequestsClient) createHandleResponse(resp *http.Response) (RoleEligibilityScheduleRequestsCreateResponse, error) {
	result := RoleEligibilityScheduleRequestsCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleEligibilityScheduleRequest); err != nil {
		return RoleEligibilityScheduleRequestsCreateResponse{}, err
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *RoleEligibilityScheduleRequestsClient) createHandleError(resp *http.Response) error {
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

// Get - Get the specified role eligibility schedule request.
// If the operation fails it returns the *CloudError error type.
func (client *RoleEligibilityScheduleRequestsClient) Get(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, options *RoleEligibilityScheduleRequestsGetOptions) (RoleEligibilityScheduleRequestsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleEligibilityScheduleRequestName, options)
	if err != nil {
		return RoleEligibilityScheduleRequestsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleEligibilityScheduleRequestsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleEligibilityScheduleRequestsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleEligibilityScheduleRequestsClient) getCreateRequest(ctx context.Context, scope string, roleEligibilityScheduleRequestName string, options *RoleEligibilityScheduleRequestsGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{roleEligibilityScheduleRequestName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleEligibilityScheduleRequestName == "" {
		return nil, errors.New("parameter roleEligibilityScheduleRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleEligibilityScheduleRequestName}", url.PathEscape(roleEligibilityScheduleRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleEligibilityScheduleRequestsClient) getHandleResponse(resp *http.Response) (RoleEligibilityScheduleRequestsGetResponse, error) {
	result := RoleEligibilityScheduleRequestsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleEligibilityScheduleRequest); err != nil {
		return RoleEligibilityScheduleRequestsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RoleEligibilityScheduleRequestsClient) getHandleError(resp *http.Response) error {
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

// ListForScope - Gets role eligibility schedule requests for a scope.
// If the operation fails it returns the *CloudError error type.
func (client *RoleEligibilityScheduleRequestsClient) ListForScope(scope string, options *RoleEligibilityScheduleRequestsListForScopeOptions) *RoleEligibilityScheduleRequestsListForScopePager {
	return &RoleEligibilityScheduleRequestsListForScopePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForScopeCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp RoleEligibilityScheduleRequestsListForScopeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RoleEligibilityScheduleRequestListResult.NextLink)
		},
	}
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleEligibilityScheduleRequestsClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleEligibilityScheduleRequestsListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleEligibilityScheduleRequestsClient) listForScopeHandleResponse(resp *http.Response) (RoleEligibilityScheduleRequestsListForScopeResponse, error) {
	result := RoleEligibilityScheduleRequestsListForScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleEligibilityScheduleRequestListResult); err != nil {
		return RoleEligibilityScheduleRequestsListForScopeResponse{}, err
	}
	return result, nil
}

// listForScopeHandleError handles the ListForScope error response.
func (client *RoleEligibilityScheduleRequestsClient) listForScopeHandleError(resp *http.Response) error {
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
