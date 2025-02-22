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
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// OutboundFirewallRulesClient contains the methods for the OutboundFirewallRules group.
// Don't use this type directly, use NewOutboundFirewallRulesClient() instead.
type OutboundFirewallRulesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewOutboundFirewallRulesClient creates a new instance of OutboundFirewallRulesClient with the specified values.
func NewOutboundFirewallRulesClient(con *arm.Connection, subscriptionID string) *OutboundFirewallRulesClient {
	return &OutboundFirewallRulesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Create a outbound firewall rule with a given name.
// If the operation fails it returns a generic error.
func (client *OutboundFirewallRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, parameters OutboundFirewallRule, options *OutboundFirewallRulesBeginCreateOrUpdateOptions) (OutboundFirewallRulesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, outboundRuleFqdn, parameters, options)
	if err != nil {
		return OutboundFirewallRulesCreateOrUpdatePollerResponse{}, err
	}
	result := OutboundFirewallRulesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OutboundFirewallRulesClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return OutboundFirewallRulesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &OutboundFirewallRulesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create a outbound firewall rule with a given name.
// If the operation fails it returns a generic error.
func (client *OutboundFirewallRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, parameters OutboundFirewallRule, options *OutboundFirewallRulesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, outboundRuleFqdn, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *OutboundFirewallRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, parameters OutboundFirewallRule, options *OutboundFirewallRulesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/outboundFirewallRules/{outboundRuleFqdn}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if outboundRuleFqdn == "" {
		return nil, errors.New("parameter outboundRuleFqdn cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outboundRuleFqdn}", url.PathEscape(outboundRuleFqdn))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *OutboundFirewallRulesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Deletes a outbound firewall rule with a given name.
// If the operation fails it returns a generic error.
func (client *OutboundFirewallRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, options *OutboundFirewallRulesBeginDeleteOptions) (OutboundFirewallRulesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, outboundRuleFqdn, options)
	if err != nil {
		return OutboundFirewallRulesDeletePollerResponse{}, err
	}
	result := OutboundFirewallRulesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OutboundFirewallRulesClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return OutboundFirewallRulesDeletePollerResponse{}, err
	}
	result.Poller = &OutboundFirewallRulesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a outbound firewall rule with a given name.
// If the operation fails it returns a generic error.
func (client *OutboundFirewallRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, options *OutboundFirewallRulesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, outboundRuleFqdn, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OutboundFirewallRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, options *OutboundFirewallRulesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/outboundFirewallRules/{outboundRuleFqdn}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if outboundRuleFqdn == "" {
		return nil, errors.New("parameter outboundRuleFqdn cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outboundRuleFqdn}", url.PathEscape(outboundRuleFqdn))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *OutboundFirewallRulesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets an outbound firewall rule.
// If the operation fails it returns a generic error.
func (client *OutboundFirewallRulesClient) Get(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, options *OutboundFirewallRulesGetOptions) (OutboundFirewallRulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, outboundRuleFqdn, options)
	if err != nil {
		return OutboundFirewallRulesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OutboundFirewallRulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OutboundFirewallRulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OutboundFirewallRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, options *OutboundFirewallRulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/outboundFirewallRules/{outboundRuleFqdn}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if outboundRuleFqdn == "" {
		return nil, errors.New("parameter outboundRuleFqdn cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outboundRuleFqdn}", url.PathEscape(outboundRuleFqdn))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OutboundFirewallRulesClient) getHandleResponse(resp *http.Response) (OutboundFirewallRulesGetResponse, error) {
	result := OutboundFirewallRulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundFirewallRule); err != nil {
		return OutboundFirewallRulesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *OutboundFirewallRulesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByServer - Gets all outbound firewall rules on a server.
// If the operation fails it returns a generic error.
func (client *OutboundFirewallRulesClient) ListByServer(resourceGroupName string, serverName string, options *OutboundFirewallRulesListByServerOptions) *OutboundFirewallRulesListByServerPager {
	return &OutboundFirewallRulesListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp OutboundFirewallRulesListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OutboundFirewallRuleListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *OutboundFirewallRulesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *OutboundFirewallRulesListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/outboundFirewallRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *OutboundFirewallRulesClient) listByServerHandleResponse(resp *http.Response) (OutboundFirewallRulesListByServerResponse, error) {
	result := OutboundFirewallRulesListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundFirewallRuleListResult); err != nil {
		return OutboundFirewallRulesListByServerResponse{}, err
	}
	return result, nil
}

// listByServerHandleError handles the ListByServer error response.
func (client *OutboundFirewallRulesClient) listByServerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
