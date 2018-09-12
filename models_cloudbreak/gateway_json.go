// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GatewayJSON gateway Json
// swagger:model GatewayJson

type GatewayJSON struct {

	// [DEPRECATED] 'enableGateway' is no longer needed to determine if gateway needs to be launched or not.
	// Presence of gateway definition in request is suffucicient. This value is only used in legacy requests,
	// when 'topologyName' or 'exposedServices' is defined in the root of Gateway, instead of using topologies.
	// When it is a legacy request and 'enableGateway' is set to 'false', gateway will not be saved and created.
	EnableGateway *bool `json:"enableGateway,omitempty"`

	// [DEPRECATED] Use the 'exposedServices' inside the 'topologies' part of the request.
	// If 'exposedServices' is specified, other deprecated properties ('topologyName' and 'enableGateway') will be used as well, and 'topologies' will be ignored.
	ExposedServices []string `json:"exposedServices"`

	// Knox gateway type
	GatewayType string `json:"gatewayType,omitempty"`

	// Knox SSO type
	SsoType string `json:"ssoType,omitempty"`

	// SSO Provider certificate
	TokenCert string `json:"tokenCert,omitempty"`

	// Topology definitions of the gateway.
	Topologies []*GatewayTopologyJSON `json:"topologies"`

	// [DEPRECATED] Use the 'topologyName' inside the 'topologies' part of the request.
	// If 'topologyName' is specified, other deprecated properties ('exposedServices' and 'enableGateway') will be used as well, and 'topologies' will be ignored.
	TopologyName string `json:"topologyName,omitempty"`
}

/* polymorph GatewayJson enableGateway false */

/* polymorph GatewayJson exposedServices false */

/* polymorph GatewayJson gatewayType false */

/* polymorph GatewayJson ssoType false */

/* polymorph GatewayJson tokenCert false */

/* polymorph GatewayJson topologies false */

/* polymorph GatewayJson topologyName false */

// Validate validates this gateway Json
func (m *GatewayJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExposedServices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGatewayType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSsoType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTopologies(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayJSON) validateExposedServices(formats strfmt.Registry) error {

	if swag.IsZero(m.ExposedServices) { // not required
		return nil
	}

	return nil
}

var gatewayJsonTypeGatewayTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CENTRAL","INDIVIDUAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gatewayJsonTypeGatewayTypePropEnum = append(gatewayJsonTypeGatewayTypePropEnum, v)
	}
}

const (
	// GatewayJSONGatewayTypeCENTRAL captures enum value "CENTRAL"
	GatewayJSONGatewayTypeCENTRAL string = "CENTRAL"
	// GatewayJSONGatewayTypeINDIVIDUAL captures enum value "INDIVIDUAL"
	GatewayJSONGatewayTypeINDIVIDUAL string = "INDIVIDUAL"
)

// prop value enum
func (m *GatewayJSON) validateGatewayTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, gatewayJsonTypeGatewayTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GatewayJSON) validateGatewayType(formats strfmt.Registry) error {

	if swag.IsZero(m.GatewayType) { // not required
		return nil
	}

	// value enum
	if err := m.validateGatewayTypeEnum("gatewayType", "body", m.GatewayType); err != nil {
		return err
	}

	return nil
}

var gatewayJsonTypeSsoTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SSO_PROVIDER","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gatewayJsonTypeSsoTypePropEnum = append(gatewayJsonTypeSsoTypePropEnum, v)
	}
}

const (
	// GatewayJSONSsoTypeSSOPROVIDER captures enum value "SSO_PROVIDER"
	GatewayJSONSsoTypeSSOPROVIDER string = "SSO_PROVIDER"
	// GatewayJSONSsoTypeNONE captures enum value "NONE"
	GatewayJSONSsoTypeNONE string = "NONE"
)

// prop value enum
func (m *GatewayJSON) validateSsoTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, gatewayJsonTypeSsoTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GatewayJSON) validateSsoType(formats strfmt.Registry) error {

	if swag.IsZero(m.SsoType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSsoTypeEnum("ssoType", "body", m.SsoType); err != nil {
		return err
	}

	return nil
}

func (m *GatewayJSON) validateTopologies(formats strfmt.Registry) error {

	if swag.IsZero(m.Topologies) { // not required
		return nil
	}

	for i := 0; i < len(m.Topologies); i++ {

		if swag.IsZero(m.Topologies[i]) { // not required
			continue
		}

		if m.Topologies[i] != nil {

			if err := m.Topologies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topologies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GatewayJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayJSON) UnmarshalBinary(b []byte) error {
	var res GatewayJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
