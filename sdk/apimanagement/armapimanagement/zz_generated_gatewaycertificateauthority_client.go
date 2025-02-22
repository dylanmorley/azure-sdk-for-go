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

// GatewayCertificateAuthorityClient contains the methods for the GatewayCertificateAuthority group.
// Don't use this type directly, use NewGatewayCertificateAuthorityClient() instead.
type GatewayCertificateAuthorityClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewGatewayCertificateAuthorityClient creates a new instance of GatewayCertificateAuthorityClient with the specified values.
func NewGatewayCertificateAuthorityClient(con *arm.Connection, subscriptionID string) *GatewayCertificateAuthorityClient {
	return &GatewayCertificateAuthorityClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Assign Certificate entity to Gateway entity as Certificate Authority.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayCertificateAuthorityClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, parameters GatewayCertificateAuthorityContract, options *GatewayCertificateAuthorityCreateOrUpdateOptions) (GatewayCertificateAuthorityCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, certificateID, parameters, options)
	if err != nil {
		return GatewayCertificateAuthorityCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayCertificateAuthorityCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return GatewayCertificateAuthorityCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GatewayCertificateAuthorityClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, parameters GatewayCertificateAuthorityContract, options *GatewayCertificateAuthorityCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/certificateAuthorities/{certificateId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if certificateID == "" {
		return nil, errors.New("parameter certificateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateId}", url.PathEscape(certificateID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *GatewayCertificateAuthorityClient) createOrUpdateHandleResponse(resp *http.Response) (GatewayCertificateAuthorityCreateOrUpdateResponse, error) {
	result := GatewayCertificateAuthorityCreateOrUpdateResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayCertificateAuthorityContract); err != nil {
		return GatewayCertificateAuthorityCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *GatewayCertificateAuthorityClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Remove relationship between Certificate Authority and Gateway entity.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayCertificateAuthorityClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, ifMatch string, options *GatewayCertificateAuthorityDeleteOptions) (GatewayCertificateAuthorityDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, certificateID, ifMatch, options)
	if err != nil {
		return GatewayCertificateAuthorityDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayCertificateAuthorityDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return GatewayCertificateAuthorityDeleteResponse{}, client.deleteHandleError(resp)
	}
	return GatewayCertificateAuthorityDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GatewayCertificateAuthorityClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, ifMatch string, options *GatewayCertificateAuthorityDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/certificateAuthorities/{certificateId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if certificateID == "" {
		return nil, errors.New("parameter certificateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateId}", url.PathEscape(certificateID))
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
func (client *GatewayCertificateAuthorityClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get assigned Gateway Certificate Authority details.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayCertificateAuthorityClient) Get(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, options *GatewayCertificateAuthorityGetOptions) (GatewayCertificateAuthorityGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, certificateID, options)
	if err != nil {
		return GatewayCertificateAuthorityGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayCertificateAuthorityGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GatewayCertificateAuthorityGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GatewayCertificateAuthorityClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, options *GatewayCertificateAuthorityGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/certificateAuthorities/{certificateId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if certificateID == "" {
		return nil, errors.New("parameter certificateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateId}", url.PathEscape(certificateID))
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
func (client *GatewayCertificateAuthorityClient) getHandleResponse(resp *http.Response) (GatewayCertificateAuthorityGetResponse, error) {
	result := GatewayCertificateAuthorityGetResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayCertificateAuthorityContract); err != nil {
		return GatewayCertificateAuthorityGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *GatewayCertificateAuthorityClient) getHandleError(resp *http.Response) error {
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

// GetEntityTag - Checks if Certificate entity is assigned to Gateway entity as Certificate Authority.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayCertificateAuthorityClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, options *GatewayCertificateAuthorityGetEntityTagOptions) (GatewayCertificateAuthorityGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, certificateID, options)
	if err != nil {
		return GatewayCertificateAuthorityGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayCertificateAuthorityGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *GatewayCertificateAuthorityClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, options *GatewayCertificateAuthorityGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/certificateAuthorities/{certificateId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if certificateID == "" {
		return nil, errors.New("parameter certificateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateId}", url.PathEscape(certificateID))
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
func (client *GatewayCertificateAuthorityClient) getEntityTagHandleResponse(resp *http.Response) (GatewayCertificateAuthorityGetEntityTagResponse, error) {
	result := GatewayCertificateAuthorityGetEntityTagResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Lists the collection of Certificate Authorities for the specified Gateway entity.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayCertificateAuthorityClient) ListByService(resourceGroupName string, serviceName string, gatewayID string, options *GatewayCertificateAuthorityListByServiceOptions) *GatewayCertificateAuthorityListByServicePager {
	return &GatewayCertificateAuthorityListByServicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, options)
		},
		advancer: func(ctx context.Context, resp GatewayCertificateAuthorityListByServiceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.GatewayCertificateAuthorityCollection.NextLink)
		},
	}
}

// listByServiceCreateRequest creates the ListByService request.
func (client *GatewayCertificateAuthorityClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, options *GatewayCertificateAuthorityListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/certificateAuthorities"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
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

// listByServiceHandleResponse handles the ListByService response.
func (client *GatewayCertificateAuthorityClient) listByServiceHandleResponse(resp *http.Response) (GatewayCertificateAuthorityListByServiceResponse, error) {
	result := GatewayCertificateAuthorityListByServiceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayCertificateAuthorityCollection); err != nil {
		return GatewayCertificateAuthorityListByServiceResponse{}, err
	}
	return result, nil
}

// listByServiceHandleError handles the ListByService error response.
func (client *GatewayCertificateAuthorityClient) listByServiceHandleError(resp *http.Response) error {
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
