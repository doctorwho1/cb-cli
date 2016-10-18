package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*HostGroupBase host group base

swagger:model HostGroupBase
*/
type HostGroupBase struct {

	/* instance group or resource constraint for a hostgroup

	Required: true
	*/
	Constraint *Constraint `json:"constraint"`

	/* name of the resource

	Required: true
	*/
	Name string `json:"name"`

	/* referenced recipe ids

	Unique: true
	*/
	RecipeIds []int64 `json:"recipeIds,omitempty"`
}

// Validate validates this host group base
func (m *HostGroupBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstraint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecipeIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostGroupBase) validateConstraint(formats strfmt.Registry) error {

	if m.Constraint != nil {

		if err := m.Constraint.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *HostGroupBase) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *HostGroupBase) validateRecipeIds(formats strfmt.Registry) error {

	if swag.IsZero(m.RecipeIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("recipeIds", "body", m.RecipeIds); err != nil {
		return err
	}

	for i := 0; i < len(m.RecipeIds); i++ {

		if err := validate.Required("recipeIds"+"."+strconv.Itoa(i), "body", int64(m.RecipeIds[i])); err != nil {
			return err
		}

	}

	return nil
}
