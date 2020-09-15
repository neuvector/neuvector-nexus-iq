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

// RESTAuthToken REST auth token
//
// swagger:model RESTAuthToken
type RESTAuthToken struct {

	// redirect endpoint
	// Required: true
	RedirectEndpoint *string `json:"redirect_endpoint"`

	// state
	// Required: true
	State *string `json:"state"`

	// token
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this REST auth token
func (m *RESTAuthToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRedirectEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTAuthToken) validateRedirectEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("redirect_endpoint", "body", m.RedirectEndpoint); err != nil {
		return err
	}

	return nil
}

func (m *RESTAuthToken) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *RESTAuthToken) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTAuthToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTAuthToken) UnmarshalBinary(b []byte) error {
	var res RESTAuthToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
