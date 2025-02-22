//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// SubscriptionClient contains the methods for the Subscription group.
// Don't use this type directly, use NewSubscriptionClient() instead.
type SubscriptionClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSubscriptionClient creates a new instance of SubscriptionClient with the specified values.
func NewSubscriptionClient(con *arm.Connection, subscriptionID string) *SubscriptionClient {
	return &SubscriptionClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates the subscription of specified user to the specified product.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SubscriptionClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, sid string, parameters SubscriptionCreateParameters, options *SubscriptionCreateOrUpdateOptions) (SubscriptionCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, sid, parameters, options)
	if err != nil {
		return SubscriptionCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SubscriptionCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SubscriptionClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, parameters SubscriptionCreateParameters, options *SubscriptionCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Notify != nil {
		reqQP.Set("notify", strconv.FormatBool(*options.Notify))
	}
	reqQP.Set("api-version", "2021-01-01-preview")
	if options != nil && options.AppType != nil {
		reqQP.Set("appType", string(*options.AppType))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SubscriptionClient) createOrUpdateHandleResponse(resp *http.Response) (SubscriptionCreateOrUpdateResponse, error) {
	result := SubscriptionCreateOrUpdateResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionContract); err != nil {
		return SubscriptionCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SubscriptionClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes the specified subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SubscriptionClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, sid string, ifMatch string, options *SubscriptionDeleteOptions) (SubscriptionDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, sid, ifMatch, options)
	if err != nil {
		return SubscriptionDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SubscriptionDeleteResponse{}, client.deleteHandleError(resp)
	}
	return SubscriptionDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SubscriptionClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, ifMatch string, options *SubscriptionDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SubscriptionClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the specified Subscription entity.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SubscriptionClient) Get(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionGetOptions) (SubscriptionGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, sid, options)
	if err != nil {
		return SubscriptionGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SubscriptionClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubscriptionClient) getHandleResponse(resp *http.Response) (SubscriptionGetResponse, error) {
	result := SubscriptionGetResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionContract); err != nil {
		return SubscriptionGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SubscriptionClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetEntityTag - Gets the entity state (Etag) version of the apimanagement subscription specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SubscriptionClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionGetEntityTagOptions) (SubscriptionGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, sid, options)
	if err != nil {
		return SubscriptionGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *SubscriptionClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *SubscriptionClient) getEntityTagHandleResponse(resp *http.Response) (SubscriptionGetEntityTagResponse, error) {
	result := SubscriptionGetEntityTagResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// List - Lists all subscriptions of the API Management service instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SubscriptionClient) List(resourceGroupName string, serviceName string, options *SubscriptionListOptions) *SubscriptionListPager {
	return &SubscriptionListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
		},
		advancer: func(ctx context.Context, resp SubscriptionListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SubscriptionCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SubscriptionClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *SubscriptionListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2021-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SubscriptionClient) listHandleResponse(resp *http.Response) (SubscriptionListResponse, error) {
	result := SubscriptionListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionCollection); err != nil {
		return SubscriptionListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SubscriptionClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListSecrets - Gets the specified Subscription keys.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SubscriptionClient) ListSecrets(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionListSecretsOptions) (SubscriptionListSecretsResponse, error) {
	req, err := client.listSecretsCreateRequest(ctx, resourceGroupName, serviceName, sid, options)
	if err != nil {
		return SubscriptionListSecretsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionListSecretsResponse{}, client.listSecretsHandleError(resp)
	}
	return client.listSecretsHandleResponse(resp)
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *SubscriptionClient) listSecretsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionListSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}/listSecrets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *SubscriptionClient) listSecretsHandleResponse(resp *http.Response) (SubscriptionListSecretsResponse, error) {
	result := SubscriptionListSecretsResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionKeysContract); err != nil {
		return SubscriptionListSecretsResponse{}, err
	}
	return result, nil
}

// listSecretsHandleError handles the ListSecrets error response.
func (client *SubscriptionClient) listSecretsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// RegeneratePrimaryKey - Regenerates primary key of existing subscription of the API Management service instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SubscriptionClient) RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionRegeneratePrimaryKeyOptions) (SubscriptionRegeneratePrimaryKeyResponse, error) {
	req, err := client.regeneratePrimaryKeyCreateRequest(ctx, resourceGroupName, serviceName, sid, options)
	if err != nil {
		return SubscriptionRegeneratePrimaryKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionRegeneratePrimaryKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return SubscriptionRegeneratePrimaryKeyResponse{}, client.regeneratePrimaryKeyHandleError(resp)
	}
	return SubscriptionRegeneratePrimaryKeyResponse{RawResponse: resp}, nil
}

// regeneratePrimaryKeyCreateRequest creates the RegeneratePrimaryKey request.
func (client *SubscriptionClient) regeneratePrimaryKeyCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionRegeneratePrimaryKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}/regeneratePrimaryKey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// regeneratePrimaryKeyHandleError handles the RegeneratePrimaryKey error response.
func (client *SubscriptionClient) regeneratePrimaryKeyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// RegenerateSecondaryKey - Regenerates secondary key of existing subscription of the API Management service instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SubscriptionClient) RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionRegenerateSecondaryKeyOptions) (SubscriptionRegenerateSecondaryKeyResponse, error) {
	req, err := client.regenerateSecondaryKeyCreateRequest(ctx, resourceGroupName, serviceName, sid, options)
	if err != nil {
		return SubscriptionRegenerateSecondaryKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionRegenerateSecondaryKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return SubscriptionRegenerateSecondaryKeyResponse{}, client.regenerateSecondaryKeyHandleError(resp)
	}
	return SubscriptionRegenerateSecondaryKeyResponse{RawResponse: resp}, nil
}

// regenerateSecondaryKeyCreateRequest creates the RegenerateSecondaryKey request.
func (client *SubscriptionClient) regenerateSecondaryKeyCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, options *SubscriptionRegenerateSecondaryKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}/regenerateSecondaryKey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// regenerateSecondaryKeyHandleError handles the RegenerateSecondaryKey error response.
func (client *SubscriptionClient) regenerateSecondaryKeyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Updates the details of a subscription specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SubscriptionClient) Update(ctx context.Context, resourceGroupName string, serviceName string, sid string, ifMatch string, parameters SubscriptionUpdateParameters, options *SubscriptionUpdateOptions) (SubscriptionUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, sid, ifMatch, parameters, options)
	if err != nil {
		return SubscriptionUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SubscriptionClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, sid string, ifMatch string, parameters SubscriptionUpdateParameters, options *SubscriptionUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if sid == "" {
		return nil, errors.New("parameter sid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sid}", url.PathEscape(sid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Notify != nil {
		reqQP.Set("notify", strconv.FormatBool(*options.Notify))
	}
	reqQP.Set("api-version", "2021-01-01-preview")
	if options != nil && options.AppType != nil {
		reqQP.Set("appType", string(*options.AppType))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *SubscriptionClient) updateHandleResponse(resp *http.Response) (SubscriptionUpdateResponse, error) {
	result := SubscriptionUpdateResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionContract); err != nil {
		return SubscriptionUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *SubscriptionClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
