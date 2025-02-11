// Code generated by go-swagger; DO NOT EDIT.

package smart_device

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

// GetYourSmartDevicesReader is a Reader for the GetYourSmartDevices structure.
type GetYourSmartDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetYourSmartDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetYourSmartDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /smart-device] getYourSmartDevices", response, response.Code())
	}
}

// NewGetYourSmartDevicesOK creates a GetYourSmartDevicesOK with default headers values
func NewGetYourSmartDevicesOK() *GetYourSmartDevicesOK {
	return &GetYourSmartDevicesOK{}
}

/*
GetYourSmartDevicesOK describes a response with status code 200, with default header values.

GetYourSmartDevicesOK get your smart devices o k
*/
type GetYourSmartDevicesOK struct {
	Payload *GetYourSmartDevicesOKBody
}

// IsSuccess returns true when this get your smart devices o k response has a 2xx status code
func (o *GetYourSmartDevicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get your smart devices o k response has a 3xx status code
func (o *GetYourSmartDevicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get your smart devices o k response has a 4xx status code
func (o *GetYourSmartDevicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get your smart devices o k response has a 5xx status code
func (o *GetYourSmartDevicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get your smart devices o k response a status code equal to that given
func (o *GetYourSmartDevicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get your smart devices o k response
func (o *GetYourSmartDevicesOK) Code() int {
	return 200
}

func (o *GetYourSmartDevicesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /smart-device][%d] getYourSmartDevicesOK %s", 200, payload)
}

func (o *GetYourSmartDevicesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /smart-device][%d] getYourSmartDevicesOK %s", 200, payload)
}

func (o *GetYourSmartDevicesOK) GetPayload() *GetYourSmartDevicesOKBody {
	return o.Payload
}

func (o *GetYourSmartDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetYourSmartDevicesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetYourSmartDevicesOKBody get your smart devices o k body
// Example: {"data":[{"alias":"veniam","other_data":{"graph_color":"#f0c27c","hardware_id":"P","local_key":"J"},"uuid":"0d8968ca-f418-4bbd-af21-2bde7613614b"},{"alias":"nostrum","other_data":{"graph_color":"#608742","hardware_id":"_","local_key":")"},"uuid":"d6032cf8-cbe7-4fbb-a710-c63fd9f6dae6"}]}
swagger:model GetYourSmartDevicesOKBody
*/
type GetYourSmartDevicesOKBody struct {

	// data
	// Example: [{"alias":"veniam","other_data":{"graph_color":"#f0c27c","hardware_id":"P","local_key":"J"},"uuid":"0d8968ca-f418-4bbd-af21-2bde7613614b"},{"alias":"nostrum","other_data":{"graph_color":"#608742","hardware_id":"_","local_key":")"},"uuid":"d6032cf8-cbe7-4fbb-a710-c63fd9f6dae6"}]
	Data []*GetYourSmartDevicesOKBodyDataItems0 `json:"data"`
}

// Validate validates this get your smart devices o k body
func (o *GetYourSmartDevicesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetYourSmartDevicesOKBody) validateData(formats strfmt.Registry) error {
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
					return ve.ValidateName("getYourSmartDevicesOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getYourSmartDevicesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get your smart devices o k body based on the context it is used
func (o *GetYourSmartDevicesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetYourSmartDevicesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getYourSmartDevicesOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getYourSmartDevicesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetYourSmartDevicesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetYourSmartDevicesOKBody) UnmarshalBinary(b []byte) error {
	var res GetYourSmartDevicesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetYourSmartDevicesOKBodyDataItems0 get your smart devices o k body data items0
swagger:model GetYourSmartDevicesOKBodyDataItems0
*/
type GetYourSmartDevicesOKBodyDataItems0 struct {

	// alias
	// Example: veniam
	Alias string `json:"alias,omitempty"`

	// other data
	OtherData *GetYourSmartDevicesOKBodyDataItems0OtherData `json:"other_data,omitempty"`

	// uuid
	// Example: 0d8968ca-f418-4bbd-af21-2bde7613614b
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this get your smart devices o k body data items0
func (o *GetYourSmartDevicesOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOtherData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetYourSmartDevicesOKBodyDataItems0) validateOtherData(formats strfmt.Registry) error {
	if swag.IsZero(o.OtherData) { // not required
		return nil
	}

	if o.OtherData != nil {
		if err := o.OtherData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("other_data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("other_data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get your smart devices o k body data items0 based on the context it is used
func (o *GetYourSmartDevicesOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOtherData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetYourSmartDevicesOKBodyDataItems0) contextValidateOtherData(ctx context.Context, formats strfmt.Registry) error {

	if o.OtherData != nil {

		if swag.IsZero(o.OtherData) { // not required
			return nil
		}

		if err := o.OtherData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("other_data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("other_data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetYourSmartDevicesOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetYourSmartDevicesOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res GetYourSmartDevicesOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetYourSmartDevicesOKBodyDataItems0OtherData get your smart devices o k body data items0 other data
swagger:model GetYourSmartDevicesOKBodyDataItems0OtherData
*/
type GetYourSmartDevicesOKBodyDataItems0OtherData struct {

	// graph color
	// Example: #f0c27c
	GraphColor string `json:"graph_color,omitempty"`

	// hardware id
	// Example: P
	HardwareID string `json:"hardware_id,omitempty"`

	// local key
	// Example: J
	LocalKey string `json:"local_key,omitempty"`
}

// Validate validates this get your smart devices o k body data items0 other data
func (o *GetYourSmartDevicesOKBodyDataItems0OtherData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get your smart devices o k body data items0 other data based on context it is used
func (o *GetYourSmartDevicesOKBodyDataItems0OtherData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetYourSmartDevicesOKBodyDataItems0OtherData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetYourSmartDevicesOKBodyDataItems0OtherData) UnmarshalBinary(b []byte) error {
	var res GetYourSmartDevicesOKBodyDataItems0OtherData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
