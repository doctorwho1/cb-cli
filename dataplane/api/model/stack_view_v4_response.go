// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StackViewV4Response stack view v4 response
// swagger:model StackViewV4Response
type StackViewV4Response struct {

	// cluster object on stack
	Cluster *ClusterViewV4Response `json:"cluster,omitempty"`

	// creation time of the stack in long
	Created int64 `json:"created,omitempty"`

	// environment which the stack is assigned to
	Environment *EnvironmentSettingsV4Response `json:"environment,omitempty"`

	// specific version of HDP
	HdpVersion string `json:"hdpVersion,omitempty"`

	// id of the stack
	ID int64 `json:"id,omitempty"`

	// name of the stack
	// Required: true
	Name *string `json:"name"`

	// node count of the stack
	NodeCount int32 `json:"nodeCount,omitempty"`

	// status of the stack
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED CREATE_FAILED ENABLE_SECURITY_FAILED PRE_DELETE_IN_PROGRESS DELETE_IN_PROGRESS DELETE_FAILED DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED]
	Status string `json:"status,omitempty"`

	// termination completion time of stack in long
	Terminated int64 `json:"terminated,omitempty"`

	// the related user
	User *UserViewV4Response `json:"user,omitempty"`
}

// Validate validates this stack view v4 response
func (m *StackViewV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackViewV4Response) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *StackViewV4Response) validateEnvironment(formats strfmt.Registry) error {

	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	if m.Environment != nil {
		if err := m.Environment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

func (m *StackViewV4Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var stackViewV4ResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackViewV4ResponseTypeStatusPropEnum = append(stackViewV4ResponseTypeStatusPropEnum, v)
	}
}

const (

	// StackViewV4ResponseStatusREQUESTED captures enum value "REQUESTED"
	StackViewV4ResponseStatusREQUESTED string = "REQUESTED"

	// StackViewV4ResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	StackViewV4ResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// StackViewV4ResponseStatusAVAILABLE captures enum value "AVAILABLE"
	StackViewV4ResponseStatusAVAILABLE string = "AVAILABLE"

	// StackViewV4ResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	StackViewV4ResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// StackViewV4ResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	StackViewV4ResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// StackViewV4ResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	StackViewV4ResponseStatusUPDATEFAILED string = "UPDATE_FAILED"

	// StackViewV4ResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	StackViewV4ResponseStatusCREATEFAILED string = "CREATE_FAILED"

	// StackViewV4ResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	StackViewV4ResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"

	// StackViewV4ResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	StackViewV4ResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"

	// StackViewV4ResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	StackViewV4ResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// StackViewV4ResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	StackViewV4ResponseStatusDELETEFAILED string = "DELETE_FAILED"

	// StackViewV4ResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	StackViewV4ResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// StackViewV4ResponseStatusSTOPPED captures enum value "STOPPED"
	StackViewV4ResponseStatusSTOPPED string = "STOPPED"

	// StackViewV4ResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	StackViewV4ResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// StackViewV4ResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	StackViewV4ResponseStatusSTARTREQUESTED string = "START_REQUESTED"

	// StackViewV4ResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	StackViewV4ResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// StackViewV4ResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	StackViewV4ResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// StackViewV4ResponseStatusSTARTFAILED captures enum value "START_FAILED"
	StackViewV4ResponseStatusSTARTFAILED string = "START_FAILED"

	// StackViewV4ResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	StackViewV4ResponseStatusSTOPFAILED string = "STOP_FAILED"

	// StackViewV4ResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	StackViewV4ResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// StackViewV4ResponseStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	StackViewV4ResponseStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"
)

// prop value enum
func (m *StackViewV4Response) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackViewV4ResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackViewV4Response) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *StackViewV4Response) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackViewV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackViewV4Response) UnmarshalBinary(b []byte) error {
	var res StackViewV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
