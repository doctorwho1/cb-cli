// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StackRepoDetailsJSON stack repo details Json
// swagger:model StackRepoDetailsJson

type StackRepoDetailsJSON struct {

	// stack
	Stack map[string]string `json:"stack,omitempty"`

	// util
	Util map[string]string `json:"util,omitempty"`
}

/* polymorph StackRepoDetailsJson stack false */

/* polymorph StackRepoDetailsJson util false */

// Validate validates this stack repo details Json
func (m *StackRepoDetailsJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StackRepoDetailsJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackRepoDetailsJSON) UnmarshalBinary(b []byte) error {
	var res StackRepoDetailsJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
