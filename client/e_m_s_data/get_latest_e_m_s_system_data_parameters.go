// Code generated by go-swagger; DO NOT EDIT.

package e_m_s_data

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

// NewGetLatestEMSSystemDataParams creates a new GetLatestEMSSystemDataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLatestEMSSystemDataParams() *GetLatestEMSSystemDataParams {
	return &GetLatestEMSSystemDataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLatestEMSSystemDataParamsWithTimeout creates a new GetLatestEMSSystemDataParams object
// with the ability to set a timeout on a request.
func NewGetLatestEMSSystemDataParamsWithTimeout(timeout time.Duration) *GetLatestEMSSystemDataParams {
	return &GetLatestEMSSystemDataParams{
		timeout: timeout,
	}
}

// NewGetLatestEMSSystemDataParamsWithContext creates a new GetLatestEMSSystemDataParams object
// with the ability to set a context for a request.
func NewGetLatestEMSSystemDataParamsWithContext(ctx context.Context) *GetLatestEMSSystemDataParams {
	return &GetLatestEMSSystemDataParams{
		Context: ctx,
	}
}

// NewGetLatestEMSSystemDataParamsWithHTTPClient creates a new GetLatestEMSSystemDataParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLatestEMSSystemDataParamsWithHTTPClient(client *http.Client) *GetLatestEMSSystemDataParams {
	return &GetLatestEMSSystemDataParams{
		HTTPClient: client,
	}
}

/*
GetLatestEMSSystemDataParams contains all the parameters to send to the API endpoint

	for the get latest e m s system data operation.

	Typically these are written to a http.Request.
*/
type GetLatestEMSSystemDataParams struct {

	// Authorization.
	Authorization *string

	/* InverterSerialNumber.

	   Optional parameter. The serial number of the EMS
	*/
	InverterSerialNumber string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get latest e m s system data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLatestEMSSystemDataParams) WithDefaults() *GetLatestEMSSystemDataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get latest e m s system data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLatestEMSSystemDataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get latest e m s system data params
func (o *GetLatestEMSSystemDataParams) WithTimeout(timeout time.Duration) *GetLatestEMSSystemDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get latest e m s system data params
func (o *GetLatestEMSSystemDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get latest e m s system data params
func (o *GetLatestEMSSystemDataParams) WithContext(ctx context.Context) *GetLatestEMSSystemDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get latest e m s system data params
func (o *GetLatestEMSSystemDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get latest e m s system data params
func (o *GetLatestEMSSystemDataParams) WithHTTPClient(client *http.Client) *GetLatestEMSSystemDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get latest e m s system data params
func (o *GetLatestEMSSystemDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get latest e m s system data params
func (o *GetLatestEMSSystemDataParams) WithAuthorization(authorization *string) *GetLatestEMSSystemDataParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get latest e m s system data params
func (o *GetLatestEMSSystemDataParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithInverterSerialNumber adds the inverterSerialNumber to the get latest e m s system data params
func (o *GetLatestEMSSystemDataParams) WithInverterSerialNumber(inverterSerialNumber string) *GetLatestEMSSystemDataParams {
	o.SetInverterSerialNumber(inverterSerialNumber)
	return o
}

// SetInverterSerialNumber adds the inverterSerialNumber to the get latest e m s system data params
func (o *GetLatestEMSSystemDataParams) SetInverterSerialNumber(inverterSerialNumber string) {
	o.InverterSerialNumber = inverterSerialNumber
}

// WriteToRequest writes these params to a swagger request
func (o *GetLatestEMSSystemDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
