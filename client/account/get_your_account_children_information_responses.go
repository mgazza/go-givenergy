// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetYourAccountChildrenInformationReader is a Reader for the GetYourAccountChildrenInformation structure.
type GetYourAccountChildrenInformationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetYourAccountChildrenInformationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetYourAccountChildrenInformationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /account-children] getYourAccountChildrenInformation", response, response.Code())
	}
}

// NewGetYourAccountChildrenInformationOK creates a GetYourAccountChildrenInformationOK with default headers values
func NewGetYourAccountChildrenInformationOK() *GetYourAccountChildrenInformationOK {
	return &GetYourAccountChildrenInformationOK{}
}

/*
GetYourAccountChildrenInformationOK describes a response with status code 200, with default header values.

GetYourAccountChildrenInformationOK get your account children information o k
*/
type GetYourAccountChildrenInformationOK struct {
	Payload *GetYourAccountChildrenInformationOKBody
}

// IsSuccess returns true when this get your account children information o k response has a 2xx status code
func (o *GetYourAccountChildrenInformationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get your account children information o k response has a 3xx status code
func (o *GetYourAccountChildrenInformationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get your account children information o k response has a 4xx status code
func (o *GetYourAccountChildrenInformationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get your account children information o k response has a 5xx status code
func (o *GetYourAccountChildrenInformationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get your account children information o k response a status code equal to that given
func (o *GetYourAccountChildrenInformationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get your account children information o k response
func (o *GetYourAccountChildrenInformationOK) Code() int {
	return 200
}

func (o *GetYourAccountChildrenInformationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /account-children][%d] getYourAccountChildrenInformationOK %s", 200, payload)
}

func (o *GetYourAccountChildrenInformationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /account-children][%d] getYourAccountChildrenInformationOK %s", 200, payload)
}

func (o *GetYourAccountChildrenInformationOK) GetPayload() *GetYourAccountChildrenInformationOKBody {
	return o.Payload
}

func (o *GetYourAccountChildrenInformationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetYourAccountChildrenInformationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetYourAccountChildrenInformationOKBody get your account children information o k body
// Example: {"data":[{"address":"84 Alexandra Square","company":null,"country":"UNITED_KINGDOM","email":"stephanie07@gmail.com","first_name":"Paul","id":19,"name":"wright.dylan.123","postcode":"WS11 1ZY","role":"SUPER_ENGINEER","standard_timezone":"Europe/London","surname":"Edwards","telephone_number":4527298086,"timezone":"GMT"},{"address":"39 Grace Ferry","company":null,"country":"UNITED_KINGDOM","email":"jmorgan@gmail.com","first_name":"Archie","id":20,"name":"michael37.776","postcode":"LE4 4JR","role":"OWNER","standard_timezone":"Atlantic/Azores","surname":"Jones","telephone_number":"0047 3552779","timezone":"GMT -1"}]}
swagger:model GetYourAccountChildrenInformationOKBody
*/
type GetYourAccountChildrenInformationOKBody struct {

	// data
	// Example: [{"address":"84 Alexandra Square","company":null,"country":"UNITED_KINGDOM","email":"stephanie07@gmail.com","first_name":"Paul","id":19,"name":"wright.dylan.123","postcode":"WS11 1ZY","role":"SUPER_ENGINEER","standard_timezone":"Europe/London","surname":"Edwards","telephone_number":4527298086,"timezone":"GMT"},{"address":"39 Grace Ferry","company":null,"country":"UNITED_KINGDOM","email":"jmorgan@gmail.com","first_name":"Archie","id":20,"name":"michael37.776","postcode":"LE4 4JR","role":"OWNER","standard_timezone":"Atlantic/Azores","surname":"Jones","telephone_number":"0047 3552779","timezone":"GMT -1"}]
	Data []*GetYourAccountChildrenInformationOKBodyDataItems0 `json:"data"`
}

// Validate validates this get your account children information o k body
func (o *GetYourAccountChildrenInformationOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetYourAccountChildrenInformationOKBody) validateData(formats strfmt.Registry) error {
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
					return ve.ValidateName("getYourAccountChildrenInformationOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getYourAccountChildrenInformationOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get your account children information o k body based on the context it is used
func (o *GetYourAccountChildrenInformationOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetYourAccountChildrenInformationOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getYourAccountChildrenInformationOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getYourAccountChildrenInformationOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetYourAccountChildrenInformationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetYourAccountChildrenInformationOKBody) UnmarshalBinary(b []byte) error {
	var res GetYourAccountChildrenInformationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetYourAccountChildrenInformationOKBodyDataItems0 get your account children information o k body data items0
swagger:model GetYourAccountChildrenInformationOKBodyDataItems0
*/
type GetYourAccountChildrenInformationOKBodyDataItems0 struct {

	// address
	// Example: 84 Alexandra Square
	Address string `json:"address,omitempty"`

	// company
	Company string `json:"company,omitempty"`

	// country
	// Example: UNITED_KINGDOM
	Country string `json:"country,omitempty"`

	// email
	// Example: stephanie07@gmail.com
	Email string `json:"email,omitempty"`

	// first name
	// Example: Paul
	FirstName string `json:"first_name,omitempty"`

	// id
	// Example: 19
	ID int64 `json:"id,omitempty"`

	// name
	// Example: wright.dylan.123
	Name string `json:"name,omitempty"`

	// postcode
	// Example: WS11 1ZY
	Postcode string `json:"postcode,omitempty"`

	// role
	// Example: SUPER_ENGINEER
	Role string `json:"role,omitempty"`

	// standard timezone
	// Example: Europe/London
	StandardTimezone string `json:"standard_timezone,omitempty"`

	// surname
	// Example: Edwards
	Surname string `json:"surname,omitempty"`

	// telephone number
	// Example: 4527298086
	TelephoneNumber string `json:"telephone_number,omitempty"`

	// timezone
	// Example: GMT
	Timezone string `json:"timezone,omitempty"`
}

// Validate validates this get your account children information o k body data items0
func (o *GetYourAccountChildrenInformationOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get your account children information o k body data items0 based on context it is used
func (o *GetYourAccountChildrenInformationOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetYourAccountChildrenInformationOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetYourAccountChildrenInformationOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res GetYourAccountChildrenInformationOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
