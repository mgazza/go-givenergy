// Code generated by go-swagger; DO NOT EDIT.

package site

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetYourSitesReader is a Reader for the GetYourSites structure.
type GetYourSitesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetYourSitesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetYourSitesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetYourSitesOK creates a GetYourSitesOK with default headers values
func NewGetYourSitesOK() *GetYourSitesOK {
	return &GetYourSitesOK{}
}

/*
GetYourSitesOK describes a response with status code 200, with default header values.

GetYourSitesOK get your sites o k
*/
type GetYourSitesOK struct {
	Payload *GetYourSitesOKBody
}

// IsSuccess returns true when this get your sites o k response has a 2xx status code
func (o *GetYourSitesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get your sites o k response has a 3xx status code
func (o *GetYourSitesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get your sites o k response has a 4xx status code
func (o *GetYourSitesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get your sites o k response has a 5xx status code
func (o *GetYourSitesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get your sites o k response a status code equal to that given
func (o *GetYourSitesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get your sites o k response
func (o *GetYourSitesOK) Code() int {
	return 200
}

func (o *GetYourSitesOK) Error() string {
	return fmt.Sprintf("[GET /site][%d] getYourSitesOK  %+v", 200, o.Payload)
}

func (o *GetYourSitesOK) String() string {
	return fmt.Sprintf("[GET /site][%d] getYourSitesOK  %+v", 200, o.Payload)
}

func (o *GetYourSitesOK) GetPayload() *GetYourSitesOKBody {
	return o.Payload
}

func (o *GetYourSitesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetYourSitesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetYourSitesOKBody get your sites o k body
// Example: {"data":[{"account":"yreynolds.814","country":"FAROE_IS","date_created":"1988-03-24","id":7,"latitude":79.284177,"longitude":1.10994,"name":"Jeremy Hill","products":[],"timezone":"GMT","type":"SINGLE_PHASE","tz_timezone":"Europe/Astrakhan"},{"account":"rlewis.7","country":"NIUE","date_created":"2023-02-07","id":8,"latitude":-77.744751,"longitude":-175.172025,"name":"Riley Adams","products":[],"timezone":"GMT","type":"SINGLE_PHASE","tz_timezone":"Pacific/Honolulu"}]}
swagger:model GetYourSitesOKBody
*/
type GetYourSitesOKBody struct {

	// data
	// Example: [{"account":"yreynolds.814","country":"FAROE_IS","date_created":"1988-03-24","id":7,"latitude":79.284177,"longitude":1.10994,"name":"Jeremy Hill","products":[],"timezone":"GMT","type":"SINGLE_PHASE","tz_timezone":"Europe/Astrakhan"},{"account":"rlewis.7","country":"NIUE","date_created":"2023-02-07","id":8,"latitude":-77.744751,"longitude":-175.172025,"name":"Riley Adams","products":[],"timezone":"GMT","type":"SINGLE_PHASE","tz_timezone":"Pacific/Honolulu"}]
	Data []*GetYourSitesOKBodyDataItems0 `json:"data"`
}

// Validate validates this get your sites o k body
func (o *GetYourSitesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetYourSitesOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getYourSitesOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getYourSitesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get your sites o k body based on the context it is used
func (o *GetYourSitesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetYourSitesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getYourSitesOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getYourSitesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetYourSitesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetYourSitesOKBody) UnmarshalBinary(b []byte) error {
	var res GetYourSitesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetYourSitesOKBodyDataItems0 get your sites o k body data items0
swagger:model GetYourSitesOKBodyDataItems0
*/
type GetYourSitesOKBodyDataItems0 struct {

	// account
	// Example: yreynolds.814
	Account string `json:"account,omitempty"`

	// country
	// Example: FAROE_IS
	Country string `json:"country,omitempty"`

	// date created
	// Example: 1988-03-24
	DateCreated string `json:"date_created,omitempty"`

	// id
	// Example: 7
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

// Validate validates this get your sites o k body data items0
func (o *GetYourSitesOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get your sites o k body data items0 based on context it is used
func (o *GetYourSitesOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetYourSitesOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetYourSitesOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res GetYourSitesOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
