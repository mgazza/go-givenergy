// Code generated by go-swagger; DO NOT EDIT.

package inverter_data

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

// GetLatestEnergyDataReader is a Reader for the GetLatestEnergyData structure.
type GetLatestEnergyDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLatestEnergyDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLatestEnergyDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inverter/{inverter_serial_number}/meter-data/latest] getLatestEnergyData", response, response.Code())
	}
}

// NewGetLatestEnergyDataOK creates a GetLatestEnergyDataOK with default headers values
func NewGetLatestEnergyDataOK() *GetLatestEnergyDataOK {
	return &GetLatestEnergyDataOK{}
}

/*
GetLatestEnergyDataOK describes a response with status code 200, with default header values.

GetLatestEnergyDataOK get latest energy data o k
*/
type GetLatestEnergyDataOK struct {
	Payload *GetLatestEnergyDataOKBody
}

// IsSuccess returns true when this get latest energy data o k response has a 2xx status code
func (o *GetLatestEnergyDataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get latest energy data o k response has a 3xx status code
func (o *GetLatestEnergyDataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest energy data o k response has a 4xx status code
func (o *GetLatestEnergyDataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get latest energy data o k response has a 5xx status code
func (o *GetLatestEnergyDataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest energy data o k response a status code equal to that given
func (o *GetLatestEnergyDataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get latest energy data o k response
func (o *GetLatestEnergyDataOK) Code() int {
	return 200
}

func (o *GetLatestEnergyDataOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /inverter/{inverter_serial_number}/meter-data/latest][%d] getLatestEnergyDataOK %s", 200, payload)
}

func (o *GetLatestEnergyDataOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /inverter/{inverter_serial_number}/meter-data/latest][%d] getLatestEnergyDataOK %s", 200, payload)
}

func (o *GetLatestEnergyDataOK) GetPayload() *GetLatestEnergyDataOKBody {
	return o.Payload
}

func (o *GetLatestEnergyDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLatestEnergyDataOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetLatestEnergyDataOKBody get latest energy data o k body
// Example: {"data":{"is_metered":true,"time":"1992-09-29T14:37:47Z","today":{"ac_charge":30996.9,"battery":{"charge":91852.3,"discharge":65419.6},"consumption":938.6,"grid":{"export":4294.6,"import":9880.9},"solar":5690.3},"total":{"ac_charge":26925.1,"battery":{"charge":8763.7,"discharge":8763.7},"consumption":66779,"grid":{"export":5659,"import":184.9},"solar":57462.4}}}
swagger:model GetLatestEnergyDataOKBody
*/
type GetLatestEnergyDataOKBody struct {

	// data
	Data *GetLatestEnergyDataOKBodyData `json:"data,omitempty"`
}

// Validate validates this get latest energy data o k body
func (o *GetLatestEnergyDataOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestEnergyDataOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEnergyDataOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEnergyDataOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get latest energy data o k body based on the context it is used
func (o *GetLatestEnergyDataOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestEnergyDataOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEnergyDataOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEnergyDataOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBody) UnmarshalBinary(b []byte) error {
	var res GetLatestEnergyDataOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestEnergyDataOKBodyData get latest energy data o k body data
swagger:model GetLatestEnergyDataOKBodyData
*/
type GetLatestEnergyDataOKBodyData struct {

	// is metered
	// Example: true
	IsMetered bool `json:"is_metered,omitempty"`

	// time
	// Example: 1992-09-29T14:37:47Z
	Time string `json:"time,omitempty"`

	// today
	Today *GetLatestEnergyDataOKBodyDataToday `json:"today,omitempty"`

	// total
	Total *GetLatestEnergyDataOKBodyDataTotal `json:"total,omitempty"`
}

// Validate validates this get latest energy data o k body data
func (o *GetLatestEnergyDataOKBodyData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateToday(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestEnergyDataOKBodyData) validateToday(formats strfmt.Registry) error {
	if swag.IsZero(o.Today) { // not required
		return nil
	}

	if o.Today != nil {
		if err := o.Today.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "today")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "today")
			}
			return err
		}
	}

	return nil
}

func (o *GetLatestEnergyDataOKBodyData) validateTotal(formats strfmt.Registry) error {
	if swag.IsZero(o.Total) { // not required
		return nil
	}

	if o.Total != nil {
		if err := o.Total.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "total")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "total")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get latest energy data o k body data based on the context it is used
func (o *GetLatestEnergyDataOKBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateToday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestEnergyDataOKBodyData) contextValidateToday(ctx context.Context, formats strfmt.Registry) error {

	if o.Today != nil {

		if swag.IsZero(o.Today) { // not required
			return nil
		}

		if err := o.Today.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "today")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "today")
			}
			return err
		}
	}

	return nil
}

func (o *GetLatestEnergyDataOKBodyData) contextValidateTotal(ctx context.Context, formats strfmt.Registry) error {

	if o.Total != nil {

		if swag.IsZero(o.Total) { // not required
			return nil
		}

		if err := o.Total.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "total")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "total")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBodyData) UnmarshalBinary(b []byte) error {
	var res GetLatestEnergyDataOKBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestEnergyDataOKBodyDataToday get latest energy data o k body data today
swagger:model GetLatestEnergyDataOKBodyDataToday
*/
type GetLatestEnergyDataOKBodyDataToday struct {

	// ac charge
	// Example: 30996.9
	AcCharge float64 `json:"ac_charge,omitempty"`

	// battery
	Battery *GetLatestEnergyDataOKBodyDataTodayBattery `json:"battery,omitempty"`

	// consumption
	// Example: 938.6
	Consumption float64 `json:"consumption,omitempty"`

	// grid
	Grid *GetLatestEnergyDataOKBodyDataTodayGrid `json:"grid,omitempty"`

	// solar
	// Example: 5690.3
	Solar float64 `json:"solar,omitempty"`
}

// Validate validates this get latest energy data o k body data today
func (o *GetLatestEnergyDataOKBodyDataToday) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBattery(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGrid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestEnergyDataOKBodyDataToday) validateBattery(formats strfmt.Registry) error {
	if swag.IsZero(o.Battery) { // not required
		return nil
	}

	if o.Battery != nil {
		if err := o.Battery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "today" + "." + "battery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "today" + "." + "battery")
			}
			return err
		}
	}

	return nil
}

func (o *GetLatestEnergyDataOKBodyDataToday) validateGrid(formats strfmt.Registry) error {
	if swag.IsZero(o.Grid) { // not required
		return nil
	}

	if o.Grid != nil {
		if err := o.Grid.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "today" + "." + "grid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "today" + "." + "grid")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get latest energy data o k body data today based on the context it is used
func (o *GetLatestEnergyDataOKBodyDataToday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBattery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateGrid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestEnergyDataOKBodyDataToday) contextValidateBattery(ctx context.Context, formats strfmt.Registry) error {

	if o.Battery != nil {

		if swag.IsZero(o.Battery) { // not required
			return nil
		}

		if err := o.Battery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "today" + "." + "battery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "today" + "." + "battery")
			}
			return err
		}
	}

	return nil
}

func (o *GetLatestEnergyDataOKBodyDataToday) contextValidateGrid(ctx context.Context, formats strfmt.Registry) error {

	if o.Grid != nil {

		if swag.IsZero(o.Grid) { // not required
			return nil
		}

		if err := o.Grid.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "today" + "." + "grid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "today" + "." + "grid")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBodyDataToday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBodyDataToday) UnmarshalBinary(b []byte) error {
	var res GetLatestEnergyDataOKBodyDataToday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestEnergyDataOKBodyDataTodayBattery get latest energy data o k body data today battery
swagger:model GetLatestEnergyDataOKBodyDataTodayBattery
*/
type GetLatestEnergyDataOKBodyDataTodayBattery struct {

	// charge
	// Example: 91852.3
	Charge float64 `json:"charge,omitempty"`

	// discharge
	// Example: 65419.6
	Discharge float64 `json:"discharge,omitempty"`
}

// Validate validates this get latest energy data o k body data today battery
func (o *GetLatestEnergyDataOKBodyDataTodayBattery) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get latest energy data o k body data today battery based on context it is used
func (o *GetLatestEnergyDataOKBodyDataTodayBattery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBodyDataTodayBattery) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBodyDataTodayBattery) UnmarshalBinary(b []byte) error {
	var res GetLatestEnergyDataOKBodyDataTodayBattery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestEnergyDataOKBodyDataTodayGrid get latest energy data o k body data today grid
swagger:model GetLatestEnergyDataOKBodyDataTodayGrid
*/
type GetLatestEnergyDataOKBodyDataTodayGrid struct {

	// export
	// Example: 4294.6
	Export float64 `json:"export,omitempty"`

	// import
	// Example: 9880.9
	Import float64 `json:"import,omitempty"`
}

// Validate validates this get latest energy data o k body data today grid
func (o *GetLatestEnergyDataOKBodyDataTodayGrid) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get latest energy data o k body data today grid based on context it is used
func (o *GetLatestEnergyDataOKBodyDataTodayGrid) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBodyDataTodayGrid) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBodyDataTodayGrid) UnmarshalBinary(b []byte) error {
	var res GetLatestEnergyDataOKBodyDataTodayGrid
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestEnergyDataOKBodyDataTotal get latest energy data o k body data total
swagger:model GetLatestEnergyDataOKBodyDataTotal
*/
type GetLatestEnergyDataOKBodyDataTotal struct {

	// ac charge
	// Example: 26925.1
	AcCharge float64 `json:"ac_charge,omitempty"`

	// battery
	Battery *GetLatestEnergyDataOKBodyDataTotalBattery `json:"battery,omitempty"`

	// consumption
	// Example: 66779
	Consumption int64 `json:"consumption,omitempty"`

	// grid
	Grid *GetLatestEnergyDataOKBodyDataTotalGrid `json:"grid,omitempty"`

	// solar
	// Example: 57462.4
	Solar float64 `json:"solar,omitempty"`
}

// Validate validates this get latest energy data o k body data total
func (o *GetLatestEnergyDataOKBodyDataTotal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBattery(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGrid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestEnergyDataOKBodyDataTotal) validateBattery(formats strfmt.Registry) error {
	if swag.IsZero(o.Battery) { // not required
		return nil
	}

	if o.Battery != nil {
		if err := o.Battery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "total" + "." + "battery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "total" + "." + "battery")
			}
			return err
		}
	}

	return nil
}

func (o *GetLatestEnergyDataOKBodyDataTotal) validateGrid(formats strfmt.Registry) error {
	if swag.IsZero(o.Grid) { // not required
		return nil
	}

	if o.Grid != nil {
		if err := o.Grid.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "total" + "." + "grid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "total" + "." + "grid")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get latest energy data o k body data total based on the context it is used
func (o *GetLatestEnergyDataOKBodyDataTotal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBattery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateGrid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLatestEnergyDataOKBodyDataTotal) contextValidateBattery(ctx context.Context, formats strfmt.Registry) error {

	if o.Battery != nil {

		if swag.IsZero(o.Battery) { // not required
			return nil
		}

		if err := o.Battery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "total" + "." + "battery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "total" + "." + "battery")
			}
			return err
		}
	}

	return nil
}

func (o *GetLatestEnergyDataOKBodyDataTotal) contextValidateGrid(ctx context.Context, formats strfmt.Registry) error {

	if o.Grid != nil {

		if swag.IsZero(o.Grid) { // not required
			return nil
		}

		if err := o.Grid.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "total" + "." + "grid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getLatestEnergyDataOK" + "." + "data" + "." + "total" + "." + "grid")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBodyDataTotal) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBodyDataTotal) UnmarshalBinary(b []byte) error {
	var res GetLatestEnergyDataOKBodyDataTotal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestEnergyDataOKBodyDataTotalBattery get latest energy data o k body data total battery
swagger:model GetLatestEnergyDataOKBodyDataTotalBattery
*/
type GetLatestEnergyDataOKBodyDataTotalBattery struct {

	// charge
	// Example: 8763.7
	Charge float64 `json:"charge,omitempty"`

	// discharge
	// Example: 8763.7
	Discharge float64 `json:"discharge,omitempty"`
}

// Validate validates this get latest energy data o k body data total battery
func (o *GetLatestEnergyDataOKBodyDataTotalBattery) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get latest energy data o k body data total battery based on context it is used
func (o *GetLatestEnergyDataOKBodyDataTotalBattery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBodyDataTotalBattery) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBodyDataTotalBattery) UnmarshalBinary(b []byte) error {
	var res GetLatestEnergyDataOKBodyDataTotalBattery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLatestEnergyDataOKBodyDataTotalGrid get latest energy data o k body data total grid
swagger:model GetLatestEnergyDataOKBodyDataTotalGrid
*/
type GetLatestEnergyDataOKBodyDataTotalGrid struct {

	// export
	// Example: 5659
	Export int64 `json:"export,omitempty"`

	// import
	// Example: 184.9
	Import float64 `json:"import,omitempty"`
}

// Validate validates this get latest energy data o k body data total grid
func (o *GetLatestEnergyDataOKBodyDataTotalGrid) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get latest energy data o k body data total grid based on context it is used
func (o *GetLatestEnergyDataOKBodyDataTotalGrid) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBodyDataTotalGrid) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLatestEnergyDataOKBodyDataTotalGrid) UnmarshalBinary(b []byte) error {
	var res GetLatestEnergyDataOKBodyDataTotalGrid
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
