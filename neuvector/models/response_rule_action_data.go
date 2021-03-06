// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RESTResponseRuleActionData REST response rule action data
//
// swagger:model RESTResponseRuleActionData
type RESTResponseRuleActionData struct {

	// insert
	Insert *RESTResponseRuleInsert `json:"insert,omitempty"`
}

// Validate validates this REST response rule action data
func (m *RESTResponseRuleActionData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInsert(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTResponseRuleActionData) validateInsert(formats strfmt.Registry) error {

	if swag.IsZero(m.Insert) { // not required
		return nil
	}

	if m.Insert != nil {
		if err := m.Insert.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("insert")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTResponseRuleActionData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTResponseRuleActionData) UnmarshalBinary(b []byte) error {
	var res RESTResponseRuleActionData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
