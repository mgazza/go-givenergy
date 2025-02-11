// Code generated by go-swagger; DO NOT EDIT.

package inverter_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new inverter control API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new inverter control API client with basic auth credentials.
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

// New creates a new inverter control API client with a bearer token for authentication.
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
Client for inverter control API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithAccept allows the client to force the Accept header
// to negotiate a specific Producer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithAccept(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ProducesMediaTypes = []string{mime}
	}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// WithAcceptTextPlain sets the Accept header to "text/plain".
func WithAcceptTextPlain(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"text/plain"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetPresetValues(params *GetPresetValuesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPresetValuesOK, error)

	GetSettingPresets(params *GetSettingPresetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSettingPresetsOK, error)

	GetSettingsList(params *GetSettingsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSettingsListOK, error)

	ModifySetting(params *ModifySettingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifySettingOK, error)

	ModifySettingMultiple(params *ModifySettingMultipleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifySettingMultipleOK, error)

	ReadSetting(params *ReadSettingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadSettingOK, error)

	ReadSettingMultiple(params *ReadSettingMultipleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadSettingMultipleOK, error)

	SendCustomCommandToInverter(params *SendCustomCommandToInverterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SendCustomCommandToInverterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetPresetValues gets preset values

	Get a preset's current values

The returned data will be formatted in the same way as the preset's parameters. Examples can be found on the right

<aside class="notice">If a particular value is not yet known as it hasn't been set on or read from the inverter,
an assumed default value or null may be returned, depending on the preset and value</aside>
*/
func (a *Client) GetPresetValues(params *GetPresetValuesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPresetValuesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPresetValuesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPresetValues",
		Method:             "GET",
		PathPattern:        "/inverter/{inverter_serial_number}/presets/{id}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPresetValuesReader{formats: a.formats},
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
	success, ok := result.(*GetPresetValuesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPresetValues: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSettingPresets gets setting presets

Retrieves a list of available setting presets for a given inverter
*/
func (a *Client) GetSettingPresets(params *GetSettingPresetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSettingPresetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSettingPresetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSettingPresets",
		Method:             "GET",
		PathPattern:        "/inverter/{inverter_serial_number}/presets",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSettingPresetsReader{formats: a.formats},
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
	success, ok := result.(*GetSettingPresetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSettingPresets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSettingsList gets settings list

Returns a set of inverter settings available to your account
*/
func (a *Client) GetSettingsList(params *GetSettingsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSettingsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSettingsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSettingsList",
		Method:             "GET",
		PathPattern:        "/inverter/{inverter_serial_number}/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSettingsListReader{formats: a.formats},
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
	success, ok := result.(*GetSettingsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSettingsList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ModifySetting modifies setting

Write a value to the setting on the inverter
*/
func (a *Client) ModifySetting(params *ModifySettingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifySettingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifySettingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "modifySetting",
		Method:             "POST",
		PathPattern:        "/inverter/{inverter_serial_number}/settings/{setting_id}/write",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifySettingReader{formats: a.formats},
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
	success, ok := result.(*ModifySettingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for modifySetting: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ModifySettingMultiple modifies setting multiple

	Write a value to a setting on multiple inverters

<aside class="warning">Only attempt to modify the setting for inverters that you are certain exist and which you have access to.
Improper use of this endpoint may lead to restrictions being imposed on your account</aside>
*/
func (a *Client) ModifySettingMultiple(params *ModifySettingMultipleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifySettingMultipleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifySettingMultipleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "modifySettingMultiple",
		Method:             "POST",
		PathPattern:        "/multi-control/write",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifySettingMultipleReader{formats: a.formats},
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
	success, ok := result.(*ModifySettingMultipleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for modifySettingMultiple: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReadSetting reads setting

Read a specific setting on the inverter
*/
func (a *Client) ReadSetting(params *ReadSettingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadSettingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadSettingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "readSetting",
		Method:             "POST",
		PathPattern:        "/inverter/{inverter_serial_number}/settings/{setting_id}/read",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadSettingReader{formats: a.formats},
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
	success, ok := result.(*ReadSettingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readSetting: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ReadSettingMultiple reads setting multiple

	Read a specific setting on multiple inverters

<aside class="warning">Only attempt to read the setting for inverters that you are certain exist and which you have access to.
Improper use of this endpoint may lead to restrictions being imposed on your account</aside>
*/
func (a *Client) ReadSettingMultiple(params *ReadSettingMultipleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadSettingMultipleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadSettingMultipleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "readSettingMultiple",
		Method:             "POST",
		PathPattern:        "/multi-control/read",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadSettingMultipleReader{formats: a.formats},
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
	success, ok := result.(*ReadSettingMultipleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readSettingMultiple: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SendCustomCommandToInverter sends custom command to inverter

	Transparently sends the given hexadecimal string to the inverter

<aside class="notice">

	A response code will only be returned if the command was not successful

</aside>

See [Remote Control Codes](#remote-control-codes) for all possible response codes
*/
func (a *Client) SendCustomCommandToInverter(params *SendCustomCommandToInverterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SendCustomCommandToInverterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendCustomCommandToInverterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "sendCustomCommandToInverter",
		Method:             "POST",
		PathPattern:        "/inverter/{inverter_serial_number}/debug/transparent/send",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendCustomCommandToInverterReader{formats: a.formats},
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
	success, ok := result.(*SendCustomCommandToInverterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendCustomCommandToInverter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
