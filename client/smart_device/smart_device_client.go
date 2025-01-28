// Code generated by go-swagger; DO NOT EDIT.

package smart_device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new smart device API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for smart device API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSmartDevice(params *CreateSmartDeviceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSmartDeviceOK, error)

	CreateSmartDeviceDataPoint(params *CreateSmartDeviceDataPointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSmartDeviceDataPointOK, error)

	GetSmartDeviceByID(params *GetSmartDeviceByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSmartDeviceByIDOK, error)

	GetSmartDeviceDataPointsByID(params *GetSmartDeviceDataPointsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSmartDeviceDataPointsByIDOK, error)

	GetYourSmartDevices(params *GetYourSmartDevicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetYourSmartDevicesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateSmartDevice creates smart device

Register a new smart device to your account
*/
func (a *Client) CreateSmartDevice(params *CreateSmartDeviceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSmartDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSmartDeviceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSmartDevice",
		Method:             "POST",
		PathPattern:        "/smart-device",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSmartDeviceReader{formats: a.formats},
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
	success, ok := result.(*CreateSmartDeviceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createSmartDevice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateSmartDeviceDataPoint creates smart device data point

Store a data point against a smart device
*/
func (a *Client) CreateSmartDeviceDataPoint(params *CreateSmartDeviceDataPointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSmartDeviceDataPointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSmartDeviceDataPointParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSmartDeviceDataPoint",
		Method:             "POST",
		PathPattern:        "/smart-device/{smartDevice_uuid}/data",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSmartDeviceDataPointReader{formats: a.formats},
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
	success, ok := result.(*CreateSmartDeviceDataPointOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createSmartDeviceDataPoint: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSmartDeviceByID gets smart device by ID

Get a smart device's information
*/
func (a *Client) GetSmartDeviceByID(params *GetSmartDeviceByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSmartDeviceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSmartDeviceByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSmartDeviceByID",
		Method:             "GET",
		PathPattern:        "/smart-device/{smartDevice_uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSmartDeviceByIDReader{formats: a.formats},
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
	success, ok := result.(*GetSmartDeviceByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSmartDeviceByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSmartDeviceDataPointsByID gets smart device data points by ID

Get a list of a smart device's data points
*/
func (a *Client) GetSmartDeviceDataPointsByID(params *GetSmartDeviceDataPointsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSmartDeviceDataPointsByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSmartDeviceDataPointsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSmartDeviceDataPointsByID",
		Method:             "GET",
		PathPattern:        "/smart-device/{smartDevice_uuid}/data",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSmartDeviceDataPointsByIDReader{formats: a.formats},
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
	success, ok := result.(*GetSmartDeviceDataPointsByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSmartDeviceDataPointsByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetYourSmartDevices gets your smart devices

List the smart devices registered to your account
*/
func (a *Client) GetYourSmartDevices(params *GetYourSmartDevicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetYourSmartDevicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetYourSmartDevicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getYourSmartDevices",
		Method:             "GET",
		PathPattern:        "/smart-device",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetYourSmartDevicesReader{formats: a.formats},
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
	success, ok := result.(*GetYourSmartDevicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getYourSmartDevices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
