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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TimeZonesClient contains the methods for the TimeZones group.
// Don't use this type directly, use NewTimeZonesClient() instead.
type TimeZonesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewTimeZonesClient creates a new instance of TimeZonesClient with the specified values.
func NewTimeZonesClient(con *arm.Connection, subscriptionID string) *TimeZonesClient {
	return &TimeZonesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Gets a managed instance time zone.
// If the operation fails it returns a generic error.
func (client *TimeZonesClient) Get(ctx context.Context, locationName string, timeZoneID string, options *TimeZonesGetOptions) (TimeZonesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, timeZoneID, options)
	if err != nil {
		return TimeZonesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TimeZonesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TimeZonesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TimeZonesClient) getCreateRequest(ctx context.Context, locationName string, timeZoneID string, options *TimeZonesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/timeZones/{timeZoneId}"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if timeZoneID == "" {
		return nil, errors.New("parameter timeZoneID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{timeZoneId}", url.PathEscape(timeZoneID))
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
func (client *TimeZonesClient) getHandleResponse(resp *http.Response) (TimeZonesGetResponse, error) {
	result := TimeZonesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TimeZone); err != nil {
		return TimeZonesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *TimeZonesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByLocation - Gets a list of managed instance time zones by location.
// If the operation fails it returns a generic error.
func (client *TimeZonesClient) ListByLocation(locationName string, options *TimeZonesListByLocationOptions) *TimeZonesListByLocationPager {
	return &TimeZonesListByLocationPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByLocationCreateRequest(ctx, locationName, options)
		},
		advancer: func(ctx context.Context, resp TimeZonesListByLocationResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.TimeZoneListResult.NextLink)
		},
	}
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *TimeZonesClient) listByLocationCreateRequest(ctx context.Context, locationName string, options *TimeZonesListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/timeZones"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
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

// listByLocationHandleResponse handles the ListByLocation response.
func (client *TimeZonesClient) listByLocationHandleResponse(resp *http.Response) (TimeZonesListByLocationResponse, error) {
	result := TimeZonesListByLocationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TimeZoneListResult); err != nil {
		return TimeZonesListByLocationResponse{}, err
	}
	return result, nil
}

// listByLocationHandleError handles the ListByLocation error response.
func (client *TimeZonesClient) listByLocationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
