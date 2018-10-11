// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SimpleEnvironmentResponse simple environment response
// swagger:model SimpleEnvironmentResponse

type SimpleEnvironmentResponse struct {

	// Cloud platform of the environment.
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// description of the resource
	Description string `json:"description,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// name of the resource
	Name string `json:"name,omitempty"`

	// Regions of the environment.
	// Unique: true
	Regions []string `json:"regions"`

	// workspace
	Workspace *WorkspaceResourceResponse `json:"workspace,omitempty"`
}

/* polymorph SimpleEnvironmentResponse cloudPlatform false */

/* polymorph SimpleEnvironmentResponse description false */

/* polymorph SimpleEnvironmentResponse id false */

/* polymorph SimpleEnvironmentResponse name false */

/* polymorph SimpleEnvironmentResponse regions false */

/* polymorph SimpleEnvironmentResponse workspace false */

// Validate validates this simple environment response
func (m *SimpleEnvironmentResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimpleEnvironmentResponse) validateRegions(formats strfmt.Registry) error {

	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	if err := validate.UniqueItems("regions", "body", m.Regions); err != nil {
		return err
	}

	return nil
}

func (m *SimpleEnvironmentResponse) validateWorkspace(formats strfmt.Registry) error {

	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {

		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SimpleEnvironmentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimpleEnvironmentResponse) UnmarshalBinary(b []byte) error {
	var res SimpleEnvironmentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
