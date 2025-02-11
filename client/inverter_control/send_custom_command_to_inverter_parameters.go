// Code generated by go-swagger; DO NOT EDIT.

package inverter_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewSendCustomCommandToInverterParams creates a new SendCustomCommandToInverterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSendCustomCommandToInverterParams() *SendCustomCommandToInverterParams {
	return &SendCustomCommandToInverterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSendCustomCommandToInverterParamsWithTimeout creates a new SendCustomCommandToInverterParams object
// with the ability to set a timeout on a request.
func NewSendCustomCommandToInverterParamsWithTimeout(timeout time.Duration) *SendCustomCommandToInverterParams {
	return &SendCustomCommandToInverterParams{
		timeout: timeout,
	}
}

// NewSendCustomCommandToInverterParamsWithContext creates a new SendCustomCommandToInverterParams object
// with the ability to set a context for a request.
func NewSendCustomCommandToInverterParamsWithContext(ctx context.Context) *SendCustomCommandToInverterParams {
	return &SendCustomCommandToInverterParams{
		Context: ctx,
	}
}

// NewSendCustomCommandToInverterParamsWithHTTPClient creates a new SendCustomCommandToInverterParams object
// with the ability to set a custom HTTPClient for a request.
func NewSendCustomCommandToInverterParamsWithHTTPClient(client *http.Client) *SendCustomCommandToInverterParams {
	return &SendCustomCommandToInverterParams{
		HTTPClient: client,
	}
}

/*
SendCustomCommandToInverterParams contains all the parameters to send to the API endpoint

	for the send custom command to inverter operation.

	Typically these are written to a http.Request.
*/
type SendCustomCommandToInverterParams struct {

	// Authorization.
	Authorization *string

	// Body.
	Body SendCustomCommandToInverterBody

	/* InverterSerialNumber.

	   The serial number of the inverter.
	*/
	InverterSerialNumber string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the send custom command to inverter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SendCustomCommandToInverterParams) WithDefaults() *SendCustomCommandToInverterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the send custom command to inverter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SendCustomCommandToInverterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the send custom command to inverter params
func (o *SendCustomCommandToInverterParams) WithTimeout(timeout time.Duration) *SendCustomCommandToInverterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send custom command to inverter params
func (o *SendCustomCommandToInverterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send custom command to inverter params
func (o *SendCustomCommandToInverterParams) WithContext(ctx context.Context) *SendCustomCommandToInverterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send custom command to inverter params
func (o *SendCustomCommandToInverterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send custom command to inverter params
func (o *SendCustomCommandToInverterParams) WithHTTPClient(client *http.Client) *SendCustomCommandToInverterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send custom command to inverter params
func (o *SendCustomCommandToInverterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the send custom command to inverter params
func (o *SendCustomCommandToInverterParams) WithAuthorization(authorization *string) *SendCustomCommandToInverterParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the send custom command to inverter params
func (o *SendCustomCommandToInverterParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithBody adds the body to the send custom command to inverter params
func (o *SendCustomCommandToInverterParams) WithBody(body SendCustomCommandToInverterBody) *SendCustomCommandToInverterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send custom command to inverter params
func (o *SendCustomCommandToInverterParams) SetBody(body SendCustomCommandToInverterBody) {
	o.Body = body
}

// WithInverterSerialNumber adds the inverterSerialNumber to the send custom command to inverter params
func (o *SendCustomCommandToInverterParams) WithInverterSerialNumber(inverterSerialNumber string) *SendCustomCommandToInverterParams {
	o.SetInverterSerialNumber(inverterSerialNumber)
	return o
}

// SetInverterSerialNumber adds the inverterSerialNumber to the send custom command to inverter params
func (o *SendCustomCommandToInverterParams) SetInverterSerialNumber(inverterSerialNumber string) {
	o.InverterSerialNumber = inverterSerialNumber
}

// WriteToRequest writes these params to a swagger request
func (o *SendCustomCommandToInverterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// header param Authorization
		if err := r.SetHeaderParam("Authorization", *o.Authorization); err != nil {
			return err
		}
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param inverter_serial_number
	if err := r.SetPathParam("inverter_serial_number", o.InverterSerialNumber); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
