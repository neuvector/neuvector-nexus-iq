// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RESTScanReportData REST scan report data
//
// swagger:model RESTScanReportData
type RESTScanReportData struct {

	// report
	// Required: true
	Report *RESTScanReport `json:"report"`
}

// Validate validates this REST scan report data
func (m *RESTScanReportData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReport(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTScanReportData) validateReport(formats strfmt.Registry) error {

	if err := validate.Required("report", "body", m.Report); err != nil {
		return err
	}

	if m.Report != nil {
		if err := m.Report.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("report")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTScanReportData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTScanReportData) UnmarshalBinary(b []byte) error {
	var res RESTScanReportData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
