// Code generated by go-swagger; DO NOT EDIT.

package e_v_charger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetEVChargerByUUIDReader is a Reader for the GetEVChargerByUUID structure.
type GetEVChargerByUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEVChargerByUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEVChargerByUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ev-charger/{charger_uuid}] getEVChargerByUUID", response, response.Code())
	}
}

// NewGetEVChargerByUUIDOK creates a GetEVChargerByUUIDOK with default headers values
func NewGetEVChargerByUUIDOK() *GetEVChargerByUUIDOK {
	return &GetEVChargerByUUIDOK{}
}

/*
GetEVChargerByUUIDOK describes a response with status code 200, with default header values.

GetEVChargerByUUIDOK get e v charger by Uuid o k
*/
type GetEVChargerByUUIDOK struct {
	Payload *GetEVChargerByUUIDOKBody
}

// IsSuccess returns true when this get e v charger by Uuid o k response has a 2xx status code
func (o *GetEVChargerByUUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get e v charger by Uuid o k response has a 3xx status code
func (o *GetEVChargerByUUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get e v charger by Uuid o k response has a 4xx status code
func (o *GetEVChargerByUUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get e v charger by Uuid o k response has a 5xx status code
func (o *GetEVChargerByUUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get e v charger by Uuid o k response a status code equal to that given
func (o *GetEVChargerByUUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get e v charger by Uuid o k response
func (o *GetEVChargerByUUIDOK) Code() int {
	return 200
}

func (o *GetEVChargerByUUIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ev-charger/{charger_uuid}][%d] getEVChargerByUuidOK %s", 200, payload)
}

func (o *GetEVChargerByUUIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ev-charger/{charger_uuid}][%d] getEVChargerByUuidOK %s", 200, payload)
}

func (o *GetEVChargerByUUIDOK) GetPayload() *GetEVChargerByUUIDOKBody {
	return o.Payload
}

func (o *GetEVChargerByUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEVChargerByUUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetEVChargerByUUIDOKBody get e v charger by UUID o k body
// Example: {"data":{"alias":"Garage","online":false,"serial_number":"70458130637647","status":"Preparing","type":"GivEnergy","uuid":"9b0fa358-7f85-421e-b5cb-2ee5192e4a73","went_offline_at":null}}
swagger:model GetEVChargerByUUIDOKBody
*/
type GetEVChargerByUUIDOKBody struct {

	// data
	Data *GetEVChargerByUUIDOKBodyData `json:"data,omitempty"`
}

// Validate validates this get e v charger by UUID o k body
func (o *GetEVChargerByUUIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEVChargerByUUIDOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getEVChargerByUuidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getEVChargerByUuidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get e v charger by UUID o k body based on the context it is used
func (o *GetEVChargerByUUIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEVChargerByUUIDOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getEVChargerByUuidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getEVChargerByUuidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEVChargerByUUIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEVChargerByUUIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetEVChargerByUUIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetEVChargerByUUIDOKBodyData get e v charger by UUID o k body data
swagger:model GetEVChargerByUUIDOKBodyData
*/
type GetEVChargerByUUIDOKBodyData struct {

	// alias
	// Example: Garage
	Alias string `json:"alias,omitempty"`

	// online
	// Example: false
	Online bool `json:"online,omitempty"`

	// serial number
	// Example: 70458130637647
	SerialNumber string `json:"serial_number,omitempty"`

	// status
	// Example: Preparing
	Status string `json:"status,omitempty"`

	// type
	// Example: GivEnergy
	Type string `json:"type,omitempty"`

	// uuid
	// Example: 9b0fa358-7f85-421e-b5cb-2ee5192e4a73
	UUID string `json:"uuid,omitempty"`

	// went offline at
	WentOfflineAt string `json:"went_offline_at,omitempty"`
}

// Validate validates this get e v charger by UUID o k body data
func (o *GetEVChargerByUUIDOKBodyData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get e v charger by UUID o k body data based on context it is used
func (o *GetEVChargerByUUIDOKBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetEVChargerByUUIDOKBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEVChargerByUUIDOKBodyData) UnmarshalBinary(b []byte) error {
	var res GetEVChargerByUUIDOKBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
