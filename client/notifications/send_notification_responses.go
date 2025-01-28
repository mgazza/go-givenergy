// Code generated by go-swagger; DO NOT EDIT.

package notifications

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
	"github.com/go-openapi/validate"
)

// SendNotificationReader is a Reader for the SendNotification structure.
type SendNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSendNotificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSendNotificationOK creates a SendNotificationOK with default headers values
func NewSendNotificationOK() *SendNotificationOK {
	return &SendNotificationOK{}
}

/*
SendNotificationOK describes a response with status code 200, with default header values.

failure
*/
type SendNotificationOK struct {
	Payload *SendNotificationOKBody
}

// IsSuccess returns true when this send notification o k response has a 2xx status code
func (o *SendNotificationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this send notification o k response has a 3xx status code
func (o *SendNotificationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send notification o k response has a 4xx status code
func (o *SendNotificationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this send notification o k response has a 5xx status code
func (o *SendNotificationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this send notification o k response a status code equal to that given
func (o *SendNotificationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the send notification o k response
func (o *SendNotificationOK) Code() int {
	return 200
}

func (o *SendNotificationOK) Error() string {
	return fmt.Sprintf("[POST /notification/send][%d] sendNotificationOK  %+v", 200, o.Payload)
}

func (o *SendNotificationOK) String() string {
	return fmt.Sprintf("[POST /notification/send][%d] sendNotificationOK  %+v", 200, o.Payload)
}

func (o *SendNotificationOK) GetPayload() *SendNotificationOKBody {
	return o.Payload
}

func (o *SendNotificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SendNotificationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SendNotificationBody send notification body
swagger:model SendNotificationBody
*/
type SendNotificationBody struct {

	// Body of the notification. Must meet the API requirements for notification content. Must be at least 20 characters.
	// Example: Here is some information about the amazing thing that has happened
	// Required: true
	Body *string `json:"body"`

	// The notification's icon. Must be a valid [material design icon](https://pictogrammers.com/library/mdi/?welcome) prefixed with `mdi-`. Must start with one of <code>mdi-</code>.
	// Example: mdi-account-outline
	// Required: true
	Icon *string `json:"icon"`

	// One or more platforms to send the notification to.
	// Example: ["persist"]
	// Required: true
	Platforms []string `json:"platforms"`

	// Title of the notification. Must meet the API requirements for notification content. Must be at least 10 characters.
	// Example: Something amazing has happened!
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this send notification body
func (o *SendNotificationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIcon(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePlatforms(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SendNotificationBody) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"body", "body", o.Body); err != nil {
		return err
	}

	return nil
}

func (o *SendNotificationBody) validateIcon(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"icon", "body", o.Icon); err != nil {
		return err
	}

	return nil
}

func (o *SendNotificationBody) validatePlatforms(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"platforms", "body", o.Platforms); err != nil {
		return err
	}

	return nil
}

func (o *SendNotificationBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this send notification body based on context it is used
func (o *SendNotificationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SendNotificationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendNotificationBody) UnmarshalBinary(b []byte) error {
	var res SendNotificationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SendNotificationOKBody send notification o k body
// Example: {"success":false}
swagger:model SendNotificationOKBody
*/
type SendNotificationOKBody struct {

	// success
	// Example: false
	Success bool `json:"success,omitempty"`
}

// Validate validates this send notification o k body
func (o *SendNotificationOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this send notification o k body based on context it is used
func (o *SendNotificationOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SendNotificationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendNotificationOKBody) UnmarshalBinary(b []byte) error {
	var res SendNotificationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
