// Code generated by go-swagger; DO NOT EDIT.

package e_v_charger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new e v charger API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new e v charger API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new e v charger API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for e v charger API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetChargingSessions(params *GetChargingSessionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetChargingSessionsOK, error)

	GetCommandData(params *GetCommandDataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCommandDataOK, error)

	GetDataPoints(params *GetDataPointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDataPointsOK, error)

	GetEVChargerByUUID(params *GetEVChargerByUUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEVChargerByUUIDOK, error)

	GetSupportedCommands(params *GetSupportedCommandsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSupportedCommandsOK, error)

	GetYourEVChargers(params *GetYourEVChargersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetYourEVChargersOK, error)

	SendCommand(params *SendCommandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SendCommandOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetChargingSessions gets charging sessions

Fetch a list of all charging sessions for the given EV charger
*/
func (a *Client) GetChargingSessions(params *GetChargingSessionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetChargingSessionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetChargingSessionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getChargingSessions",
		Method:             "GET",
		PathPattern:        "/ev-charger/{charger_uuid}/charging-sessions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetChargingSessionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetChargingSessionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getChargingSessions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetCommandData gets command data

	Fetch the current data for the given command and EV Charger

<aside class="notice">A 422 error will be thrown if an invalid command ID is provided or if the given EV charger does not support the given command</aside>

The below response example is for the <code>change-mode</code> command
*/
func (a *Client) GetCommandData(params *GetCommandDataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCommandDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCommandDataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCommandData",
		Method:             "GET",
		PathPattern:        "/ev-charger/{charger_uuid}/commands/{command_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCommandDataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCommandDataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCommandData: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDataPoints gets data points
*/
func (a *Client) GetDataPoints(params *GetDataPointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDataPointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDataPointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDataPoints",
		Method:             "GET",
		PathPattern:        "/ev-charger/{charger_uuid}/meter-data",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDataPointsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDataPointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDataPoints: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEVChargerByUUID gets e v charger by UUID

Return information about a single EV charger by its UUID
*/
func (a *Client) GetEVChargerByUUID(params *GetEVChargerByUUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEVChargerByUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEVChargerByUUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEVChargerByUUID",
		Method:             "GET",
		PathPattern:        "/ev-charger/{charger_uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEVChargerByUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEVChargerByUUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEVChargerByUUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSupportedCommands gets supported commands

Fetch a list of all commands that are supported for the given EV charger
*/
func (a *Client) GetSupportedCommands(params *GetSupportedCommandsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSupportedCommandsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSupportedCommandsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSupportedCommands",
		Method:             "GET",
		PathPattern:        "/ev-charger/{charger_uuid}/commands",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSupportedCommandsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSupportedCommandsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSupportedCommands: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetYourEVChargers gets your e v chargers

Return a list of EV chargers registered to your account
*/
func (a *Client) GetYourEVChargers(params *GetYourEVChargersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetYourEVChargersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetYourEVChargersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getYourEVChargers",
		Method:             "GET",
		PathPattern:        "/ev-charger",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetYourEVChargersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetYourEVChargersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getYourEVChargers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SendCommand sends command

Send a command to an EV Charger
*/
func (a *Client) SendCommand(params *SendCommandParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SendCommandOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendCommandParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "sendCommand",
		Method:             "POST",
		PathPattern:        "/ev-charger/{charger_uuid}/commands/{command_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendCommandReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SendCommandOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendCommand: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
