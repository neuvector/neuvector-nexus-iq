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

// RESTResponseRule REST response rule
//
// swagger:model RESTResponseRule
type RESTResponseRule struct {

	// actions
	// Required: true
	Actions []string `json:"actions"`

	// cfg type
	// Required: true
	CfgType *string `json:"cfg_type"`

	// comment
	Comment string `json:"comment,omitempty"`

	// conditions
	Conditions []*RESTCLUSEventCondition `json:"conditions"`

	// disable
	Disable bool `json:"disable,omitempty"`

	// event
	// Required: true
	Event *string `json:"event"`

	// group
	Group string `json:"group,omitempty"`

	// id
	// Required: true
	ID *uint32 `json:"id"`
}

// Validate validates this REST response rule
func (m *RESTResponseRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCfgType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RESTResponseRule) validateActions(formats strfmt.Registry) error {

	if err := validate.Required("actions", "body", m.Actions); err != nil {
		return err
	}

	return nil
}

func (m *RESTResponseRule) validateCfgType(formats strfmt.Registry) error {

	if err := validate.Required("cfg_type", "body", m.CfgType); err != nil {
		return err
	}

	return nil
}

func (m *RESTResponseRule) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RESTResponseRule) validateEvent(formats strfmt.Registry) error {

	if err := validate.Required("event", "body", m.Event); err != nil {
		return err
	}

	return nil
}

func (m *RESTResponseRule) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RESTResponseRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RESTResponseRule) UnmarshalBinary(b []byte) error {
	var res RESTResponseRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
