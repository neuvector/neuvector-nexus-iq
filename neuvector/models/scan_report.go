// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RESTScanReport REST scan report
//
// swagger:model RESTScanReport
type RESTScanReport struct {

	// envs
	// Required: true
	Envs []string `json:"envs"`

	// labels
	// Required: true
	Labels *RESTScanReportLabels `json:"labels"`

	// modules
	// Required: true
	Modules []*RESTScanModule `json:"modules"`

	// vulnerabilities
	// Required: true
	Vulnerabilities []*RESTVulnerability `json:"vulnerabilities"`
}

// Validate validates this REST scan report
func (m *RESTScanReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVulnerabilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTScanReport) validateEnvs(formats strfmt.Registry) error {

	if err := validate.Required("envs", "body", m.Envs); err != nil {
		return err
	}

	return nil
}

func (m *RESTScanReport) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("labels", "body", m.Labels); err != nil {
		return err
	}

	if m.Labels != nil {
		if err := m.Labels.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels")
			}
			return err
		}
	}

	return nil
}

func (m *RESTScanReport) validateModules(formats strfmt.Registry) error {

	if err := validate.Required("modules", "body", m.Modules); err != nil {
		return err
	}

	for i := 0; i < len(m.Modules); i++ {
		if swag.IsZero(m.Modules[i]) { // not required
			continue
		}

		if m.Modules[i] != nil {
			if err := m.Modules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RESTScanReport) validateVulnerabilities(formats strfmt.Registry) error {

	if err := validate.Required("vulnerabilities", "body", m.Vulnerabilities); err != nil {
		return err
	}

	for i := 0; i < len(m.Vulnerabilities); i++ {
		if swag.IsZero(m.Vulnerabilities[i]) { // not required
			continue
		}

		if m.Vulnerabilities[i] != nil {
			if err := m.Vulnerabilities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vulnerabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTScanReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTScanReport) UnmarshalBinary(b []byte) error {
	var res RESTScanReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RESTScanReportLabels REST scan report labels
//
// swagger:model RESTScanReportLabels
type RESTScanReportLabels struct {

	// key
	Key []string `json:"key"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this REST scan report labels
func (m *RESTScanReportLabels) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RESTScanReportLabels) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTScanReportLabels) UnmarshalBinary(b []byte) error {
	var res RESTScanReportLabels
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
