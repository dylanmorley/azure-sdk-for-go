//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// ConfigurationsCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ConfigurationsCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ConfigurationsCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ConfigurationsCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ConfigurationsCreateOrUpdateResponse will be returned.
func (p *ConfigurationsCreateOrUpdatePoller) FinalResponse(ctx context.Context) (ConfigurationsCreateOrUpdateResponse, error) {
	respType := ConfigurationsCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Configuration)
	if err != nil {
		return ConfigurationsCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ConfigurationsCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// DatabasesCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type DatabasesCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DatabasesCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *DatabasesCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DatabasesCreateOrUpdateResponse will be returned.
func (p *DatabasesCreateOrUpdatePoller) FinalResponse(ctx context.Context) (DatabasesCreateOrUpdateResponse, error) {
	respType := DatabasesCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Database)
	if err != nil {
		return DatabasesCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DatabasesCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// DatabasesDeletePoller provides polling facilities until the operation reaches a terminal state.
type DatabasesDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DatabasesDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *DatabasesDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DatabasesDeleteResponse will be returned.
func (p *DatabasesDeletePoller) FinalResponse(ctx context.Context) (DatabasesDeleteResponse, error) {
	respType := DatabasesDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return DatabasesDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DatabasesDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// FirewallRulesCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type FirewallRulesCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *FirewallRulesCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *FirewallRulesCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final FirewallRulesCreateOrUpdateResponse will be returned.
func (p *FirewallRulesCreateOrUpdatePoller) FinalResponse(ctx context.Context) (FirewallRulesCreateOrUpdateResponse, error) {
	respType := FirewallRulesCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.FirewallRule)
	if err != nil {
		return FirewallRulesCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *FirewallRulesCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// FirewallRulesDeletePoller provides polling facilities until the operation reaches a terminal state.
type FirewallRulesDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *FirewallRulesDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *FirewallRulesDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final FirewallRulesDeleteResponse will be returned.
func (p *FirewallRulesDeletePoller) FinalResponse(ctx context.Context) (FirewallRulesDeleteResponse, error) {
	respType := FirewallRulesDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return FirewallRulesDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *FirewallRulesDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateEndpointConnectionsCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionsCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateEndpointConnectionsCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *PrivateEndpointConnectionsCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateEndpointConnectionsCreateOrUpdateResponse will be returned.
func (p *PrivateEndpointConnectionsCreateOrUpdatePoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionsCreateOrUpdateResponse, error) {
	respType := PrivateEndpointConnectionsCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionsCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateEndpointConnectionsCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateEndpointConnectionsDeletePoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionsDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateEndpointConnectionsDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *PrivateEndpointConnectionsDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateEndpointConnectionsDeleteResponse will be returned.
func (p *PrivateEndpointConnectionsDeletePoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionsDeleteResponse, error) {
	respType := PrivateEndpointConnectionsDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return PrivateEndpointConnectionsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateEndpointConnectionsDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateEndpointConnectionsUpdateTagsPoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionsUpdateTagsPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateEndpointConnectionsUpdateTagsPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *PrivateEndpointConnectionsUpdateTagsPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateEndpointConnectionsUpdateTagsResponse will be returned.
func (p *PrivateEndpointConnectionsUpdateTagsPoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionsUpdateTagsResponse, error) {
	respType := PrivateEndpointConnectionsUpdateTagsResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionsUpdateTagsResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateEndpointConnectionsUpdateTagsPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ServerAdministratorsCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ServerAdministratorsCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ServerAdministratorsCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ServerAdministratorsCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ServerAdministratorsCreateOrUpdateResponse will be returned.
func (p *ServerAdministratorsCreateOrUpdatePoller) FinalResponse(ctx context.Context) (ServerAdministratorsCreateOrUpdateResponse, error) {
	respType := ServerAdministratorsCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ServerAdministratorResource)
	if err != nil {
		return ServerAdministratorsCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ServerAdministratorsCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ServerAdministratorsDeletePoller provides polling facilities until the operation reaches a terminal state.
type ServerAdministratorsDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ServerAdministratorsDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ServerAdministratorsDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ServerAdministratorsDeleteResponse will be returned.
func (p *ServerAdministratorsDeletePoller) FinalResponse(ctx context.Context) (ServerAdministratorsDeleteResponse, error) {
	respType := ServerAdministratorsDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ServerAdministratorsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ServerAdministratorsDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ServerKeysCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ServerKeysCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ServerKeysCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ServerKeysCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ServerKeysCreateOrUpdateResponse will be returned.
func (p *ServerKeysCreateOrUpdatePoller) FinalResponse(ctx context.Context) (ServerKeysCreateOrUpdateResponse, error) {
	respType := ServerKeysCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ServerKey)
	if err != nil {
		return ServerKeysCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ServerKeysCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ServerKeysDeletePoller provides polling facilities until the operation reaches a terminal state.
type ServerKeysDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ServerKeysDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ServerKeysDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ServerKeysDeleteResponse will be returned.
func (p *ServerKeysDeletePoller) FinalResponse(ctx context.Context) (ServerKeysDeleteResponse, error) {
	respType := ServerKeysDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ServerKeysDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ServerKeysDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ServerParametersListUpdateConfigurationsPoller provides polling facilities until the operation reaches a terminal state.
type ServerParametersListUpdateConfigurationsPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ServerParametersListUpdateConfigurationsPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ServerParametersListUpdateConfigurationsPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ServerParametersListUpdateConfigurationsResponse will be returned.
func (p *ServerParametersListUpdateConfigurationsPoller) FinalResponse(ctx context.Context) (ServerParametersListUpdateConfigurationsResponse, error) {
	respType := ServerParametersListUpdateConfigurationsResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ConfigurationListResult)
	if err != nil {
		return ServerParametersListUpdateConfigurationsResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ServerParametersListUpdateConfigurationsPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ServerSecurityAlertPoliciesCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ServerSecurityAlertPoliciesCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ServerSecurityAlertPoliciesCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ServerSecurityAlertPoliciesCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ServerSecurityAlertPoliciesCreateOrUpdateResponse will be returned.
func (p *ServerSecurityAlertPoliciesCreateOrUpdatePoller) FinalResponse(ctx context.Context) (ServerSecurityAlertPoliciesCreateOrUpdateResponse, error) {
	respType := ServerSecurityAlertPoliciesCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ServerSecurityAlertPolicy)
	if err != nil {
		return ServerSecurityAlertPoliciesCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ServerSecurityAlertPoliciesCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ServersCreatePoller provides polling facilities until the operation reaches a terminal state.
type ServersCreatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ServersCreatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ServersCreatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ServersCreateResponse will be returned.
func (p *ServersCreatePoller) FinalResponse(ctx context.Context) (ServersCreateResponse, error) {
	respType := ServersCreateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Server)
	if err != nil {
		return ServersCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ServersCreatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ServersDeletePoller provides polling facilities until the operation reaches a terminal state.
type ServersDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ServersDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ServersDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ServersDeleteResponse will be returned.
func (p *ServersDeletePoller) FinalResponse(ctx context.Context) (ServersDeleteResponse, error) {
	respType := ServersDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ServersDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ServersDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ServersRestartPoller provides polling facilities until the operation reaches a terminal state.
type ServersRestartPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ServersRestartPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ServersRestartPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ServersRestartResponse will be returned.
func (p *ServersRestartPoller) FinalResponse(ctx context.Context) (ServersRestartResponse, error) {
	respType := ServersRestartResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ServersRestartResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ServersRestartPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ServersUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ServersUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ServersUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ServersUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ServersUpdateResponse will be returned.
func (p *ServersUpdatePoller) FinalResponse(ctx context.Context) (ServersUpdateResponse, error) {
	respType := ServersUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Server)
	if err != nil {
		return ServersUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ServersUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualNetworkRulesCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VirtualNetworkRulesCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualNetworkRulesCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualNetworkRulesCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualNetworkRulesCreateOrUpdateResponse will be returned.
func (p *VirtualNetworkRulesCreateOrUpdatePoller) FinalResponse(ctx context.Context) (VirtualNetworkRulesCreateOrUpdateResponse, error) {
	respType := VirtualNetworkRulesCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.VirtualNetworkRule)
	if err != nil {
		return VirtualNetworkRulesCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualNetworkRulesCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualNetworkRulesDeletePoller provides polling facilities until the operation reaches a terminal state.
type VirtualNetworkRulesDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualNetworkRulesDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualNetworkRulesDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualNetworkRulesDeleteResponse will be returned.
func (p *VirtualNetworkRulesDeletePoller) FinalResponse(ctx context.Context) (VirtualNetworkRulesDeleteResponse, error) {
	respType := VirtualNetworkRulesDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VirtualNetworkRulesDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualNetworkRulesDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
