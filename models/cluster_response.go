package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*ClusterResponse cluster response

swagger:model ClusterResponse
*/
type ClusterResponse struct {

	/* public ambari ip of the stack
	 */
	AmbariServerIP *string `json:"ambariServerIp,omitempty"`

	/* ambari stack details
	 */
	AmbariStackDetails *AmbariStackDetails `json:"ambariStackDetails,omitempty"`

	/* Additional information for ambari cluster
	 */
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	/* blueprint id for the cluster
	 */
	BlueprintID *int64 `json:"blueprintId,omitempty"`

	/* blueprint inputs in the cluster

	Unique: true
	*/
	BlueprintInputs []*BlueprintInput `json:"blueprintInputs,omitempty"`

	/* name of the cluster
	 */
	Cluster *string `json:"cluster,omitempty"`

	/* config recommendation strategy
	 */
	ConfigStrategy *string `json:"configStrategy,omitempty"`

	/* description of the resource
	 */
	Description *string `json:"description,omitempty"`

	/* shipyard service enabled in the cluster
	 */
	EnableShipyard *bool `json:"enableShipyard,omitempty"`

	/* host groups

	Unique: true
	*/
	HostGroups []*HostGroupResponse `json:"hostGroups,omitempty"`

	/* duration - how long the cluster is running in hours
	 */
	HoursUp *int32 `json:"hoursUp,omitempty"`

	/* id of the resource
	 */
	ID *int64 `json:"id,omitempty"`

	/* LDAP config id for the cluster
	 */
	LdapConfigID *int64 `json:"ldapConfigId,omitempty"`

	/* flag for default LDAP support
	 */
	LdapRequired *bool `json:"ldapRequired,omitempty"`

	/* duration - how long the cluster is running in minutes (minus hours)
	 */
	MinutesUp *int32 `json:"minutesUp,omitempty"`

	/* name of the resource
	 */
	Name *string `json:"name,omitempty"`

	/* ambari password
	 */
	Password *string `json:"password,omitempty"`

	/* RDS configuration id for the cluster
	 */
	RdsConfigID *int64 `json:"rdsConfigId,omitempty"`

	/* secure
	 */
	Secure *bool `json:"secure,omitempty"`

	/* most important services in the cluster
	 */
	ServiceEndPoints map[string]string `json:"serviceEndPoints,omitempty"`

	/* SSSD config id for the cluster
	 */
	SssdConfigID *int64 `json:"sssdConfigId,omitempty"`

	/* status of the cluster
	 */
	Status *string `json:"status,omitempty"`

	/* status message of the cluster
	 */
	StatusReason *string `json:"statusReason,omitempty"`

	/* ambari username
	 */
	UserName *string `json:"userName,omitempty"`
}

// Validate validates this cluster response
func (m *ClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlueprintInputs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConfigStrategy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHostGroups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServiceEndPoints(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterResponse) validateBlueprintInputs(formats strfmt.Registry) error {

	if swag.IsZero(m.BlueprintInputs) { // not required
		return nil
	}

	if err := validate.UniqueItems("blueprintInputs", "body", m.BlueprintInputs); err != nil {
		return err
	}

	for i := 0; i < len(m.BlueprintInputs); i++ {

		if m.BlueprintInputs[i] != nil {

			if err := m.BlueprintInputs[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

var clusterResponseTypeConfigStrategyPropEnum []interface{}

func (m *ClusterResponse) validateConfigStrategyEnum(path, location string, value string) error {
	if clusterResponseTypeConfigStrategyPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["NEVER_APPLY","ONLY_STACK_DEFAULTS_APPLY","ALWAYS_APPLY","ALWAYS_APPLY_DONT_OVERRIDE_CUSTOM_VALUES"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			clusterResponseTypeConfigStrategyPropEnum = append(clusterResponseTypeConfigStrategyPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, clusterResponseTypeConfigStrategyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterResponse) validateConfigStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigStrategy) { // not required
		return nil
	}

	if err := m.validateConfigStrategyEnum("configStrategy", "body", *m.ConfigStrategy); err != nil {
		return err
	}

	return nil
}

func (m *ClusterResponse) validateHostGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.HostGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("hostGroups", "body", m.HostGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.HostGroups); i++ {

		if m.HostGroups[i] != nil {

			if err := m.HostGroups[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ClusterResponse) validateServiceEndPoints(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceEndPoints) { // not required
		return nil
	}

	if err := validate.Required("serviceEndPoints", "body", m.ServiceEndPoints); err != nil {
		return err
	}

	return nil
}
