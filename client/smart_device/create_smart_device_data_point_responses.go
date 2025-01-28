// Code generated by go-swagger; DO NOT EDIT.

package smart_device

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
	"github.com/go-openapi/validate"
)

// CreateSmartDeviceDataPointReader is a Reader for the CreateSmartDeviceDataPoint structure.
type CreateSmartDeviceDataPointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSmartDeviceDataPointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSmartDeviceDataPointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSmartDeviceDataPointOK creates a CreateSmartDeviceDataPointOK with default headers values
func NewCreateSmartDeviceDataPointOK() *CreateSmartDeviceDataPointOK {
	return &CreateSmartDeviceDataPointOK{}
}

/*
CreateSmartDeviceDataPointOK describes a response with status code 200, with default header values.

CreateSmartDeviceDataPointOK create smart device data point o k
*/
type CreateSmartDeviceDataPointOK struct {
	Payload *CreateSmartDeviceDataPointOKBody
}

// IsSuccess returns true when this create smart device data point o k response has a 2xx status code
func (o *CreateSmartDeviceDataPointOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create smart device data point o k response has a 3xx status code
func (o *CreateSmartDeviceDataPointOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create smart device data point o k response has a 4xx status code
func (o *CreateSmartDeviceDataPointOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create smart device data point o k response has a 5xx status code
func (o *CreateSmartDeviceDataPointOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create smart device data point o k response a status code equal to that given
func (o *CreateSmartDeviceDataPointOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create smart device data point o k response
func (o *CreateSmartDeviceDataPointOK) Code() int {
	return 200
}

func (o *CreateSmartDeviceDataPointOK) Error() string {
	return fmt.Sprintf("[POST /smart-device/{smartDevice_uuid}/data][%d] createSmartDeviceDataPointOK  %+v", 200, o.Payload)
}

func (o *CreateSmartDeviceDataPointOK) String() string {
	return fmt.Sprintf("[POST /smart-device/{smartDevice_uuid}/data][%d] createSmartDeviceDataPointOK  %+v", 200, o.Payload)
}

func (o *CreateSmartDeviceDataPointOK) GetPayload() *CreateSmartDeviceDataPointOKBody {
	return o.Payload
}

func (o *CreateSmartDeviceDataPointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateSmartDeviceDataPointOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateSmartDeviceDataPointBody create smart device data point body
swagger:model CreateSmartDeviceDataPointBody
*/
type CreateSmartDeviceDataPointBody struct {

	// power
	// Example: 17
	// Required: true
	Power *int64 `json:"power"`

	// The date & time that the data point was created
	// Example: consequatur
	// Required: true
	Time *string `json:"time"`
}

// Validate validates this create smart device data point body
func (o *CreateSmartDeviceDataPointBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePower(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateSmartDeviceDataPointBody) validatePower(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"power", "body", o.Power); err != nil {
		return err
	}

	return nil
}

func (o *CreateSmartDeviceDataPointBody) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"time", "body", o.Time); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create smart device data point body based on context it is used
func (o *CreateSmartDeviceDataPointBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateSmartDeviceDataPointBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateSmartDeviceDataPointBody) UnmarshalBinary(b []byte) error {
	var res CreateSmartDeviceDataPointBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateSmartDeviceDataPointOKBody create smart device data point o k body
// Example: {"data":[{"power":622,"time":"1980-04-12T13:41:24Z"},{"power":437,"time":"1993-12-21T13:55:59Z"}]}
swagger:model CreateSmartDeviceDataPointOKBody
*/
type CreateSmartDeviceDataPointOKBody struct {

	// data
	// Example: [{"power":622,"time":"1980-04-12T13:41:24Z"},{"power":437,"time":"1993-12-21T13:55:59Z"}]
	Data []*CreateSmartDeviceDataPointOKBodyDataItems0 `json:"data"`
}

// Validate validates this create smart device data point o k body
func (o *CreateSmartDeviceDataPointOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateSmartDeviceDataPointOKBody) validateData(formats strfmt.Registry) error {
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
					return ve.ValidateName("createSmartDeviceDataPointOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createSmartDeviceDataPointOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create smart device data point o k body based on the context it is used
func (o *CreateSmartDeviceDataPointOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateSmartDeviceDataPointOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createSmartDeviceDataPointOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createSmartDeviceDataPointOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateSmartDeviceDataPointOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateSmartDeviceDataPointOKBody) UnmarshalBinary(b []byte) error {
	var res CreateSmartDeviceDataPointOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateSmartDeviceDataPointOKBodyDataItems0 create smart device data point o k body data items0
swagger:model CreateSmartDeviceDataPointOKBodyDataItems0
*/
type CreateSmartDeviceDataPointOKBodyDataItems0 struct {

	// power
	// Example: 622
	Power int64 `json:"power,omitempty"`

	// time
	// Example: 1980-04-12T13:41:24Z
	Time string `json:"time,omitempty"`
}

// Validate validates this create smart device data point o k body data items0
func (o *CreateSmartDeviceDataPointOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create smart device data point o k body data items0 based on context it is used
func (o *CreateSmartDeviceDataPointOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateSmartDeviceDataPointOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateSmartDeviceDataPointOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res CreateSmartDeviceDataPointOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
