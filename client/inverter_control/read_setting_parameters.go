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
	"github.com/go-openapi/swag"
)

// NewReadSettingParams creates a new ReadSettingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadSettingParams() *ReadSettingParams {
	return &ReadSettingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadSettingParamsWithTimeout creates a new ReadSettingParams object
// with the ability to set a timeout on a request.
func NewReadSettingParamsWithTimeout(timeout time.Duration) *ReadSettingParams {
	return &ReadSettingParams{
		timeout: timeout,
	}
}

// NewReadSettingParamsWithContext creates a new ReadSettingParams object
// with the ability to set a context for a request.
func NewReadSettingParamsWithContext(ctx context.Context) *ReadSettingParams {
	return &ReadSettingParams{
		Context: ctx,
	}
}

// NewReadSettingParamsWithHTTPClient creates a new ReadSettingParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadSettingParamsWithHTTPClient(client *http.Client) *ReadSettingParams {
	return &ReadSettingParams{
		HTTPClient: client,
	}
}

/*
ReadSettingParams contains all the parameters to send to the API endpoint

	for the read setting operation.

	Typically these are written to a http.Request.
*/
type ReadSettingParams struct {

	// Authorization.
	Authorization *string

	// Body.
	Body ReadSettingBody

	/* InverterSerialNumber.

	   The serial number of the inverter.
	*/
	InverterSerialNumber string

	/* SettingID.

	   The ID of the setting.
	*/
	SettingID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadSettingParams) WithDefaults() *ReadSettingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadSettingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read setting params
func (o *ReadSettingParams) WithTimeout(timeout time.Duration) *ReadSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read setting params
func (o *ReadSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read setting params
func (o *ReadSettingParams) WithContext(ctx context.Context) *ReadSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read setting params
func (o *ReadSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read setting params
func (o *ReadSettingParams) WithHTTPClient(client *http.Client) *ReadSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read setting params
func (o *ReadSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the read setting params
func (o *ReadSettingParams) WithAuthorization(authorization *string) *ReadSettingParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the read setting params
func (o *ReadSettingParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithBody adds the body to the read setting params
func (o *ReadSettingParams) WithBody(body ReadSettingBody) *ReadSettingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the read setting params
func (o *ReadSettingParams) SetBody(body ReadSettingBody) {
	o.Body = body
}

// WithInverterSerialNumber adds the inverterSerialNumber to the read setting params
func (o *ReadSettingParams) WithInverterSerialNumber(inverterSerialNumber string) *ReadSettingParams {
	o.SetInverterSerialNumber(inverterSerialNumber)
	return o
}

// SetInverterSerialNumber adds the inverterSerialNumber to the read setting params
func (o *ReadSettingParams) SetInverterSerialNumber(inverterSerialNumber string) {
	o.InverterSerialNumber = inverterSerialNumber
}

// WithSettingID adds the settingID to the read setting params
func (o *ReadSettingParams) WithSettingID(settingID int64) *ReadSettingParams {
	o.SetSettingID(settingID)
	return o
}

// SetSettingID adds the settingId to the read setting params
func (o *ReadSettingParams) SetSettingID(settingID int64) {
	o.SettingID = settingID
}

// WriteToRequest writes these params to a swagger request
func (o *ReadSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param setting_id
	if err := r.SetPathParam("setting_id", swag.FormatInt64(o.SettingID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
