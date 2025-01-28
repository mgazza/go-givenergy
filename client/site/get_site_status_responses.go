// Code generated by go-swagger; DO NOT EDIT.

package site

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetSiteStatusReader is a Reader for the GetSiteStatus structure.
type GetSiteStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSiteStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSiteStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSiteStatusOK creates a GetSiteStatusOK with default headers values
func NewGetSiteStatusOK() *GetSiteStatusOK {
	return &GetSiteStatusOK{}
}

/*
GetSiteStatusOK describes a response with status code 200, with default header values.

GetSiteStatusOK get site status o k
*/
type GetSiteStatusOK struct {
	Payload *GetSiteStatusOKBody
}

// IsSuccess returns true when this get site status o k response has a 2xx status code
func (o *GetSiteStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get site status o k response has a 3xx status code
func (o *GetSiteStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get site status o k response has a 4xx status code
func (o *GetSiteStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get site status o k response has a 5xx status code
func (o *GetSiteStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get site status o k response a status code equal to that given
func (o *GetSiteStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get site status o k response
func (o *GetSiteStatusOK) Code() int {
	return 200
}

func (o *GetSiteStatusOK) Error() string {
	return fmt.Sprintf("[GET /site/{site_plant_id}/status][%d] getSiteStatusOK  %+v", 200, o.Payload)
}

func (o *GetSiteStatusOK) String() string {
	return fmt.Sprintf("[GET /site/{site_plant_id}/status][%d] getSiteStatusOK  %+v", 200, o.Payload)
}

func (o *GetSiteStatusOK) GetPayload() *GetSiteStatusOKBody {
	return o.Payload
}

func (o *GetSiteStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSiteStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetSiteStatusOKBody get site status o k body
// Example: {"status":"vacant"}
swagger:model GetSiteStatusOKBody
*/
type GetSiteStatusOKBody struct {

	// status
	// Example: vacant
	Status string `json:"status,omitempty"`
}

// Validate validates this get site status o k body
func (o *GetSiteStatusOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get site status o k body based on context it is used
func (o *GetSiteStatusOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSiteStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSiteStatusOKBody) UnmarshalBinary(b []byte) error {
	var res GetSiteStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
