// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TagsV4Response tags v4 response
// swagger:model TagsV4Response
type TagsV4Response struct {

	// stack related application tags
	Application map[string]string `json:"application,omitempty"`

	// stack related default tags
	Defaults map[string]string `json:"defaults,omitempty"`

	// stack related userdefined tags
	UserDefined map[string]string `json:"userDefined,omitempty"`
}

// Validate validates this tags v4 response
func (m *TagsV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TagsV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagsV4Response) UnmarshalBinary(b []byte) error {
	var res TagsV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
