// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AccountPreferencesRequest account preferences request
// swagger:model AccountPreferencesRequest

type AccountPreferencesRequest struct {

	// default tags for the resources created
	DefaultTags map[string]string `json:"defaultTags,omitempty"`

	// list of the cloudplatforms visible on the UI
	Platforms string `json:"platforms,omitempty"`

	// smartsense enabled on the UI
	SmartsenseEnabled *bool `json:"smartsenseEnabled,omitempty"`
}

/* polymorph AccountPreferencesRequest defaultTags false */

/* polymorph AccountPreferencesRequest platforms false */

/* polymorph AccountPreferencesRequest smartsenseEnabled false */

// Validate validates this account preferences request
func (m *AccountPreferencesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AccountPreferencesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPreferencesRequest) UnmarshalBinary(b []byte) error {
	var res AccountPreferencesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}