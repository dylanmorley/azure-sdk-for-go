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
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// TableResourcesClient contains the methods for the TableResources group.
// Don't use this type directly, use NewTableResourcesClient() instead.
type TableResourcesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewTableResourcesClient creates a new instance of TableResourcesClient with the specified values.
func NewTableResourcesClient(con *arm.Connection, subscriptionID string) *TableResourcesClient {
	return &TableResourcesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateUpdateTable - Create or update an Azure Cosmos DB Table
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) BeginCreateUpdateTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, createUpdateTableParameters TableCreateUpdateParameters, options *TableResourcesBeginCreateUpdateTableOptions) (TableResourcesCreateUpdateTablePollerResponse, error) {
	resp, err := client.createUpdateTable(ctx, resourceGroupName, accountName, tableName, createUpdateTableParameters, options)
	if err != nil {
		return TableResourcesCreateUpdateTablePollerResponse{}, err
	}
	result := TableResourcesCreateUpdateTablePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TableResourcesClient.CreateUpdateTable", "", resp, client.pl, client.createUpdateTableHandleError)
	if err != nil {
		return TableResourcesCreateUpdateTablePollerResponse{}, err
	}
	result.Poller = &TableResourcesCreateUpdateTablePoller{
		pt: pt,
	}
	return result, nil
}

// CreateUpdateTable - Create or update an Azure Cosmos DB Table
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) createUpdateTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, createUpdateTableParameters TableCreateUpdateParameters, options *TableResourcesBeginCreateUpdateTableOptions) (*http.Response, error) {
	req, err := client.createUpdateTableCreateRequest(ctx, resourceGroupName, accountName, tableName, createUpdateTableParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createUpdateTableHandleError(resp)
	}
	return resp, nil
}

// createUpdateTableCreateRequest creates the CreateUpdateTable request.
func (client *TableResourcesClient) createUpdateTableCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, createUpdateTableParameters TableCreateUpdateParameters, options *TableResourcesBeginCreateUpdateTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, createUpdateTableParameters)
}

// createUpdateTableHandleError handles the CreateUpdateTable error response.
func (client *TableResourcesClient) createUpdateTableHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDeleteTable - Deletes an existing Azure Cosmos DB Table.
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) BeginDeleteTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginDeleteTableOptions) (TableResourcesDeleteTablePollerResponse, error) {
	resp, err := client.deleteTable(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return TableResourcesDeleteTablePollerResponse{}, err
	}
	result := TableResourcesDeleteTablePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TableResourcesClient.DeleteTable", "", resp, client.pl, client.deleteTableHandleError)
	if err != nil {
		return TableResourcesDeleteTablePollerResponse{}, err
	}
	result.Poller = &TableResourcesDeleteTablePoller{
		pt: pt,
	}
	return result, nil
}

// DeleteTable - Deletes an existing Azure Cosmos DB Table.
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) deleteTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginDeleteTableOptions) (*http.Response, error) {
	req, err := client.deleteTableCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteTableHandleError(resp)
	}
	return resp, nil
}

// deleteTableCreateRequest creates the DeleteTable request.
func (client *TableResourcesClient) deleteTableCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginDeleteTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteTableHandleError handles the DeleteTable error response.
func (client *TableResourcesClient) deleteTableHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// GetTable - Gets the Tables under an existing Azure Cosmos DB database account with the provided name.
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) GetTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesGetTableOptions) (TableResourcesGetTableResponse, error) {
	req, err := client.getTableCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return TableResourcesGetTableResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TableResourcesGetTableResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TableResourcesGetTableResponse{}, client.getTableHandleError(resp)
	}
	return client.getTableHandleResponse(resp)
}

// getTableCreateRequest creates the GetTable request.
func (client *TableResourcesClient) getTableCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesGetTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getTableHandleResponse handles the GetTable response.
func (client *TableResourcesClient) getTableHandleResponse(resp *http.Response) (TableResourcesGetTableResponse, error) {
	result := TableResourcesGetTableResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TableGetResults); err != nil {
		return TableResourcesGetTableResponse{}, err
	}
	return result, nil
}

// getTableHandleError handles the GetTable error response.
func (client *TableResourcesClient) getTableHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// GetTableThroughput - Gets the RUs per second of the Table under an existing Azure Cosmos DB database account with the provided name.
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) GetTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesGetTableThroughputOptions) (TableResourcesGetTableThroughputResponse, error) {
	req, err := client.getTableThroughputCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return TableResourcesGetTableThroughputResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TableResourcesGetTableThroughputResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TableResourcesGetTableThroughputResponse{}, client.getTableThroughputHandleError(resp)
	}
	return client.getTableThroughputHandleResponse(resp)
}

// getTableThroughputCreateRequest creates the GetTableThroughput request.
func (client *TableResourcesClient) getTableThroughputCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesGetTableThroughputOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getTableThroughputHandleResponse handles the GetTableThroughput response.
func (client *TableResourcesClient) getTableThroughputHandleResponse(resp *http.Response) (TableResourcesGetTableThroughputResponse, error) {
	result := TableResourcesGetTableThroughputResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ThroughputSettingsGetResults); err != nil {
		return TableResourcesGetTableThroughputResponse{}, err
	}
	return result, nil
}

// getTableThroughputHandleError handles the GetTableThroughput error response.
func (client *TableResourcesClient) getTableThroughputHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListTables - Lists the Tables under an existing Azure Cosmos DB database account.
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) ListTables(ctx context.Context, resourceGroupName string, accountName string, options *TableResourcesListTablesOptions) (TableResourcesListTablesResponse, error) {
	req, err := client.listTablesCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return TableResourcesListTablesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TableResourcesListTablesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TableResourcesListTablesResponse{}, client.listTablesHandleError(resp)
	}
	return client.listTablesHandleResponse(resp)
}

// listTablesCreateRequest creates the ListTables request.
func (client *TableResourcesClient) listTablesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *TableResourcesListTablesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listTablesHandleResponse handles the ListTables response.
func (client *TableResourcesClient) listTablesHandleResponse(resp *http.Response) (TableResourcesListTablesResponse, error) {
	result := TableResourcesListTablesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TableListResult); err != nil {
		return TableResourcesListTablesResponse{}, err
	}
	return result, nil
}

// listTablesHandleError handles the ListTables error response.
func (client *TableResourcesClient) listTablesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginMigrateTableToAutoscale - Migrate an Azure Cosmos DB Table from manual throughput to autoscale
// If the operation fails it returns the *CloudError error type.
func (client *TableResourcesClient) BeginMigrateTableToAutoscale(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginMigrateTableToAutoscaleOptions) (TableResourcesMigrateTableToAutoscalePollerResponse, error) {
	resp, err := client.migrateTableToAutoscale(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return TableResourcesMigrateTableToAutoscalePollerResponse{}, err
	}
	result := TableResourcesMigrateTableToAutoscalePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TableResourcesClient.MigrateTableToAutoscale", "", resp, client.pl, client.migrateTableToAutoscaleHandleError)
	if err != nil {
		return TableResourcesMigrateTableToAutoscalePollerResponse{}, err
	}
	result.Poller = &TableResourcesMigrateTableToAutoscalePoller{
		pt: pt,
	}
	return result, nil
}

// MigrateTableToAutoscale - Migrate an Azure Cosmos DB Table from manual throughput to autoscale
// If the operation fails it returns the *CloudError error type.
func (client *TableResourcesClient) migrateTableToAutoscale(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginMigrateTableToAutoscaleOptions) (*http.Response, error) {
	req, err := client.migrateTableToAutoscaleCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.migrateTableToAutoscaleHandleError(resp)
	}
	return resp, nil
}

// migrateTableToAutoscaleCreateRequest creates the MigrateTableToAutoscale request.
func (client *TableResourcesClient) migrateTableToAutoscaleCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginMigrateTableToAutoscaleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default/migrateToAutoscale"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// migrateTableToAutoscaleHandleError handles the MigrateTableToAutoscale error response.
func (client *TableResourcesClient) migrateTableToAutoscaleHandleError(resp *http.Response) error {
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

// BeginMigrateTableToManualThroughput - Migrate an Azure Cosmos DB Table from autoscale to manual throughput
// If the operation fails it returns the *CloudError error type.
func (client *TableResourcesClient) BeginMigrateTableToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginMigrateTableToManualThroughputOptions) (TableResourcesMigrateTableToManualThroughputPollerResponse, error) {
	resp, err := client.migrateTableToManualThroughput(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return TableResourcesMigrateTableToManualThroughputPollerResponse{}, err
	}
	result := TableResourcesMigrateTableToManualThroughputPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TableResourcesClient.MigrateTableToManualThroughput", "", resp, client.pl, client.migrateTableToManualThroughputHandleError)
	if err != nil {
		return TableResourcesMigrateTableToManualThroughputPollerResponse{}, err
	}
	result.Poller = &TableResourcesMigrateTableToManualThroughputPoller{
		pt: pt,
	}
	return result, nil
}

// MigrateTableToManualThroughput - Migrate an Azure Cosmos DB Table from autoscale to manual throughput
// If the operation fails it returns the *CloudError error type.
func (client *TableResourcesClient) migrateTableToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginMigrateTableToManualThroughputOptions) (*http.Response, error) {
	req, err := client.migrateTableToManualThroughputCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.migrateTableToManualThroughputHandleError(resp)
	}
	return resp, nil
}

// migrateTableToManualThroughputCreateRequest creates the MigrateTableToManualThroughput request.
func (client *TableResourcesClient) migrateTableToManualThroughputCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginMigrateTableToManualThroughputOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default/migrateToManualThroughput"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// migrateTableToManualThroughputHandleError handles the MigrateTableToManualThroughput error response.
func (client *TableResourcesClient) migrateTableToManualThroughputHandleError(resp *http.Response) error {
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

// BeginUpdateTableThroughput - Update RUs per second of an Azure Cosmos DB Table
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) BeginUpdateTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, updateThroughputParameters ThroughputSettingsUpdateParameters, options *TableResourcesBeginUpdateTableThroughputOptions) (TableResourcesUpdateTableThroughputPollerResponse, error) {
	resp, err := client.updateTableThroughput(ctx, resourceGroupName, accountName, tableName, updateThroughputParameters, options)
	if err != nil {
		return TableResourcesUpdateTableThroughputPollerResponse{}, err
	}
	result := TableResourcesUpdateTableThroughputPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TableResourcesClient.UpdateTableThroughput", "", resp, client.pl, client.updateTableThroughputHandleError)
	if err != nil {
		return TableResourcesUpdateTableThroughputPollerResponse{}, err
	}
	result.Poller = &TableResourcesUpdateTableThroughputPoller{
		pt: pt,
	}
	return result, nil
}

// UpdateTableThroughput - Update RUs per second of an Azure Cosmos DB Table
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) updateTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, updateThroughputParameters ThroughputSettingsUpdateParameters, options *TableResourcesBeginUpdateTableThroughputOptions) (*http.Response, error) {
	req, err := client.updateTableThroughputCreateRequest(ctx, resourceGroupName, accountName, tableName, updateThroughputParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateTableThroughputHandleError(resp)
	}
	return resp, nil
}

// updateTableThroughputCreateRequest creates the UpdateTableThroughput request.
func (client *TableResourcesClient) updateTableThroughputCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, updateThroughputParameters ThroughputSettingsUpdateParameters, options *TableResourcesBeginUpdateTableThroughputOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, updateThroughputParameters)
}

// updateTableThroughputHandleError handles the UpdateTableThroughput error response.
func (client *TableResourcesClient) updateTableThroughputHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
