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
)

// GetSmartDeviceDataPointsByIDReader is a Reader for the GetSmartDeviceDataPointsByID structure.
type GetSmartDeviceDataPointsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSmartDeviceDataPointsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSmartDeviceDataPointsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSmartDeviceDataPointsByIDOK creates a GetSmartDeviceDataPointsByIDOK with default headers values
func NewGetSmartDeviceDataPointsByIDOK() *GetSmartDeviceDataPointsByIDOK {
	return &GetSmartDeviceDataPointsByIDOK{}
}

/*
GetSmartDeviceDataPointsByIDOK describes a response with status code 200, with default header values.

GetSmartDeviceDataPointsByIDOK get smart device data points by Id o k
*/
type GetSmartDeviceDataPointsByIDOK struct {
	Payload *GetSmartDeviceDataPointsByIDOKBody
}

// IsSuccess returns true when this get smart device data points by Id o k response has a 2xx status code
func (o *GetSmartDeviceDataPointsByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get smart device data points by Id o k response has a 3xx status code
func (o *GetSmartDeviceDataPointsByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get smart device data points by Id o k response has a 4xx status code
func (o *GetSmartDeviceDataPointsByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get smart device data points by Id o k response has a 5xx status code
func (o *GetSmartDeviceDataPointsByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get smart device data points by Id o k response a status code equal to that given
func (o *GetSmartDeviceDataPointsByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get smart device data points by Id o k response
func (o *GetSmartDeviceDataPointsByIDOK) Code() int {
	return 200
}

func (o *GetSmartDeviceDataPointsByIDOK) Error() string {
	return fmt.Sprintf("[GET /smart-device/{smartDevice_uuid}/data][%d] getSmartDeviceDataPointsByIdOK  %+v", 200, o.Payload)
}

func (o *GetSmartDeviceDataPointsByIDOK) String() string {
	return fmt.Sprintf("[GET /smart-device/{smartDevice_uuid}/data][%d] getSmartDeviceDataPointsByIdOK  %+v", 200, o.Payload)
}

func (o *GetSmartDeviceDataPointsByIDOK) GetPayload() *GetSmartDeviceDataPointsByIDOKBody {
	return o.Payload
}

func (o *GetSmartDeviceDataPointsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSmartDeviceDataPointsByIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetSmartDeviceDataPointsByIDOKBody get smart device data points by ID o k body
// Example: {"data":[{"power":509,"time":"1987-04-16T18:03:05Z"},{"power":431,"time":"2016-05-03T04:07:01Z"}]}
swagger:model GetSmartDeviceDataPointsByIDOKBody
*/
type GetSmartDeviceDataPointsByIDOKBody struct {

	// data
	// Example: [{"power":509,"time":"1987-04-16T18:03:05Z"},{"power":431,"time":"2016-05-03T04:07:01Z"}]
	Data []*GetSmartDeviceDataPointsByIDOKBodyDataItems0 `json:"data"`
}

// Validate validates this get smart device data points by ID o k body
func (o *GetSmartDeviceDataPointsByIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSmartDeviceDataPointsByIDOKBody) validateData(formats strfmt.Registry) error {
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
					return ve.ValidateName("getSmartDeviceDataPointsByIdOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getSmartDeviceDataPointsByIdOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get smart device data points by ID o k body based on the context it is used
func (o *GetSmartDeviceDataPointsByIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSmartDeviceDataPointsByIDOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getSmartDeviceDataPointsByIdOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getSmartDeviceDataPointsByIdOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSmartDeviceDataPointsByIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSmartDeviceDataPointsByIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetSmartDeviceDataPointsByIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetSmartDeviceDataPointsByIDOKBodyDataItems0 get smart device data points by ID o k body data items0
swagger:model GetSmartDeviceDataPointsByIDOKBodyDataItems0
*/
type GetSmartDeviceDataPointsByIDOKBodyDataItems0 struct {

	// power
	// Example: 509
	Power int64 `json:"power,omitempty"`

	// time
	// Example: 1987-04-16T18:03:05Z
	Time string `json:"time,omitempty"`
}

// Validate validates this get smart device data points by ID o k body data items0
func (o *GetSmartDeviceDataPointsByIDOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get smart device data points by ID o k body data items0 based on context it is used
func (o *GetSmartDeviceDataPointsByIDOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSmartDeviceDataPointsByIDOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSmartDeviceDataPointsByIDOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res GetSmartDeviceDataPointsByIDOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
