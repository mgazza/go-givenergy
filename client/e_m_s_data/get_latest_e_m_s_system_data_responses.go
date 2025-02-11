// Code generated by go-swagger; DO NOT EDIT.

package e_m_s_data

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

// GetLatestEMSSystemDataReader is a Reader for the GetLatestEMSSystemData structure.
type GetLatestEMSSystemDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLatestEMSSystemDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLatestEMSSystemDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ems/{inverter_serial_number}/system-data/latest] getLatestEMSSystemData", response, response.Code())
	}
}

// NewGetLatestEMSSystemDataOK creates a GetLatestEMSSystemDataOK with default headers values
func NewGetLatestEMSSystemDataOK() *GetLatestEMSSystemDataOK {
	return &GetLatestEMSSystemDataOK{}
}

/*
GetLatestEMSSystemDataOK describes a response with status code 200, with default header values.

GetLatestEMSSystemDataOK get latest e m s system data o k
*/
type GetLatestEMSSystemDataOK struct {
	Payload *GetLatestEMSSystemDataOKBody
}

// IsSuccess returns true when this get latest e m s system data o k response has a 2xx status code
func (o *GetLatestEMSSystemDataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get latest e m s system data o k response has a 3xx status code
func (o *GetLatestEMSSystemDataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest e m s system data o k response has a 4xx status code
func (o *GetLatestEMSSystemDataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get latest e m s system data o k response has a 5xx status code
func (o *GetLatestEMSSystemDataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest e m s system data o k response a status code equal to that given
func (o *GetLatestEMSSystemDataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get latest e m s system data o k response
func (o *GetLatestEMSSystemDataOK) Code() int {
	return 200
}

func (o *GetLatestEMSSystemDataOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ems/{inverter_serial_number}/system-data/latest][%d] getLatestEMSSystemDataOK %s", 200, payload)
}

func (o *GetLatestEMSSystemDataOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ems/{inverter_serial_number}/system-data/latest][%d] getLatestEMSSystemDataOK %s", 200, payload)
}

func (o *GetLatestEMSSystemDataOK) GetPayload() *GetLatestEMSSystemDataOKBody {
	return o.Payload
}

func (o *GetLatestEMSSystemDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLatestEMSSystemDataOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetLatestEMSSystemDataOKBody get latest e m s system data o k body
// Example: {"data":{"battery_power":9214,"battery_wh_remaining":58276,"calculated_load_power":15168,"car_power":-27828,"error_code":40330,"grid_power":16772,"inverters":[{"active_power":-8597,"number":1,"serial_number":"il6029p206","soc":88,"status":4,"temperature":3.2},{"active_power":-2221,"number":2,"serial_number":"nl3168l361","soc":79,"status":4,"temperature":10.3},{"active_power":24191,"number":3,"serial_number":"uy0262d251","soc":43,"status":2,"temperature":28.4},{"active_power":null,"number":4,"serial_number":null,"soc":null,"status":0,"temperature":null}],"measured_load_power":19053,"meters":[{"active_power":null,"number":1,"status":0,"type":0},{"active_power":null,"number":2,"status":2,"type":1},{"active_power":null,"number":3,"status":0,"type":3},{"active_power":null,"number":4,"status":2,"type":2},{"active_power":null,"number":5,"status":0,"type":0},{"active_power":null,"number":6,"status":0,"type":0},{"active_power":null,"number":7,"status":0,"type":0},{"active_power":null,"number":8,"status":0,"type":0}],"status":3,"time":null,"total_generation":8035,"warning_code":48384}}
swagger:model GetLatestEMSSystemDataOKBody
*/
type GetLatestEMSSystemDataOKBody struct {

	// data
	Data *GetLatestEMSSystemDataOKBodyData `json:"data,omitempty"`
}

// Validate validates this get latest e m s system data o k body
func (o *GetLatestEMSSystemDataOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestEMSSystemDataOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEMSSystemDataOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEMSSystemDataOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get latest e m s system data o k body based on the context it is used
func (o *GetLatestEMSSystemDataOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestEMSSystemDataOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEMSSystemDataOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEMSSystemDataOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestEMSSystemDataOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestEMSSystemDataOKBody) UnmarshalBinary(b []byte) error {
	var res GetLatestEMSSystemDataOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestEMSSystemDataOKBodyData get latest e m s system data o k body data
swagger:model GetLatestEMSSystemDataOKBodyData
*/
type GetLatestEMSSystemDataOKBodyData struct {

	// battery power
	// Example: 9214
	BatteryPower int64 `json:"battery_power,omitempty"`

	// battery wh remaining
	// Example: 58276
	BatteryWhRemaining int64 `json:"battery_wh_remaining,omitempty"`

	// calculated load power
	// Example: 15168
	CalculatedLoadPower int64 `json:"calculated_load_power,omitempty"`

	// car power
	// Example: -27828
	CarPower int64 `json:"car_power,omitempty"`

	// error code
	// Example: 40330
	ErrorCode int64 `json:"error_code,omitempty"`

	// grid power
	// Example: 16772
	GridPower int64 `json:"grid_power,omitempty"`

	// inverters
	// Example: [{"active_power":-8597,"number":1,"serial_number":"il6029p206","soc":88,"status":4,"temperature":3.2},{"active_power":-2221,"number":2,"serial_number":"nl3168l361","soc":79,"status":4,"temperature":10.3},{"active_power":24191,"number":3,"serial_number":"uy0262d251","soc":43,"status":2,"temperature":28.4},{"active_power":null,"number":4,"serial_number":null,"soc":null,"status":0,"temperature":null}]
	Inverters []*GetLatestEMSSystemDataOKBodyDataInvertersItems0 `json:"inverters"`

	// measured load power
	// Example: 19053
	MeasuredLoadPower int64 `json:"measured_load_power,omitempty"`

	// meters
	// Example: [{"active_power":null,"number":1,"status":0,"type":0},{"active_power":null,"number":2,"status":2,"type":1},{"active_power":null,"number":3,"status":0,"type":3},{"active_power":null,"number":4,"status":2,"type":2},{"active_power":null,"number":5,"status":0,"type":0},{"active_power":null,"number":6,"status":0,"type":0},{"active_power":null,"number":7,"status":0,"type":0},{"active_power":null,"number":8,"status":0,"type":0}]
	Meters []*GetLatestEMSSystemDataOKBodyDataMetersItems0 `json:"meters"`

	// status
	// Example: 3
	Status int64 `json:"status,omitempty"`

	// time
	Time string `json:"time,omitempty"`

	// total generation
	// Example: 8035
	TotalGeneration int64 `json:"total_generation,omitempty"`

	// warning code
	// Example: 48384
	WarningCode int64 `json:"warning_code,omitempty"`
}

// Validate validates this get latest e m s system data o k body data
func (o *GetLatestEMSSystemDataOKBodyData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInverters(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMeters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestEMSSystemDataOKBodyData) validateInverters(formats strfmt.Registry) error {
	if swag.IsZero(o.Inverters) { // not required
		return nil
	}

	for i := 0; i < len(o.Inverters); i++ {
		if swag.IsZero(o.Inverters[i]) { // not required
			continue
		}

		if o.Inverters[i] != nil {
			if err := o.Inverters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLatestEMSSystemDataOK" + "." + "data" + "." + "inverters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getLatestEMSSystemDataOK" + "." + "data" + "." + "inverters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetLatestEMSSystemDataOKBodyData) validateMeters(formats strfmt.Registry) error {
	if swag.IsZero(o.Meters) { // not required
		return nil
	}

	for i := 0; i < len(o.Meters); i++ {
		if swag.IsZero(o.Meters[i]) { // not required
			continue
		}

		if o.Meters[i] != nil {
			if err := o.Meters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLatestEMSSystemDataOK" + "." + "data" + "." + "meters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getLatestEMSSystemDataOK" + "." + "data" + "." + "meters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get latest e m s system data o k body data based on the context it is used
func (o *GetLatestEMSSystemDataOKBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateInverters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMeters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestEMSSystemDataOKBodyData) contextValidateInverters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Inverters); i++ {

		if o.Inverters[i] != nil {

			if swag.IsZero(o.Inverters[i]) { // not required
				return nil
			}

			if err := o.Inverters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLatestEMSSystemDataOK" + "." + "data" + "." + "inverters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getLatestEMSSystemDataOK" + "." + "data" + "." + "inverters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetLatestEMSSystemDataOKBodyData) contextValidateMeters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Meters); i++ {

		if o.Meters[i] != nil {

			if swag.IsZero(o.Meters[i]) { // not required
				return nil
			}

			if err := o.Meters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLatestEMSSystemDataOK" + "." + "data" + "." + "meters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getLatestEMSSystemDataOK" + "." + "data" + "." + "meters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestEMSSystemDataOKBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestEMSSystemDataOKBodyData) UnmarshalBinary(b []byte) error {
	var res GetLatestEMSSystemDataOKBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestEMSSystemDataOKBodyDataInvertersItems0 get latest e m s system data o k body data inverters items0
swagger:model GetLatestEMSSystemDataOKBodyDataInvertersItems0
*/
type GetLatestEMSSystemDataOKBodyDataInvertersItems0 struct {

	// active power
	// Example: -8597
	ActivePower int64 `json:"active_power,omitempty"`

	// number
	// Example: 1
	Number int64 `json:"number,omitempty"`

	// serial number
	// Example: il6029p206
	SerialNumber string `json:"serial_number,omitempty"`

	// soc
	// Example: 88
	Soc int64 `json:"soc,omitempty"`

	// status
	// Example: 4
	Status int64 `json:"status,omitempty"`

	// temperature
	// Example: 3.2
	Temperature float64 `json:"temperature,omitempty"`
}

// Validate validates this get latest e m s system data o k body data inverters items0
func (o *GetLatestEMSSystemDataOKBodyDataInvertersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get latest e m s system data o k body data inverters items0 based on context it is used
func (o *GetLatestEMSSystemDataOKBodyDataInvertersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestEMSSystemDataOKBodyDataInvertersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestEMSSystemDataOKBodyDataInvertersItems0) UnmarshalBinary(b []byte) error {
	var res GetLatestEMSSystemDataOKBodyDataInvertersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestEMSSystemDataOKBodyDataMetersItems0 get latest e m s system data o k body data meters items0
swagger:model GetLatestEMSSystemDataOKBodyDataMetersItems0
*/
type GetLatestEMSSystemDataOKBodyDataMetersItems0 struct {

	// active power
	ActivePower string `json:"active_power,omitempty"`

	// number
	// Example: 1
	Number int64 `json:"number,omitempty"`

	// status
	// Example: 0
	Status int64 `json:"status,omitempty"`

	// type
	// Example: 0
	Type int64 `json:"type,omitempty"`
}

// Validate validates this get latest e m s system data o k body data meters items0
func (o *GetLatestEMSSystemDataOKBodyDataMetersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get latest e m s system data o k body data meters items0 based on context it is used
func (o *GetLatestEMSSystemDataOKBodyDataMetersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestEMSSystemDataOKBodyDataMetersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestEMSSystemDataOKBodyDataMetersItems0) UnmarshalBinary(b []byte) error {
	var res GetLatestEMSSystemDataOKBodyDataMetersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
