//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// RestorableSQLContainersClient contains the methods for the RestorableSQLContainers group.
// Don't use this type directly, use NewRestorableSQLContainersClient() instead.
type RestorableSQLContainersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRestorableSQLContainersClient creates a new instance of RestorableSQLContainersClient with the specified values.
func NewRestorableSQLContainersClient(con *arm.Connection, subscriptionID string) *RestorableSQLContainersClient {
	return &RestorableSQLContainersClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - Show the event feed of all mutations done on all the Azure Cosmos DB SQL containers under a specific database. This helps in scenario where container
// was accidentally deleted. This API requires
// 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission
// If the operation fails it returns the *CloudError error type.
func (client *RestorableSQLContainersClient) List(ctx context.Context, location string, instanceID string, options *RestorableSQLContainersListOptions) (RestorableSQLContainersListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, instanceID, options)
	if err != nil {
		return RestorableSQLContainersListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorableSQLContainersListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorableSQLContainersListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *RestorableSQLContainersClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableSQLContainersListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableSqlContainers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	if options != nil && options.RestorableSQLDatabaseRid != nil {
		reqQP.Set("restorableSqlDatabaseRid", *options.RestorableSQLDatabaseRid)
	}
	if options != nil && options.StartTime != nil {
		reqQP.Set("startTime", *options.StartTime)
	}
	if options != nil && options.EndTime != nil {
		reqQP.Set("endTime", *options.EndTime)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableSQLContainersClient) listHandleResponse(resp *http.Response) (RestorableSQLContainersListResponse, error) {
	result := RestorableSQLContainersListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableSQLContainersListResult); err != nil {
		return RestorableSQLContainersListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *RestorableSQLContainersClient) listHandleError(resp *http.Response) error {
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
