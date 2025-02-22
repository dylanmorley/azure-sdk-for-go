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

// ManagedDatabaseRecommendedSensitivityLabelsClient contains the methods for the ManagedDatabaseRecommendedSensitivityLabels group.
// Don't use this type directly, use NewManagedDatabaseRecommendedSensitivityLabelsClient() instead.
type ManagedDatabaseRecommendedSensitivityLabelsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewManagedDatabaseRecommendedSensitivityLabelsClient creates a new instance of ManagedDatabaseRecommendedSensitivityLabelsClient with the specified values.
func NewManagedDatabaseRecommendedSensitivityLabelsClient(con *arm.Connection, subscriptionID string) *ManagedDatabaseRecommendedSensitivityLabelsClient {
	return &ManagedDatabaseRecommendedSensitivityLabelsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Update - Update recommended sensitivity labels states of a given database using an operations batch.
// If the operation fails it returns a generic error.
func (client *ManagedDatabaseRecommendedSensitivityLabelsClient) Update(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters RecommendedSensitivityLabelUpdateList, options *ManagedDatabaseRecommendedSensitivityLabelsUpdateOptions) (ManagedDatabaseRecommendedSensitivityLabelsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, parameters, options)
	if err != nil {
		return ManagedDatabaseRecommendedSensitivityLabelsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedDatabaseRecommendedSensitivityLabelsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedDatabaseRecommendedSensitivityLabelsUpdateResponse{}, client.updateHandleError(resp)
	}
	return ManagedDatabaseRecommendedSensitivityLabelsUpdateResponse{RawResponse: resp}, nil
}

// updateCreateRequest creates the Update request.
func (client *ManagedDatabaseRecommendedSensitivityLabelsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters RecommendedSensitivityLabelUpdateList, options *ManagedDatabaseRecommendedSensitivityLabelsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/recommendedSensitivityLabels"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *ManagedDatabaseRecommendedSensitivityLabelsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
