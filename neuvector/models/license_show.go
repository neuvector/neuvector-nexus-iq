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

// RESTLicenseShow REST license show
//
// swagger:model RESTLicenseShow
type RESTLicenseShow struct {

	// day to expire
	// Required: true
	DayToExpire *int64 `json:"day_to_expire"`

	// info
	// Required: true
	Info *RESTLicenseInfo `json:"info"`
}

// Validate validates this REST license show
func (m *RESTLicenseShow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDayToExpire(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTLicenseShow) validateDayToExpire(formats strfmt.Registry) error {

	if err := validate.Required("day_to_expire", "body", m.DayToExpire); err != nil {
		return err
	}

	return nil
}

func (m *RESTLicenseShow) validateInfo(formats strfmt.Registry) error {

	if err := validate.Required("info", "body", m.Info); err != nil {
		return err
	}

	if m.Info != nil {
		if err := m.Info.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTLicenseShow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTLicenseShow) UnmarshalBinary(b []byte) error {
	var res RESTLicenseShow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
