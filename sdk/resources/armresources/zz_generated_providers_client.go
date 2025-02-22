//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

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

// ProvidersClient contains the methods for the Providers group.
// Don't use this type directly, use NewProvidersClient() instead.
type ProvidersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewProvidersClient creates a new instance of ProvidersClient with the specified values.
func NewProvidersClient(con *arm.Connection, subscriptionID string) *ProvidersClient {
	return &ProvidersClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Gets the specified resource provider.
// If the operation fails it returns the *CloudError error type.
func (client *ProvidersClient) Get(ctx context.Context, resourceProviderNamespace string, options *ProvidersGetOptions) (ProvidersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceProviderNamespace, options)
	if err != nil {
		return ProvidersGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProvidersGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProvidersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProvidersClient) getCreateRequest(ctx context.Context, resourceProviderNamespace string, options *ProvidersGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProvidersClient) getHandleResponse(resp *http.Response) (ProvidersGetResponse, error) {
	result := ProvidersGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Provider); err != nil {
		return ProvidersGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ProvidersClient) getHandleError(resp *http.Response) error {
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

// GetAtTenantScope - Gets the specified resource provider at the tenant level.
// If the operation fails it returns the *CloudError error type.
func (client *ProvidersClient) GetAtTenantScope(ctx context.Context, resourceProviderNamespace string, options *ProvidersGetAtTenantScopeOptions) (ProvidersGetAtTenantScopeResponse, error) {
	req, err := client.getAtTenantScopeCreateRequest(ctx, resourceProviderNamespace, options)
	if err != nil {
		return ProvidersGetAtTenantScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProvidersGetAtTenantScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProvidersGetAtTenantScopeResponse{}, client.getAtTenantScopeHandleError(resp)
	}
	return client.getAtTenantScopeHandleResponse(resp)
}

// getAtTenantScopeCreateRequest creates the GetAtTenantScope request.
func (client *ProvidersClient) getAtTenantScopeCreateRequest(ctx context.Context, resourceProviderNamespace string, options *ProvidersGetAtTenantScopeOptions) (*policy.Request, error) {
	urlPath := "/providers/{resourceProviderNamespace}"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAtTenantScopeHandleResponse handles the GetAtTenantScope response.
func (client *ProvidersClient) getAtTenantScopeHandleResponse(resp *http.Response) (ProvidersGetAtTenantScopeResponse, error) {
	result := ProvidersGetAtTenantScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Provider); err != nil {
		return ProvidersGetAtTenantScopeResponse{}, err
	}
	return result, nil
}

// getAtTenantScopeHandleError handles the GetAtTenantScope error response.
func (client *ProvidersClient) getAtTenantScopeHandleError(resp *http.Response) error {
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

// List - Gets all resource providers for a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *ProvidersClient) List(options *ProvidersListOptions) *ProvidersListPager {
	return &ProvidersListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ProvidersListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ProviderListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ProvidersClient) listCreateRequest(ctx context.Context, options *ProvidersListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProvidersClient) listHandleResponse(resp *http.Response) (ProvidersListResponse, error) {
	result := ProvidersListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderListResult); err != nil {
		return ProvidersListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ProvidersClient) listHandleError(resp *http.Response) error {
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

// ListAtTenantScope - Gets all resource providers for the tenant.
// If the operation fails it returns the *CloudError error type.
func (client *ProvidersClient) ListAtTenantScope(options *ProvidersListAtTenantScopeOptions) *ProvidersListAtTenantScopePager {
	return &ProvidersListAtTenantScopePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAtTenantScopeCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ProvidersListAtTenantScopeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ProviderListResult.NextLink)
		},
	}
}

// listAtTenantScopeCreateRequest creates the ListAtTenantScope request.
func (client *ProvidersClient) listAtTenantScopeCreateRequest(ctx context.Context, options *ProvidersListAtTenantScopeOptions) (*policy.Request, error) {
	urlPath := "/providers"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAtTenantScopeHandleResponse handles the ListAtTenantScope response.
func (client *ProvidersClient) listAtTenantScopeHandleResponse(resp *http.Response) (ProvidersListAtTenantScopeResponse, error) {
	result := ProvidersListAtTenantScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderListResult); err != nil {
		return ProvidersListAtTenantScopeResponse{}, err
	}
	return result, nil
}

// listAtTenantScopeHandleError handles the ListAtTenantScope error response.
func (client *ProvidersClient) listAtTenantScopeHandleError(resp *http.Response) error {
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

// ProviderPermissions - Get the provider permissions.
// If the operation fails it returns the *CloudError error type.
func (client *ProvidersClient) ProviderPermissions(ctx context.Context, resourceProviderNamespace string, options *ProvidersProviderPermissionsOptions) (ProvidersProviderPermissionsResponse, error) {
	req, err := client.providerPermissionsCreateRequest(ctx, resourceProviderNamespace, options)
	if err != nil {
		return ProvidersProviderPermissionsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProvidersProviderPermissionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProvidersProviderPermissionsResponse{}, client.providerPermissionsHandleError(resp)
	}
	return client.providerPermissionsHandleResponse(resp)
}

// providerPermissionsCreateRequest creates the ProviderPermissions request.
func (client *ProvidersClient) providerPermissionsCreateRequest(ctx context.Context, resourceProviderNamespace string, options *ProvidersProviderPermissionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/providerPermissions"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// providerPermissionsHandleResponse handles the ProviderPermissions response.
func (client *ProvidersClient) providerPermissionsHandleResponse(resp *http.Response) (ProvidersProviderPermissionsResponse, error) {
	result := ProvidersProviderPermissionsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderPermissionListResult); err != nil {
		return ProvidersProviderPermissionsResponse{}, err
	}
	return result, nil
}

// providerPermissionsHandleError handles the ProviderPermissions error response.
func (client *ProvidersClient) providerPermissionsHandleError(resp *http.Response) error {
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

// Register - Registers a subscription with a resource provider.
// If the operation fails it returns the *CloudError error type.
func (client *ProvidersClient) Register(ctx context.Context, resourceProviderNamespace string, options *ProvidersRegisterOptions) (ProvidersRegisterResponse, error) {
	req, err := client.registerCreateRequest(ctx, resourceProviderNamespace, options)
	if err != nil {
		return ProvidersRegisterResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProvidersRegisterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProvidersRegisterResponse{}, client.registerHandleError(resp)
	}
	return client.registerHandleResponse(resp)
}

// registerCreateRequest creates the Register request.
func (client *ProvidersClient) registerCreateRequest(ctx context.Context, resourceProviderNamespace string, options *ProvidersRegisterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/register"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Properties != nil {
		return req, runtime.MarshalAsJSON(req, *options.Properties)
	}
	return req, nil
}

// registerHandleResponse handles the Register response.
func (client *ProvidersClient) registerHandleResponse(resp *http.Response) (ProvidersRegisterResponse, error) {
	result := ProvidersRegisterResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Provider); err != nil {
		return ProvidersRegisterResponse{}, err
	}
	return result, nil
}

// registerHandleError handles the Register error response.
func (client *ProvidersClient) registerHandleError(resp *http.Response) error {
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

// RegisterAtManagementGroupScope - Registers a management group with a resource provider.
// If the operation fails it returns the *CloudError error type.
func (client *ProvidersClient) RegisterAtManagementGroupScope(ctx context.Context, resourceProviderNamespace string, groupID string, options *ProvidersRegisterAtManagementGroupScopeOptions) (ProvidersRegisterAtManagementGroupScopeResponse, error) {
	req, err := client.registerAtManagementGroupScopeCreateRequest(ctx, resourceProviderNamespace, groupID, options)
	if err != nil {
		return ProvidersRegisterAtManagementGroupScopeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProvidersRegisterAtManagementGroupScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProvidersRegisterAtManagementGroupScopeResponse{}, client.registerAtManagementGroupScopeHandleError(resp)
	}
	return ProvidersRegisterAtManagementGroupScopeResponse{RawResponse: resp}, nil
}

// registerAtManagementGroupScopeCreateRequest creates the RegisterAtManagementGroupScope request.
func (client *ProvidersClient) registerAtManagementGroupScopeCreateRequest(ctx context.Context, resourceProviderNamespace string, groupID string, options *ProvidersRegisterAtManagementGroupScopeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/providers/{resourceProviderNamespace}/register"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// registerAtManagementGroupScopeHandleError handles the RegisterAtManagementGroupScope error response.
func (client *ProvidersClient) registerAtManagementGroupScopeHandleError(resp *http.Response) error {
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

// Unregister - Unregisters a subscription from a resource provider.
// If the operation fails it returns the *CloudError error type.
func (client *ProvidersClient) Unregister(ctx context.Context, resourceProviderNamespace string, options *ProvidersUnregisterOptions) (ProvidersUnregisterResponse, error) {
	req, err := client.unregisterCreateRequest(ctx, resourceProviderNamespace, options)
	if err != nil {
		return ProvidersUnregisterResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProvidersUnregisterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProvidersUnregisterResponse{}, client.unregisterHandleError(resp)
	}
	return client.unregisterHandleResponse(resp)
}

// unregisterCreateRequest creates the Unregister request.
func (client *ProvidersClient) unregisterCreateRequest(ctx context.Context, resourceProviderNamespace string, options *ProvidersUnregisterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/unregister"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// unregisterHandleResponse handles the Unregister response.
func (client *ProvidersClient) unregisterHandleResponse(resp *http.Response) (ProvidersUnregisterResponse, error) {
	result := ProvidersUnregisterResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Provider); err != nil {
		return ProvidersUnregisterResponse{}, err
	}
	return result, nil
}

// unregisterHandleError handles the Unregister error response.
func (client *ProvidersClient) unregisterHandleError(resp *http.Response) error {
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
