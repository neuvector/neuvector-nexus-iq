// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RESTSystemRequest REST system request
//
// swagger:model RESTSystemRequest
type RESTSystemRequest struct {

	// policy mode
	PolicyMode string `json:"policy_mode,omitempty"`

	// unquarantine
	Unquarantine *RESTUnquarReq `json:"unquarantine,omitempty"`
}

// Validate validates this REST system request
func (m *RESTSystemRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnquarantine(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTSystemRequest) validateUnquarantine(formats strfmt.Registry) error {

	if swag.IsZero(m.Unquarantine) { // not required
		return nil
	}

	if m.Unquarantine != nil {
		if err := m.Unquarantine.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unquarantine")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTSystemRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTSystemRequest) UnmarshalBinary(b []byte) error {
	var res RESTSystemRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
