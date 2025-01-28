// Code generated by go-swagger; DO NOT EDIT.

package meter

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

// GetHistoricMeterDataReader is a Reader for the GetHistoricMeterData structure.
type GetHistoricMeterDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHistoricMeterDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHistoricMeterDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHistoricMeterDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHistoricMeterDataOK creates a GetHistoricMeterDataOK with default headers values
func NewGetHistoricMeterDataOK() *GetHistoricMeterDataOK {
	return &GetHistoricMeterDataOK{}
}

/*
GetHistoricMeterDataOK describes a response with status code 200, with default header values.

GetHistoricMeterDataOK get historic meter data o k
*/
type GetHistoricMeterDataOK struct {
	Payload *GetHistoricMeterDataOKBody
}

// IsSuccess returns true when this get historic meter data o k response has a 2xx status code
func (o *GetHistoricMeterDataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get historic meter data o k response has a 3xx status code
func (o *GetHistoricMeterDataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get historic meter data o k response has a 4xx status code
func (o *GetHistoricMeterDataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get historic meter data o k response has a 5xx status code
func (o *GetHistoricMeterDataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get historic meter data o k response a status code equal to that given
func (o *GetHistoricMeterDataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get historic meter data o k response
func (o *GetHistoricMeterDataOK) Code() int {
	return 200
}

func (o *GetHistoricMeterDataOK) Error() string {
	return fmt.Sprintf("[GET /inverter/{inverter_serial_number}/meter/data][%d] getHistoricMeterDataOK  %+v", 200, o.Payload)
}

func (o *GetHistoricMeterDataOK) String() string {
	return fmt.Sprintf("[GET /inverter/{inverter_serial_number}/meter/data][%d] getHistoricMeterDataOK  %+v", 200, o.Payload)
}

func (o *GetHistoricMeterDataOK) GetPayload() *GetHistoricMeterDataOKBody {
	return o.Payload
}

func (o *GetHistoricMeterDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetHistoricMeterDataOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHistoricMeterDataNotFound creates a GetHistoricMeterDataNotFound with default headers values
func NewGetHistoricMeterDataNotFound() *GetHistoricMeterDataNotFound {
	return &GetHistoricMeterDataNotFound{}
}

/*
GetHistoricMeterDataNotFound describes a response with status code 404, with default header values.

GetHistoricMeterDataNotFound get historic meter data not found
*/
type GetHistoricMeterDataNotFound struct {
	Payload *GetHistoricMeterDataNotFoundBody
}

// IsSuccess returns true when this get historic meter data not found response has a 2xx status code
func (o *GetHistoricMeterDataNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get historic meter data not found response has a 3xx status code
func (o *GetHistoricMeterDataNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get historic meter data not found response has a 4xx status code
func (o *GetHistoricMeterDataNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get historic meter data not found response has a 5xx status code
func (o *GetHistoricMeterDataNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get historic meter data not found response a status code equal to that given
func (o *GetHistoricMeterDataNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get historic meter data not found response
func (o *GetHistoricMeterDataNotFound) Code() int {
	return 404
}

func (o *GetHistoricMeterDataNotFound) Error() string {
	return fmt.Sprintf("[GET /inverter/{inverter_serial_number}/meter/data][%d] getHistoricMeterDataNotFound  %+v", 404, o.Payload)
}

func (o *GetHistoricMeterDataNotFound) String() string {
	return fmt.Sprintf("[GET /inverter/{inverter_serial_number}/meter/data][%d] getHistoricMeterDataNotFound  %+v", 404, o.Payload)
}

func (o *GetHistoricMeterDataNotFound) GetPayload() *GetHistoricMeterDataNotFoundBody {
	return o.Payload
}

func (o *GetHistoricMeterDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetHistoricMeterDataNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetHistoricMeterDataBody get historic meter data body
swagger:model GetHistoricMeterDataBody
*/
type GetHistoricMeterDataBody struct {

	// Must be between 1 and 8.
	// Example: 2
	// Required: true
	Address *int64 `json:"address"`

	// Must be a date after or equal to <code>start_time</code>.
	// Example: 2106-02-22
	// Required: true
	EndTime *string `json:"end_time"`

	// Must be a date before or equal to <code>end_time</code>.
	// Example: 2019-05-22
	// Required: true
	StartTime *string `json:"start_time"`
}

// Validate validates this get historic meter data body
func (o *GetHistoricMeterDataBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHistoricMeterDataBody) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"address", "body", o.Address); err != nil {
		return err
	}

	return nil
}

func (o *GetHistoricMeterDataBody) validateEndTime(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"end_time", "body", o.EndTime); err != nil {
		return err
	}

	return nil
}

func (o *GetHistoricMeterDataBody) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"start_time", "body", o.StartTime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get historic meter data body based on context it is used
func (o *GetHistoricMeterDataBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetHistoricMeterDataBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHistoricMeterDataBody) UnmarshalBinary(b []byte) error {
	var res GetHistoricMeterDataBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetHistoricMeterDataNotFoundBody get historic meter data not found body
// Example: {"message":"Meter not found"}
swagger:model GetHistoricMeterDataNotFoundBody
*/
type GetHistoricMeterDataNotFoundBody struct {

	// message
	// Example: Meter not found
	Message string `json:"message,omitempty"`
}

// Validate validates this get historic meter data not found body
func (o *GetHistoricMeterDataNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get historic meter data not found body based on context it is used
func (o *GetHistoricMeterDataNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetHistoricMeterDataNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHistoricMeterDataNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetHistoricMeterDataNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetHistoricMeterDataOKBody get historic meter data o k body
// Example: {"data":[{"active_energy_net":940.5,"active_energy_total":503.1,"active_power_1":463.8,"active_power_2":-698.7,"active_power_3":440,"active_power_total":595.9,"apparent_power_1":-760.4,"apparent_power_2":-767.7,"apparent_power_3":289.5,"apparent_power_total":-974.2,"current_1":6120.5,"current_2":2171.8,"current_3":7712.8,"current_ln":7799,"current_total":1357.8,"export_active_energy":98.3,"export_reactive_energy":561.2,"frequency":49.57,"import_active_energy":202.9,"import_reactive_energy":626.5,"power_factor_1":0.4955,"power_factor_2":0.5087,"power_factor_3":0.314,"power_factor_total":0.331,"reactive_power_1":-595.2,"reactive_power_2":630.3,"reactive_power_3":916.3,"reactive_power_total":673.8,"time":"1997-05-28T05:10:15.000000Z","voltage_1":186.6,"voltage_2":56,"voltage_3":167.1},{"active_energy_net":761.6,"active_energy_total":989,"active_power_1":-262.9,"active_power_2":-461.5,"active_power_3":-895.7,"active_power_total":-715.1,"apparent_power_1":-910.4,"apparent_power_2":-938,"apparent_power_3":242.4,"apparent_power_total":335.6,"current_1":5568.4,"current_2":5689.8,"current_3":9184.3,"current_ln":6541.3,"current_total":6704.7,"export_active_energy":782.5,"export_reactive_energy":376.8,"frequency":49.66,"import_active_energy":776.2,"import_reactive_energy":813,"power_factor_1":0.6648,"power_factor_2":0.7496,"power_factor_3":0.2609,"power_factor_total":0.8952,"reactive_power_1":599.9,"reactive_power_2":104.9,"reactive_power_3":699.9,"reactive_power_total":464.4,"time":"1975-03-03T16:33:06.000000Z","voltage_1":197.3,"voltage_2":5.5,"voltage_3":169.8}]}
swagger:model GetHistoricMeterDataOKBody
*/
type GetHistoricMeterDataOKBody struct {

	// data
	// Example: [{"active_energy_net":940.5,"active_energy_total":503.1,"active_power_1":463.8,"active_power_2":-698.7,"active_power_3":440,"active_power_total":595.9,"apparent_power_1":-760.4,"apparent_power_2":-767.7,"apparent_power_3":289.5,"apparent_power_total":-974.2,"current_1":6120.5,"current_2":2171.8,"current_3":7712.8,"current_ln":7799,"current_total":1357.8,"export_active_energy":98.3,"export_reactive_energy":561.2,"frequency":49.57,"import_active_energy":202.9,"import_reactive_energy":626.5,"power_factor_1":0.4955,"power_factor_2":0.5087,"power_factor_3":0.314,"power_factor_total":0.331,"reactive_power_1":-595.2,"reactive_power_2":630.3,"reactive_power_3":916.3,"reactive_power_total":673.8,"time":"1997-05-28T05:10:15.000000Z","voltage_1":186.6,"voltage_2":56,"voltage_3":167.1},{"active_energy_net":761.6,"active_energy_total":989,"active_power_1":-262.9,"active_power_2":-461.5,"active_power_3":-895.7,"active_power_total":-715.1,"apparent_power_1":-910.4,"apparent_power_2":-938,"apparent_power_3":242.4,"apparent_power_total":335.6,"current_1":5568.4,"current_2":5689.8,"current_3":9184.3,"current_ln":6541.3,"current_total":6704.7,"export_active_energy":782.5,"export_reactive_energy":376.8,"frequency":49.66,"import_active_energy":776.2,"import_reactive_energy":813,"power_factor_1":0.6648,"power_factor_2":0.7496,"power_factor_3":0.2609,"power_factor_total":0.8952,"reactive_power_1":599.9,"reactive_power_2":104.9,"reactive_power_3":699.9,"reactive_power_total":464.4,"time":"1975-03-03T16:33:06.000000Z","voltage_1":197.3,"voltage_2":5.5,"voltage_3":169.8}]
	Data []*GetHistoricMeterDataOKBodyDataItems0 `json:"data"`
}

// Validate validates this get historic meter data o k body
func (o *GetHistoricMeterDataOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHistoricMeterDataOKBody) validateData(formats strfmt.Registry) error {
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
					return ve.ValidateName("getHistoricMeterDataOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getHistoricMeterDataOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get historic meter data o k body based on the context it is used
func (o *GetHistoricMeterDataOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHistoricMeterDataOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getHistoricMeterDataOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getHistoricMeterDataOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHistoricMeterDataOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHistoricMeterDataOKBody) UnmarshalBinary(b []byte) error {
	var res GetHistoricMeterDataOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetHistoricMeterDataOKBodyDataItems0 get historic meter data o k body data items0
swagger:model GetHistoricMeterDataOKBodyDataItems0
*/
type GetHistoricMeterDataOKBodyDataItems0 struct {

	// active energy net
	// Example: 940.5
	ActiveEnergyNet float64 `json:"active_energy_net,omitempty"`

	// active energy total
	// Example: 503.1
	ActiveEnergyTotal float64 `json:"active_energy_total,omitempty"`

	// active power 1
	// Example: 463.8
	ActivePower1 float64 `json:"active_power_1,omitempty"`

	// active power 2
	// Example: -698.7
	ActivePower2 float64 `json:"active_power_2,omitempty"`

	// active power 3
	// Example: 440
	ActivePower3 int64 `json:"active_power_3,omitempty"`

	// active power total
	// Example: 595.9
	ActivePowerTotal float64 `json:"active_power_total,omitempty"`

	// apparent power 1
	// Example: -760.4
	ApparentPower1 float64 `json:"apparent_power_1,omitempty"`

	// apparent power 2
	// Example: -767.7
	ApparentPower2 float64 `json:"apparent_power_2,omitempty"`

	// apparent power 3
	// Example: 289.5
	ApparentPower3 float64 `json:"apparent_power_3,omitempty"`

	// apparent power total
	// Example: -974.2
	ApparentPowerTotal float64 `json:"apparent_power_total,omitempty"`

	// current 1
	// Example: 6120.5
	Current1 float64 `json:"current_1,omitempty"`

	// current 2
	// Example: 2171.8
	Current2 float64 `json:"current_2,omitempty"`

	// current 3
	// Example: 7712.8
	Current3 float64 `json:"current_3,omitempty"`

	// current ln
	// Example: 7799
	CurrentLn int64 `json:"current_ln,omitempty"`

	// current total
	// Example: 1357.8
	CurrentTotal float64 `json:"current_total,omitempty"`

	// export active energy
	// Example: 98.3
	ExportActiveEnergy float64 `json:"export_active_energy,omitempty"`

	// export reactive energy
	// Example: 561.2
	ExportReactiveEnergy float64 `json:"export_reactive_energy,omitempty"`

	// frequency
	// Example: 49.57
	Frequency float64 `json:"frequency,omitempty"`

	// import active energy
	// Example: 202.9
	ImportActiveEnergy float64 `json:"import_active_energy,omitempty"`

	// import reactive energy
	// Example: 626.5
	ImportReactiveEnergy float64 `json:"import_reactive_energy,omitempty"`

	// power factor 1
	// Example: 0.4955
	PowerFactor1 float64 `json:"power_factor_1,omitempty"`

	// power factor 2
	// Example: 0.5087
	PowerFactor2 float64 `json:"power_factor_2,omitempty"`

	// power factor 3
	// Example: 0.314
	PowerFactor3 float64 `json:"power_factor_3,omitempty"`

	// power factor total
	// Example: 0.331
	PowerFactorTotal float64 `json:"power_factor_total,omitempty"`

	// reactive power 1
	// Example: -595.2
	ReactivePower1 float64 `json:"reactive_power_1,omitempty"`

	// reactive power 2
	// Example: 630.3
	ReactivePower2 float64 `json:"reactive_power_2,omitempty"`

	// reactive power 3
	// Example: 916.3
	ReactivePower3 float64 `json:"reactive_power_3,omitempty"`

	// reactive power total
	// Example: 673.8
	ReactivePowerTotal float64 `json:"reactive_power_total,omitempty"`

	// time
	// Example: 1997-05-28T05:10:15.000000Z
	Time string `json:"time,omitempty"`

	// voltage 1
	// Example: 186.6
	Voltage1 float64 `json:"voltage_1,omitempty"`

	// voltage 2
	// Example: 56
	Voltage2 int64 `json:"voltage_2,omitempty"`

	// voltage 3
	// Example: 167.1
	Voltage3 float64 `json:"voltage_3,omitempty"`
}

// Validate validates this get historic meter data o k body data items0
func (o *GetHistoricMeterDataOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get historic meter data o k body data items0 based on context it is used
func (o *GetHistoricMeterDataOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetHistoricMeterDataOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHistoricMeterDataOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res GetHistoricMeterDataOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
