// Code generated by go-swagger; DO NOT EDIT.

package inverter_control

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
	"github.com/go-openapi/validate"
)

// ReadSettingMultipleReader is a Reader for the ReadSettingMultiple structure.
type ReadSettingMultipleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadSettingMultipleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadSettingMultipleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /multi-control/read] readSettingMultiple", response, response.Code())
	}
}

// NewReadSettingMultipleOK creates a ReadSettingMultipleOK with default headers values
func NewReadSettingMultipleOK() *ReadSettingMultipleOK {
	return &ReadSettingMultipleOK{}
}

/*
ReadSettingMultipleOK describes a response with status code 200, with default header values.

ReadSettingMultipleOK read setting multiple o k
*/
type ReadSettingMultipleOK struct {
	Payload string
}

// IsSuccess returns true when this read setting multiple o k response has a 2xx status code
func (o *ReadSettingMultipleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read setting multiple o k response has a 3xx status code
func (o *ReadSettingMultipleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read setting multiple o k response has a 4xx status code
func (o *ReadSettingMultipleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read setting multiple o k response has a 5xx status code
func (o *ReadSettingMultipleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read setting multiple o k response a status code equal to that given
func (o *ReadSettingMultipleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read setting multiple o k response
func (o *ReadSettingMultipleOK) Code() int {
	return 200
}

func (o *ReadSettingMultipleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /multi-control/read][%d] readSettingMultipleOK %s", 200, payload)
}

func (o *ReadSettingMultipleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /multi-control/read][%d] readSettingMultipleOK %s", 200, payload)
}

func (o *ReadSettingMultipleOK) GetPayload() string {
	return o.Payload
}

func (o *ReadSettingMultipleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ReadSettingMultipleBody read setting multiple body
swagger:model ReadSettingMultipleBody
*/
type ReadSettingMultipleBody struct {

	// An array of inverter serial numbers to send the command to with a maximum length of 1000
	// Example: ["SERIAL_1","SERIAL_2"]
	// Required: true
	InverterSerials []string `json:"inverter_serials"`

	// The ID of the setting to modify
	// Example: 17
	// Required: true
	SettingID *int64 `json:"setting_id"`
}

// Validate validates this read setting multiple body
func (o *ReadSettingMultipleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInverterSerials(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSettingID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReadSettingMultipleBody) validateInverterSerials(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"inverter_serials", "body", o.InverterSerials); err != nil {
		return err
	}

	return nil
}

func (o *ReadSettingMultipleBody) validateSettingID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"setting_id", "body", o.SettingID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this read setting multiple body based on context it is used
func (o *ReadSettingMultipleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ReadSettingMultipleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReadSettingMultipleBody) UnmarshalBinary(b []byte) error {
	var res ReadSettingMultipleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
