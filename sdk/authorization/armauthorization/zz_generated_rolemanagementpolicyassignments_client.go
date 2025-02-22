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

// RoleManagementPolicyAssignmentsClient contains the methods for the RoleManagementPolicyAssignments group.
// Don't use this type directly, use NewRoleManagementPolicyAssignmentsClient() instead.
type RoleManagementPolicyAssignmentsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewRoleManagementPolicyAssignmentsClient creates a new instance of RoleManagementPolicyAssignmentsClient with the specified values.
func NewRoleManagementPolicyAssignmentsClient(con *arm.Connection) *RoleManagementPolicyAssignmentsClient {
	return &RoleManagementPolicyAssignmentsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// Create - Create a role management policy assignment
// If the operation fails it returns the *CloudError error type.
func (client *RoleManagementPolicyAssignmentsClient) Create(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, parameters RoleManagementPolicyAssignment, options *RoleManagementPolicyAssignmentsCreateOptions) (RoleManagementPolicyAssignmentsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, scope, roleManagementPolicyAssignmentName, parameters, options)
	if err != nil {
		return RoleManagementPolicyAssignmentsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleManagementPolicyAssignmentsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return RoleManagementPolicyAssignmentsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *RoleManagementPolicyAssignmentsClient) createCreateRequest(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, parameters RoleManagementPolicyAssignment, options *RoleManagementPolicyAssignmentsCreateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments/{roleManagementPolicyAssignmentName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyAssignmentName == "" {
		return nil, errors.New("parameter roleManagementPolicyAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyAssignmentName}", url.PathEscape(roleManagementPolicyAssignmentName))
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
func (client *RoleManagementPolicyAssignmentsClient) createHandleResponse(resp *http.Response) (RoleManagementPolicyAssignmentsCreateResponse, error) {
	result := RoleManagementPolicyAssignmentsCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicyAssignment); err != nil {
		return RoleManagementPolicyAssignmentsCreateResponse{}, err
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *RoleManagementPolicyAssignmentsClient) createHandleError(resp *http.Response) error {
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

// Delete - Delete a role management policy assignment
// If the operation fails it returns the *CloudError error type.
func (client *RoleManagementPolicyAssignmentsClient) Delete(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, options *RoleManagementPolicyAssignmentsDeleteOptions) (RoleManagementPolicyAssignmentsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, scope, roleManagementPolicyAssignmentName, options)
	if err != nil {
		return RoleManagementPolicyAssignmentsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleManagementPolicyAssignmentsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RoleManagementPolicyAssignmentsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return RoleManagementPolicyAssignmentsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RoleManagementPolicyAssignmentsClient) deleteCreateRequest(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, options *RoleManagementPolicyAssignmentsDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments/{roleManagementPolicyAssignmentName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyAssignmentName == "" {
		return nil, errors.New("parameter roleManagementPolicyAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyAssignmentName}", url.PathEscape(roleManagementPolicyAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *RoleManagementPolicyAssignmentsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get the specified role management policy assignment for a resource scope
// If the operation fails it returns the *CloudError error type.
func (client *RoleManagementPolicyAssignmentsClient) Get(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, options *RoleManagementPolicyAssignmentsGetOptions) (RoleManagementPolicyAssignmentsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleManagementPolicyAssignmentName, options)
	if err != nil {
		return RoleManagementPolicyAssignmentsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleManagementPolicyAssignmentsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleManagementPolicyAssignmentsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleManagementPolicyAssignmentsClient) getCreateRequest(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, options *RoleManagementPolicyAssignmentsGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments/{roleManagementPolicyAssignmentName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyAssignmentName == "" {
		return nil, errors.New("parameter roleManagementPolicyAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyAssignmentName}", url.PathEscape(roleManagementPolicyAssignmentName))
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
func (client *RoleManagementPolicyAssignmentsClient) getHandleResponse(resp *http.Response) (RoleManagementPolicyAssignmentsGetResponse, error) {
	result := RoleManagementPolicyAssignmentsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicyAssignment); err != nil {
		return RoleManagementPolicyAssignmentsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RoleManagementPolicyAssignmentsClient) getHandleError(resp *http.Response) error {
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

// ListForScope - Gets role management assignment policies for a resource scope.
// If the operation fails it returns the *CloudError error type.
func (client *RoleManagementPolicyAssignmentsClient) ListForScope(scope string, options *RoleManagementPolicyAssignmentsListForScopeOptions) *RoleManagementPolicyAssignmentsListForScopePager {
	return &RoleManagementPolicyAssignmentsListForScopePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForScopeCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp RoleManagementPolicyAssignmentsListForScopeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RoleManagementPolicyAssignmentListResult.NextLink)
		},
	}
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleManagementPolicyAssignmentsClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleManagementPolicyAssignmentsListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
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

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleManagementPolicyAssignmentsClient) listForScopeHandleResponse(resp *http.Response) (RoleManagementPolicyAssignmentsListForScopeResponse, error) {
	result := RoleManagementPolicyAssignmentsListForScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicyAssignmentListResult); err != nil {
		return RoleManagementPolicyAssignmentsListForScopeResponse{}, err
	}
	return result, nil
}

// listForScopeHandleError handles the ListForScope error response.
func (client *RoleManagementPolicyAssignmentsClient) listForScopeHandleError(resp *http.Response) error {
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
