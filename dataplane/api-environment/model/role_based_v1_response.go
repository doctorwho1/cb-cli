// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RoleBasedV1Response role based v1 response
// swagger:model RoleBasedV1Response
type RoleBasedV1Response struct {

	// app object Id
	AppObjectID string `json:"appObjectId,omitempty"`

	// code grant flow
	CodeGrantFlow bool `json:"codeGrantFlow,omitempty"`

	// sp display name
	SpDisplayName string `json:"spDisplayName,omitempty"`
}

// Validate validates this role based v1 response
func (m *RoleBasedV1Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RoleBasedV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleBasedV1Response) UnmarshalBinary(b []byte) error {
	var res RoleBasedV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
