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

// RESTRegistryConfigData REST registry config data
//
// swagger:model RESTRegistryConfigData
type RESTRegistryConfigData struct {

	// config
	// Required: true
	Config *RESTRegistryConfig `json:"config"`
}

// Validate validates this REST registry config data
func (m *RESTRegistryConfigData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTRegistryConfigData) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("config", "body", m.Config); err != nil {
		return err
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTRegistryConfigData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTRegistryConfigData) UnmarshalBinary(b []byte) error {
	var res RESTRegistryConfigData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
