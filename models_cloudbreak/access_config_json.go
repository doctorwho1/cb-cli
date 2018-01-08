// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AccessConfigJSON access config Json
// swagger:model AccessConfigJson

type AccessConfigJSON struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// properties
	Properties map[string]interface{} `json:"properties,omitempty"`
}

/* polymorph AccessConfigJson id false */

/* polymorph AccessConfigJson name false */

/* polymorph AccessConfigJson properties false */

// Validate validates this access config Json
func (m *AccessConfigJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AccessConfigJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessConfigJSON) UnmarshalBinary(b []byte) error {
	var res AccessConfigJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
