// Code generated by go-swagger; DO NOT EDIT.

package inverter_data

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

// NewRequestInverterDataParams creates a new RequestInverterDataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRequestInverterDataParams() *RequestInverterDataParams {
	return &RequestInverterDataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRequestInverterDataParamsWithTimeout creates a new RequestInverterDataParams object
// with the ability to set a timeout on a request.
func NewRequestInverterDataParamsWithTimeout(timeout time.Duration) *RequestInverterDataParams {
	return &RequestInverterDataParams{
		timeout: timeout,
	}
}

// NewRequestInverterDataParamsWithContext creates a new RequestInverterDataParams object
// with the ability to set a context for a request.
func NewRequestInverterDataParamsWithContext(ctx context.Context) *RequestInverterDataParams {
	return &RequestInverterDataParams{
		Context: ctx,
	}
}

// NewRequestInverterDataParamsWithHTTPClient creates a new RequestInverterDataParams object
// with the ability to set a custom HTTPClient for a request.
func NewRequestInverterDataParamsWithHTTPClient(client *http.Client) *RequestInverterDataParams {
	return &RequestInverterDataParams{
		HTTPClient: client,
	}
}

/*
RequestInverterDataParams contains all the parameters to send to the API endpoint

	for the request inverter data operation.

	Typically these are written to a http.Request.
*/
type RequestInverterDataParams struct {

	// Authorization.
	Authorization *string

	/* InverterSerialNumber.

	   The serial number of the inverter.
	*/
	InverterSerialNumber string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the request inverter data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RequestInverterDataParams) WithDefaults() *RequestInverterDataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the request inverter data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RequestInverterDataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the request inverter data params
func (o *RequestInverterDataParams) WithTimeout(timeout time.Duration) *RequestInverterDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request inverter data params
func (o *RequestInverterDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request inverter data params
func (o *RequestInverterDataParams) WithContext(ctx context.Context) *RequestInverterDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request inverter data params
func (o *RequestInverterDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request inverter data params
func (o *RequestInverterDataParams) WithHTTPClient(client *http.Client) *RequestInverterDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request inverter data params
func (o *RequestInverterDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the request inverter data params
func (o *RequestInverterDataParams) WithAuthorization(authorization *string) *RequestInverterDataParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the request inverter data params
func (o *RequestInverterDataParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithInverterSerialNumber adds the inverterSerialNumber to the request inverter data params
func (o *RequestInverterDataParams) WithInverterSerialNumber(inverterSerialNumber string) *RequestInverterDataParams {
	o.SetInverterSerialNumber(inverterSerialNumber)
	return o
}

// SetInverterSerialNumber adds the inverterSerialNumber to the request inverter data params
func (o *RequestInverterDataParams) SetInverterSerialNumber(inverterSerialNumber string) {
	o.InverterSerialNumber = inverterSerialNumber
}

// WriteToRequest writes these params to a swagger request
func (o *RequestInverterDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param inverter_serial_number
	if err := r.SetPathParam("inverter_serial_number", o.InverterSerialNumber); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
