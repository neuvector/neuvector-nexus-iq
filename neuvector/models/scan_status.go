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

// RESTScanStatus REST scan status
//
// swagger:model RESTScanStatus
type RESTScanStatus struct {

	// cvedb create time
	// Required: true
	CvedbCreateTime *string `json:"cvedb_create_time"`

	// cvedb version
	// Required: true
	CvedbVersion *string `json:"cvedb_version"`

	// failed
	// Required: true
	Failed *int64 `json:"failed"`

	// scanned
	// Required: true
	Scanned *int64 `json:"scanned"`

	// scanning
	// Required: true
	Scanning *int64 `json:"scanning"`

	// scheduled
	// Required: true
	Scheduled *int64 `json:"scheduled"`
}

// Validate validates this REST scan status
func (m *RESTScanStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCvedbCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCvedbVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanned(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanning(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTScanStatus) validateCvedbCreateTime(formats strfmt.Registry) error {

	if err := validate.Required("cvedb_create_time", "body", m.CvedbCreateTime); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanStatus) validateCvedbVersion(formats strfmt.Registry) error {

	if err := validate.Required("cvedb_version", "body", m.CvedbVersion); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanStatus) validateFailed(formats strfmt.Registry) error {

	if err := validate.Required("failed", "body", m.Failed); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanStatus) validateScanned(formats strfmt.Registry) error {

	if err := validate.Required("scanned", "body", m.Scanned); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanStatus) validateScanning(formats strfmt.Registry) error {

	if err := validate.Required("scanning", "body", m.Scanning); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanStatus) validateScheduled(formats strfmt.Registry) error {

	if err := validate.Required("scheduled", "body", m.Scheduled); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTScanStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTScanStatus) UnmarshalBinary(b []byte) error {
	var res RESTScanStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
