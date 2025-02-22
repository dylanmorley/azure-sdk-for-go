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

// NotebookWorkspacesClient contains the methods for the NotebookWorkspaces group.
// Don't use this type directly, use NewNotebookWorkspacesClient() instead.
type NotebookWorkspacesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewNotebookWorkspacesClient creates a new instance of NotebookWorkspacesClient with the specified values.
func NewNotebookWorkspacesClient(con *arm.Connection, subscriptionID string) *NotebookWorkspacesClient {
	return &NotebookWorkspacesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates the notebook workspace for a Cosmos DB account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotebookWorkspacesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, notebookCreateUpdateParameters NotebookWorkspaceCreateUpdateParameters, options *NotebookWorkspacesBeginCreateOrUpdateOptions) (NotebookWorkspacesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, accountName, notebookWorkspaceName, notebookCreateUpdateParameters, options)
	if err != nil {
		return NotebookWorkspacesCreateOrUpdatePollerResponse{}, err
	}
	result := NotebookWorkspacesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("NotebookWorkspacesClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return NotebookWorkspacesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &NotebookWorkspacesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates the notebook workspace for a Cosmos DB account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotebookWorkspacesClient) createOrUpdate(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, notebookCreateUpdateParameters NotebookWorkspaceCreateUpdateParameters, options *NotebookWorkspacesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, notebookCreateUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NotebookWorkspacesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, notebookCreateUpdateParameters NotebookWorkspaceCreateUpdateParameters, options *NotebookWorkspacesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, notebookCreateUpdateParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *NotebookWorkspacesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes the notebook workspace for a Cosmos DB account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotebookWorkspacesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesBeginDeleteOptions) (NotebookWorkspacesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return NotebookWorkspacesDeletePollerResponse{}, err
	}
	result := NotebookWorkspacesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("NotebookWorkspacesClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return NotebookWorkspacesDeletePollerResponse{}, err
	}
	result.Poller = &NotebookWorkspacesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the notebook workspace for a Cosmos DB account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotebookWorkspacesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NotebookWorkspacesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *NotebookWorkspacesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets the notebook workspace for a Cosmos DB account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotebookWorkspacesClient) Get(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesGetOptions) (NotebookWorkspacesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return NotebookWorkspacesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotebookWorkspacesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NotebookWorkspacesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NotebookWorkspacesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
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

// getHandleResponse handles the Get response.
func (client *NotebookWorkspacesClient) getHandleResponse(resp *http.Response) (NotebookWorkspacesGetResponse, error) {
	result := NotebookWorkspacesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookWorkspace); err != nil {
		return NotebookWorkspacesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *NotebookWorkspacesClient) getHandleError(resp *http.Response) error {
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

// ListByDatabaseAccount - Gets the notebook workspace resources of an existing Cosmos DB account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotebookWorkspacesClient) ListByDatabaseAccount(ctx context.Context, resourceGroupName string, accountName string, options *NotebookWorkspacesListByDatabaseAccountOptions) (NotebookWorkspacesListByDatabaseAccountResponse, error) {
	req, err := client.listByDatabaseAccountCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return NotebookWorkspacesListByDatabaseAccountResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotebookWorkspacesListByDatabaseAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NotebookWorkspacesListByDatabaseAccountResponse{}, client.listByDatabaseAccountHandleError(resp)
	}
	return client.listByDatabaseAccountHandleResponse(resp)
}

// listByDatabaseAccountCreateRequest creates the ListByDatabaseAccount request.
func (client *NotebookWorkspacesClient) listByDatabaseAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *NotebookWorkspacesListByDatabaseAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces"
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

// listByDatabaseAccountHandleResponse handles the ListByDatabaseAccount response.
func (client *NotebookWorkspacesClient) listByDatabaseAccountHandleResponse(resp *http.Response) (NotebookWorkspacesListByDatabaseAccountResponse, error) {
	result := NotebookWorkspacesListByDatabaseAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookWorkspaceListResult); err != nil {
		return NotebookWorkspacesListByDatabaseAccountResponse{}, err
	}
	return result, nil
}

// listByDatabaseAccountHandleError handles the ListByDatabaseAccount error response.
func (client *NotebookWorkspacesClient) listByDatabaseAccountHandleError(resp *http.Response) error {
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

// ListConnectionInfo - Retrieves the connection info for the notebook workspace
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotebookWorkspacesClient) ListConnectionInfo(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesListConnectionInfoOptions) (NotebookWorkspacesListConnectionInfoResponse, error) {
	req, err := client.listConnectionInfoCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return NotebookWorkspacesListConnectionInfoResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotebookWorkspacesListConnectionInfoResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NotebookWorkspacesListConnectionInfoResponse{}, client.listConnectionInfoHandleError(resp)
	}
	return client.listConnectionInfoHandleResponse(resp)
}

// listConnectionInfoCreateRequest creates the ListConnectionInfo request.
func (client *NotebookWorkspacesClient) listConnectionInfoCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesListConnectionInfoOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}/listConnectionInfo"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
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

// listConnectionInfoHandleResponse handles the ListConnectionInfo response.
func (client *NotebookWorkspacesClient) listConnectionInfoHandleResponse(resp *http.Response) (NotebookWorkspacesListConnectionInfoResponse, error) {
	result := NotebookWorkspacesListConnectionInfoResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookWorkspaceConnectionInfoResult); err != nil {
		return NotebookWorkspacesListConnectionInfoResponse{}, err
	}
	return result, nil
}

// listConnectionInfoHandleError handles the ListConnectionInfo error response.
func (client *NotebookWorkspacesClient) listConnectionInfoHandleError(resp *http.Response) error {
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

// BeginRegenerateAuthToken - Regenerates the auth token for the notebook workspace
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotebookWorkspacesClient) BeginRegenerateAuthToken(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesBeginRegenerateAuthTokenOptions) (NotebookWorkspacesRegenerateAuthTokenPollerResponse, error) {
	resp, err := client.regenerateAuthToken(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return NotebookWorkspacesRegenerateAuthTokenPollerResponse{}, err
	}
	result := NotebookWorkspacesRegenerateAuthTokenPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("NotebookWorkspacesClient.RegenerateAuthToken", "", resp, client.pl, client.regenerateAuthTokenHandleError)
	if err != nil {
		return NotebookWorkspacesRegenerateAuthTokenPollerResponse{}, err
	}
	result.Poller = &NotebookWorkspacesRegenerateAuthTokenPoller{
		pt: pt,
	}
	return result, nil
}

// RegenerateAuthToken - Regenerates the auth token for the notebook workspace
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotebookWorkspacesClient) regenerateAuthToken(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesBeginRegenerateAuthTokenOptions) (*http.Response, error) {
	req, err := client.regenerateAuthTokenCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.regenerateAuthTokenHandleError(resp)
	}
	return resp, nil
}

// regenerateAuthTokenCreateRequest creates the RegenerateAuthToken request.
func (client *NotebookWorkspacesClient) regenerateAuthTokenCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesBeginRegenerateAuthTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}/regenerateAuthToken"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
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

// regenerateAuthTokenHandleError handles the RegenerateAuthToken error response.
func (client *NotebookWorkspacesClient) regenerateAuthTokenHandleError(resp *http.Response) error {
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

// BeginStart - Starts the notebook workspace
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotebookWorkspacesClient) BeginStart(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesBeginStartOptions) (NotebookWorkspacesStartPollerResponse, error) {
	resp, err := client.start(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return NotebookWorkspacesStartPollerResponse{}, err
	}
	result := NotebookWorkspacesStartPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("NotebookWorkspacesClient.Start", "", resp, client.pl, client.startHandleError)
	if err != nil {
		return NotebookWorkspacesStartPollerResponse{}, err
	}
	result.Poller = &NotebookWorkspacesStartPoller{
		pt: pt,
	}
	return result, nil
}

// Start - Starts the notebook workspace
// If the operation fails it returns the *ErrorResponse error type.
func (client *NotebookWorkspacesClient) start(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.startHandleError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *NotebookWorkspacesClient) startCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}/start"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
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

// startHandleError handles the Start error response.
func (client *NotebookWorkspacesClient) startHandleError(resp *http.Response) error {
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
