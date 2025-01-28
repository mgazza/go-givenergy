// Code generated by go-swagger; DO NOT EDIT.

package inverter_data

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

// GetLatestSystemDataReader is a Reader for the GetLatestSystemData structure.
type GetLatestSystemDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLatestSystemDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLatestSystemDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inverter/{inverter_serial_number}/system-data/latest] getLatestSystemData", response, response.Code())
	}
}

// NewGetLatestSystemDataOK creates a GetLatestSystemDataOK with default headers values
func NewGetLatestSystemDataOK() *GetLatestSystemDataOK {
	return &GetLatestSystemDataOK{}
}

/*
GetLatestSystemDataOK describes a response with status code 200, with default header values.

GetLatestSystemDataOK get latest system data o k
*/
type GetLatestSystemDataOK struct {
	Payload *GetLatestSystemDataOKBody
}

// IsSuccess returns true when this get latest system data o k response has a 2xx status code
func (o *GetLatestSystemDataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get latest system data o k response has a 3xx status code
func (o *GetLatestSystemDataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest system data o k response has a 4xx status code
func (o *GetLatestSystemDataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get latest system data o k response has a 5xx status code
func (o *GetLatestSystemDataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest system data o k response a status code equal to that given
func (o *GetLatestSystemDataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get latest system data o k response
func (o *GetLatestSystemDataOK) Code() int {
	return 200
}

func (o *GetLatestSystemDataOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /inverter/{inverter_serial_number}/system-data/latest][%d] getLatestSystemDataOK %s", 200, payload)
}

func (o *GetLatestSystemDataOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /inverter/{inverter_serial_number}/system-data/latest][%d] getLatestSystemDataOK %s", 200, payload)
}

func (o *GetLatestSystemDataOK) GetPayload() *GetLatestSystemDataOKBody {
	return o.Payload
}

func (o *GetLatestSystemDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLatestSystemDataOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetLatestSystemDataOKBody get latest system data o k body
// Example: {"data":{"battery":{"percent":83,"power":-2484,"temperature":85},"consumption":38674,"grid":{"current":-4688.5,"frequency":259.5,"power":-3270,"voltage":2155.3},"inverter":{"eps_power":224,"output_frequency":49.28,"output_voltage":222.1,"power":37241,"temperature":80},"solar":{"arrays":[{"array":1,"current":918.5,"power":881,"voltage":369.8},{"array":2,"current":654.2,"power":563,"voltage":413.1}],"power":1526},"status":"Normal","time":"2007-01-13T05:46:09Z"}}
swagger:model GetLatestSystemDataOKBody
*/
type GetLatestSystemDataOKBody struct {

	// data
	Data *GetLatestSystemDataOKBodyData `json:"data,omitempty"`
}

// Validate validates this get latest system data o k body
func (o *GetLatestSystemDataOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestSystemDataOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestSystemDataOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestSystemDataOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get latest system data o k body based on the context it is used
func (o *GetLatestSystemDataOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestSystemDataOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestSystemDataOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestSystemDataOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestSystemDataOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestSystemDataOKBody) UnmarshalBinary(b []byte) error {
	var res GetLatestSystemDataOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestSystemDataOKBodyData get latest system data o k body data
swagger:model GetLatestSystemDataOKBodyData
*/
type GetLatestSystemDataOKBodyData struct {

	// battery
	Battery *GetLatestSystemDataOKBodyDataBattery `json:"battery,omitempty"`

	// consumption
	// Example: 38674
	Consumption int64 `json:"consumption,omitempty"`

	// grid
	Grid *GetLatestSystemDataOKBodyDataGrid `json:"grid,omitempty"`

	// inverter
	Inverter *GetLatestSystemDataOKBodyDataInverter `json:"inverter,omitempty"`

	// solar
	Solar *GetLatestSystemDataOKBodyDataSolar `json:"solar,omitempty"`

	// status
	// Example: Normal
	Status string `json:"status,omitempty"`

	// time
	// Example: 2007-01-13T05:46:09Z
	Time string `json:"time,omitempty"`
}

// Validate validates this get latest system data o k body data
func (o *GetLatestSystemDataOKBodyData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBattery(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGrid(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInverter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSolar(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestSystemDataOKBodyData) validateBattery(formats strfmt.Registry) error {
	if swag.IsZero(o.Battery) { // not required
		return nil
	}

	if o.Battery != nil {
		if err := o.Battery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "battery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "battery")
			}
			return err
		}
	}

	return nil
}

func (o *GetLatestSystemDataOKBodyData) validateGrid(formats strfmt.Registry) error {
	if swag.IsZero(o.Grid) { // not required
		return nil
	}

	if o.Grid != nil {
		if err := o.Grid.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "grid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "grid")
			}
			return err
		}
	}

	return nil
}

func (o *GetLatestSystemDataOKBodyData) validateInverter(formats strfmt.Registry) error {
	if swag.IsZero(o.Inverter) { // not required
		return nil
	}

	if o.Inverter != nil {
		if err := o.Inverter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "inverter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "inverter")
			}
			return err
		}
	}

	return nil
}

func (o *GetLatestSystemDataOKBodyData) validateSolar(formats strfmt.Registry) error {
	if swag.IsZero(o.Solar) { // not required
		return nil
	}

	if o.Solar != nil {
		if err := o.Solar.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "solar")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "solar")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get latest system data o k body data based on the context it is used
func (o *GetLatestSystemDataOKBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBattery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateGrid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateInverter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSolar(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestSystemDataOKBodyData) contextValidateBattery(ctx context.Context, formats strfmt.Registry) error {

	if o.Battery != nil {

		if swag.IsZero(o.Battery) { // not required
			return nil
		}

		if err := o.Battery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "battery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "battery")
			}
			return err
		}
	}

	return nil
}

func (o *GetLatestSystemDataOKBodyData) contextValidateGrid(ctx context.Context, formats strfmt.Registry) error {

	if o.Grid != nil {

		if swag.IsZero(o.Grid) { // not required
			return nil
		}

		if err := o.Grid.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "grid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "grid")
			}
			return err
		}
	}

	return nil
}

func (o *GetLatestSystemDataOKBodyData) contextValidateInverter(ctx context.Context, formats strfmt.Registry) error {

	if o.Inverter != nil {

		if swag.IsZero(o.Inverter) { // not required
			return nil
		}

		if err := o.Inverter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "inverter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "inverter")
			}
			return err
		}
	}

	return nil
}

func (o *GetLatestSystemDataOKBodyData) contextValidateSolar(ctx context.Context, formats strfmt.Registry) error {

	if o.Solar != nil {

		if swag.IsZero(o.Solar) { // not required
			return nil
		}

		if err := o.Solar.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "solar")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "solar")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestSystemDataOKBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestSystemDataOKBodyData) UnmarshalBinary(b []byte) error {
	var res GetLatestSystemDataOKBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestSystemDataOKBodyDataBattery get latest system data o k body data battery
swagger:model GetLatestSystemDataOKBodyDataBattery
*/
type GetLatestSystemDataOKBodyDataBattery struct {

	// percent
	// Example: 83
	Percent int64 `json:"percent,omitempty"`

	// power
	// Example: -2484
	Power int64 `json:"power,omitempty"`

	// temperature
	// Example: 85
	Temperature float64 `json:"temperature,omitempty"`
}

// Validate validates this get latest system data o k body data battery
func (o *GetLatestSystemDataOKBodyDataBattery) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get latest system data o k body data battery based on context it is used
func (o *GetLatestSystemDataOKBodyDataBattery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestSystemDataOKBodyDataBattery) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestSystemDataOKBodyDataBattery) UnmarshalBinary(b []byte) error {
	var res GetLatestSystemDataOKBodyDataBattery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestSystemDataOKBodyDataGrid get latest system data o k body data grid
swagger:model GetLatestSystemDataOKBodyDataGrid
*/
type GetLatestSystemDataOKBodyDataGrid struct {

	// current
	// Example: -4688.5
	Current float64 `json:"current,omitempty"`

	// frequency
	// Example: 259.5
	Frequency float64 `json:"frequency,omitempty"`

	// power
	// Example: -3270
	Power int64 `json:"power,omitempty"`

	// voltage
	// Example: 2155.3
	Voltage float64 `json:"voltage,omitempty"`
}

// Validate validates this get latest system data o k body data grid
func (o *GetLatestSystemDataOKBodyDataGrid) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get latest system data o k body data grid based on context it is used
func (o *GetLatestSystemDataOKBodyDataGrid) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestSystemDataOKBodyDataGrid) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestSystemDataOKBodyDataGrid) UnmarshalBinary(b []byte) error {
	var res GetLatestSystemDataOKBodyDataGrid
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestSystemDataOKBodyDataInverter get latest system data o k body data inverter
swagger:model GetLatestSystemDataOKBodyDataInverter
*/
type GetLatestSystemDataOKBodyDataInverter struct {

	// eps power
	// Example: 224
	EpsPower int64 `json:"eps_power,omitempty"`

	// output frequency
	// Example: 49.28
	OutputFrequency float64 `json:"output_frequency,omitempty"`

	// output voltage
	// Example: 222.1
	OutputVoltage float64 `json:"output_voltage,omitempty"`

	// power
	// Example: 37241
	Power int64 `json:"power,omitempty"`

	// temperature
	// Example: 80
	Temperature float64 `json:"temperature,omitempty"`
}

// Validate validates this get latest system data o k body data inverter
func (o *GetLatestSystemDataOKBodyDataInverter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get latest system data o k body data inverter based on context it is used
func (o *GetLatestSystemDataOKBodyDataInverter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestSystemDataOKBodyDataInverter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestSystemDataOKBodyDataInverter) UnmarshalBinary(b []byte) error {
	var res GetLatestSystemDataOKBodyDataInverter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestSystemDataOKBodyDataSolar get latest system data o k body data solar
swagger:model GetLatestSystemDataOKBodyDataSolar
*/
type GetLatestSystemDataOKBodyDataSolar struct {

	// arrays
	// Example: [{"array":1,"current":918.5,"power":881,"voltage":369.8},{"array":2,"current":654.2,"power":563,"voltage":413.1}]
	Arrays []*GetLatestSystemDataOKBodyDataSolarArraysItems0 `json:"arrays"`

	// power
	// Example: 1526
	Power int64 `json:"power,omitempty"`
}

// Validate validates this get latest system data o k body data solar
func (o *GetLatestSystemDataOKBodyDataSolar) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateArrays(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestSystemDataOKBodyDataSolar) validateArrays(formats strfmt.Registry) error {
	if swag.IsZero(o.Arrays) { // not required
		return nil
	}

	for i := 0; i < len(o.Arrays); i++ {
		if swag.IsZero(o.Arrays[i]) { // not required
			continue
		}

		if o.Arrays[i] != nil {
			if err := o.Arrays[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "solar" + "." + "arrays" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "solar" + "." + "arrays" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get latest system data o k body data solar based on the context it is used
func (o *GetLatestSystemDataOKBodyDataSolar) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateArrays(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestSystemDataOKBodyDataSolar) contextValidateArrays(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Arrays); i++ {

		if o.Arrays[i] != nil {

			if swag.IsZero(o.Arrays[i]) { // not required
				return nil
			}

			if err := o.Arrays[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "solar" + "." + "arrays" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getLatestSystemDataOK" + "." + "data" + "." + "solar" + "." + "arrays" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestSystemDataOKBodyDataSolar) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestSystemDataOKBodyDataSolar) UnmarshalBinary(b []byte) error {
	var res GetLatestSystemDataOKBodyDataSolar
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestSystemDataOKBodyDataSolarArraysItems0 get latest system data o k body data solar arrays items0
swagger:model GetLatestSystemDataOKBodyDataSolarArraysItems0
*/
type GetLatestSystemDataOKBodyDataSolarArraysItems0 struct {

	// array
	// Example: 1
	Array int64 `json:"array,omitempty"`

	// current
	// Example: 918.5
	Current float64 `json:"current,omitempty"`

	// power
	// Example: 881
	Power int64 `json:"power,omitempty"`

	// voltage
	// Example: 369.8
	Voltage float64 `json:"voltage,omitempty"`
}

// Validate validates this get latest system data o k body data solar arrays items0
func (o *GetLatestSystemDataOKBodyDataSolarArraysItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get latest system data o k body data solar arrays items0 based on context it is used
func (o *GetLatestSystemDataOKBodyDataSolarArraysItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestSystemDataOKBodyDataSolarArraysItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestSystemDataOKBodyDataSolarArraysItems0) UnmarshalBinary(b []byte) error {
	var res GetLatestSystemDataOKBodyDataSolarArraysItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
