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

// NewGetSettingsListParams creates a new GetSettingsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSettingsListParams() *GetSettingsListParams {
	return &GetSettingsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSettingsListParamsWithTimeout creates a new GetSettingsListParams object
// with the ability to set a timeout on a request.
func NewGetSettingsListParamsWithTimeout(timeout time.Duration) *GetSettingsListParams {
	return &GetSettingsListParams{
		timeout: timeout,
	}
}

// NewGetSettingsListParamsWithContext creates a new GetSettingsListParams object
// with the ability to set a context for a request.
func NewGetSettingsListParamsWithContext(ctx context.Context) *GetSettingsListParams {
	return &GetSettingsListParams{
		Context: ctx,
	}
}

// NewGetSettingsListParamsWithHTTPClient creates a new GetSettingsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSettingsListParamsWithHTTPClient(client *http.Client) *GetSettingsListParams {
	return &GetSettingsListParams{
		HTTPClient: client,
	}
}

/*
GetSettingsListParams contains all the parameters to send to the API endpoint

	for the get settings list operation.

	Typically these are written to a http.Request.
*/
type GetSettingsListParams struct {

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

// WithDefaults hydrates default values in the get settings list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSettingsListParams) WithDefaults() *GetSettingsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get settings list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSettingsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get settings list params
func (o *GetSettingsListParams) WithTimeout(timeout time.Duration) *GetSettingsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get settings list params
func (o *GetSettingsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get settings list params
func (o *GetSettingsListParams) WithContext(ctx context.Context) *GetSettingsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get settings list params
func (o *GetSettingsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get settings list params
func (o *GetSettingsListParams) WithHTTPClient(client *http.Client) *GetSettingsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get settings list params
func (o *GetSettingsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get settings list params
func (o *GetSettingsListParams) WithAuthorization(authorization *string) *GetSettingsListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get settings list params
func (o *GetSettingsListParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithInverterSerialNumber adds the inverterSerialNumber to the get settings list params
func (o *GetSettingsListParams) WithInverterSerialNumber(inverterSerialNumber string) *GetSettingsListParams {
	o.SetInverterSerialNumber(inverterSerialNumber)
	return o
}

// SetInverterSerialNumber adds the inverterSerialNumber to the get settings list params
func (o *GetSettingsListParams) SetInverterSerialNumber(inverterSerialNumber string) {
	o.InverterSerialNumber = inverterSerialNumber
}

// WriteToRequest writes these params to a swagger request
func (o *GetSettingsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
