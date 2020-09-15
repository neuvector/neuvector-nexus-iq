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

// RESTAuthData REST auth data
//
// swagger:model RESTAuthData
type RESTAuthData struct {

	// token
	Token *RESTAuthToken `json:"Token,omitempty"`

	// client ip
	// Required: true
	ClientIP *string `json:"client_ip"`

	// password
	Password *RESTAuthPassword `json:"password,omitempty"`
}

// Validate validates this REST auth data
func (m *RESTAuthData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAuthData) validateToken(formats strfmt.Registry) error {

	if swag.IsZero(m.Token) { // not required
		return nil
	}

	if m.Token != nil {
		if err := m.Token.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Token")
			}
			return err
		}
	}

	return nil
}

func (m *RESTAuthData) validateClientIP(formats strfmt.Registry) error {

	if err := validate.Required("client_ip", "body", m.ClientIP); err != nil {
		return err
	}

	return nil
}

func (m *RESTAuthData) validatePassword(formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if m.Password != nil {
		if err := m.Password.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("password")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTAuthData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAuthData) UnmarshalBinary(b []byte) error {
	var res RESTAuthData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
