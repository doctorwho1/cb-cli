// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RootVolumeV4Response root volume v4 response
// swagger:model RootVolumeV4Response
type RootVolumeV4Response struct {

	// size of the root volume
	Size int32 `json:"size,omitempty"`
}

// Validate validates this root volume v4 response
func (m *RootVolumeV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RootVolumeV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RootVolumeV4Response) UnmarshalBinary(b []byte) error {
	var res RootVolumeV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}