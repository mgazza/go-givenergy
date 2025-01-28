// Code generated by go-swagger; DO NOT EDIT.

package site

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetSingleSiteByIDReader is a Reader for the GetSingleSiteByID structure.
type GetSingleSiteByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSingleSiteByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSingleSiteByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSingleSiteByIDOK creates a GetSingleSiteByIDOK with default headers values
func NewGetSingleSiteByIDOK() *GetSingleSiteByIDOK {
	return &GetSingleSiteByIDOK{}
}

/*
GetSingleSiteByIDOK describes a response with status code 200, with default header values.

GetSingleSiteByIDOK get single site by Id o k
*/
type GetSingleSiteByIDOK struct {
	Payload *GetSingleSiteByIDOKBody
}

// IsSuccess returns true when this get single site by Id o k response has a 2xx status code
func (o *GetSingleSiteByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get single site by Id o k response has a 3xx status code
func (o *GetSingleSiteByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get single site by Id o k response has a 4xx status code
func (o *GetSingleSiteByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get single site by Id o k response has a 5xx status code
func (o *GetSingleSiteByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get single site by Id o k response a status code equal to that given
func (o *GetSingleSiteByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get single site by Id o k response
func (o *GetSingleSiteByIDOK) Code() int {
	return 200
}

func (o *GetSingleSiteByIDOK) Error() string {
	return fmt.Sprintf("[GET /site/{site_plant_id}][%d] getSingleSiteByIdOK  %+v", 200, o.Payload)
}

func (o *GetSingleSiteByIDOK) String() string {
	return fmt.Sprintf("[GET /site/{site_plant_id}][%d] getSingleSiteByIdOK  %+v", 200, o.Payload)
}

func (o *GetSingleSiteByIDOK) GetPayload() *GetSingleSiteByIDOKBody {
	return o.Payload
}

func (o *GetSingleSiteByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSingleSiteByIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetSingleSiteByIDOKBody get single site by ID o k body
// Example: {"data":{"account":"luke.phillips.146","country":"FAROE_IS","date_created":"1988-03-24","id":9,"latitude":79.284177,"longitude":1.10994,"name":"Jeremy Hill","products":[],"timezone":"GMT","type":"SINGLE_PHASE","tz_timezone":"Europe/Astrakhan"}}
swagger:model GetSingleSiteByIDOKBody
*/
type GetSingleSiteByIDOKBody struct {

	// data
	Data *GetSingleSiteByIDOKBodyData `json:"data,omitempty"`
}

// Validate validates this get single site by ID o k body
func (o *GetSingleSiteByIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSingleSiteByIDOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSingleSiteByIdOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSingleSiteByIdOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get single site by ID o k body based on the context it is used
func (o *GetSingleSiteByIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSingleSiteByIDOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSingleSiteByIdOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSingleSiteByIdOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSingleSiteByIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSingleSiteByIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetSingleSiteByIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetSingleSiteByIDOKBodyData get single site by ID o k body data
swagger:model GetSingleSiteByIDOKBodyData
*/
type GetSingleSiteByIDOKBodyData struct {

	// account
	// Example: luke.phillips.146
	Account string `json:"account,omitempty"`

	// country
	// Example: FAROE_IS
	Country string `json:"country,omitempty"`

	// date created
	// Example: 1988-03-24
	DateCreated string `json:"date_created,omitempty"`

	// id
	// Example: 9
	ID int64 `json:"id,omitempty"`

	// latitude
	// Example: 79.284177
	Latitude float64 `json:"latitude,omitempty"`

	// longitude
	// Example: 1.10994
	Longitude float64 `json:"longitude,omitempty"`

	// name
	// Example: Jeremy Hill
	Name string `json:"name,omitempty"`

	// products
	// Example: []
	Products []interface{} `json:"products"`

	// timezone
	// Example: GMT
	Timezone string `json:"timezone,omitempty"`

	// type
	// Example: SINGLE_PHASE
	Type string `json:"type,omitempty"`

	// tz timezone
	// Example: Europe/Astrakhan
	TzTimezone string `json:"tz_timezone,omitempty"`
}

// Validate validates this get single site by ID o k body data
func (o *GetSingleSiteByIDOKBodyData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get single site by ID o k body data based on context it is used
func (o *GetSingleSiteByIDOKBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSingleSiteByIDOKBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSingleSiteByIDOKBodyData) UnmarshalBinary(b []byte) error {
	var res GetSingleSiteByIDOKBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
