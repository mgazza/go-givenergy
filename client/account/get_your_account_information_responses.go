// Code generated by go-swagger; DO NOT EDIT.

package account

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

// GetYourAccountInformationReader is a Reader for the GetYourAccountInformation structure.
type GetYourAccountInformationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetYourAccountInformationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetYourAccountInformationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetYourAccountInformationOK creates a GetYourAccountInformationOK with default headers values
func NewGetYourAccountInformationOK() *GetYourAccountInformationOK {
	return &GetYourAccountInformationOK{}
}

/*
GetYourAccountInformationOK describes a response with status code 200, with default header values.

GetYourAccountInformationOK get your account information o k
*/
type GetYourAccountInformationOK struct {
	Payload *GetYourAccountInformationOKBody
}

// IsSuccess returns true when this get your account information o k response has a 2xx status code
func (o *GetYourAccountInformationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get your account information o k response has a 3xx status code
func (o *GetYourAccountInformationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get your account information o k response has a 4xx status code
func (o *GetYourAccountInformationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get your account information o k response has a 5xx status code
func (o *GetYourAccountInformationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get your account information o k response a status code equal to that given
func (o *GetYourAccountInformationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get your account information o k response
func (o *GetYourAccountInformationOK) Code() int {
	return 200
}

func (o *GetYourAccountInformationOK) Error() string {
	return fmt.Sprintf("[GET /account][%d] getYourAccountInformationOK  %+v", 200, o.Payload)
}

func (o *GetYourAccountInformationOK) String() string {
	return fmt.Sprintf("[GET /account][%d] getYourAccountInformationOK  %+v", 200, o.Payload)
}

func (o *GetYourAccountInformationOK) GetPayload() *GetYourAccountInformationOKBody {
	return o.Payload
}

func (o *GetYourAccountInformationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetYourAccountInformationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetYourAccountInformationOKBody get your account information o k body
// Example: {"data":{"address":"Flat 60\nKyle Lights","company":null,"country":"UNITED_KINGDOM","email":"kelly44@allen.co.uk","first_name":"Anna","id":2,"name":"frank28.203","postcode":"SP1 1NE","role":"SUPER_ENGINEER","standard_timezone":"Europe/London","surname":"Adams","telephone_number":"0738 286 7656","timezone":"GMT"}}
swagger:model GetYourAccountInformationOKBody
*/
type GetYourAccountInformationOKBody struct {

	// data
	Data *GetYourAccountInformationOKBodyData `json:"data,omitempty"`
}

// Validate validates this get your account information o k body
func (o *GetYourAccountInformationOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetYourAccountInformationOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getYourAccountInformationOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getYourAccountInformationOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get your account information o k body based on the context it is used
func (o *GetYourAccountInformationOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetYourAccountInformationOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getYourAccountInformationOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getYourAccountInformationOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetYourAccountInformationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetYourAccountInformationOKBody) UnmarshalBinary(b []byte) error {
	var res GetYourAccountInformationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetYourAccountInformationOKBodyData get your account information o k body data
swagger:model GetYourAccountInformationOKBodyData
*/
type GetYourAccountInformationOKBodyData struct {

	// address
	// Example: Flat 60\nKyle Lights
	Address string `json:"address,omitempty"`

	// company
	Company string `json:"company,omitempty"`

	// country
	// Example: UNITED_KINGDOM
	Country string `json:"country,omitempty"`

	// email
	// Example: kelly44@allen.co.uk
	Email string `json:"email,omitempty"`

	// first name
	// Example: Anna
	FirstName string `json:"first_name,omitempty"`

	// id
	// Example: 2
	ID int64 `json:"id,omitempty"`

	// name
	// Example: frank28.203
	Name string `json:"name,omitempty"`

	// postcode
	// Example: SP1 1NE
	Postcode string `json:"postcode,omitempty"`

	// role
	// Example: SUPER_ENGINEER
	Role string `json:"role,omitempty"`

	// standard timezone
	// Example: Europe/London
	StandardTimezone string `json:"standard_timezone,omitempty"`

	// surname
	// Example: Adams
	Surname string `json:"surname,omitempty"`

	// telephone number
	// Example: 0738 286 7656
	TelephoneNumber string `json:"telephone_number,omitempty"`

	// timezone
	// Example: GMT
	Timezone string `json:"timezone,omitempty"`
}

// Validate validates this get your account information o k body data
func (o *GetYourAccountInformationOKBodyData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get your account information o k body data based on context it is used
func (o *GetYourAccountInformationOKBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetYourAccountInformationOKBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetYourAccountInformationOKBodyData) UnmarshalBinary(b []byte) error {
	var res GetYourAccountInformationOKBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
