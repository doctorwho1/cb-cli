// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountPreference account preference
// swagger:model AccountPreference

type AccountPreference struct {

	// allowed instance types in the account (empty list for no restriction)
	AllowedInstanceTypes []string `json:"allowedInstanceTypes"`

	// lifecycle of the cluster in hours (0 for immortal clusters)
	// Required: true
	// Minimum: 0
	ClusterTimeToLive *int64 `json:"clusterTimeToLive"`

	// default tags for the resources created
	DefaultTags map[string]string `json:"defaultTags,omitempty"`

	// max number of clusters in the account (0 when unlimited)
	// Minimum: 0
	MaxNumberOfClusters *int64 `json:"maxNumberOfClusters,omitempty"`

	// max number of clusters for user within the account (0 when unlimited)
	// Minimum: 0
	MaxNumberOfClustersPerUser *int64 `json:"maxNumberOfClustersPerUser,omitempty"`

	// max number of vms in a cluster of account (0 when unlimited)
	// Minimum: 0
	MaxNumberOfNodesPerCluster *int64 `json:"maxNumberOfNodesPerCluster,omitempty"`

	// list of the cloudplatforms visible on the UI
	Platforms string `json:"platforms,omitempty"`

	// lifecycle of the account and its clusters in hours (0 for immortal account)
	// Required: true
	// Minimum: 0
	UserTimeToLive *int64 `json:"userTimeToLive"`
}

/* polymorph AccountPreference allowedInstanceTypes false */

/* polymorph AccountPreference clusterTimeToLive false */

/* polymorph AccountPreference defaultTags false */

/* polymorph AccountPreference maxNumberOfClusters false */

/* polymorph AccountPreference maxNumberOfClustersPerUser false */

/* polymorph AccountPreference maxNumberOfNodesPerCluster false */

/* polymorph AccountPreference platforms false */

/* polymorph AccountPreference userTimeToLive false */

// Validate validates this account preference
func (m *AccountPreference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedInstanceTypes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClusterTimeToLive(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaxNumberOfClusters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaxNumberOfClustersPerUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaxNumberOfNodesPerCluster(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserTimeToLive(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountPreference) validateAllowedInstanceTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.AllowedInstanceTypes) { // not required
		return nil
	}

	return nil
}

func (m *AccountPreference) validateClusterTimeToLive(formats strfmt.Registry) error {

	if err := validate.Required("clusterTimeToLive", "body", m.ClusterTimeToLive); err != nil {
		return err
	}

	if err := validate.MinimumInt("clusterTimeToLive", "body", int64(*m.ClusterTimeToLive), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *AccountPreference) validateMaxNumberOfClusters(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxNumberOfClusters) { // not required
		return nil
	}

	if err := validate.MinimumInt("maxNumberOfClusters", "body", int64(*m.MaxNumberOfClusters), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *AccountPreference) validateMaxNumberOfClustersPerUser(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxNumberOfClustersPerUser) { // not required
		return nil
	}

	if err := validate.MinimumInt("maxNumberOfClustersPerUser", "body", int64(*m.MaxNumberOfClustersPerUser), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *AccountPreference) validateMaxNumberOfNodesPerCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxNumberOfNodesPerCluster) { // not required
		return nil
	}

	if err := validate.MinimumInt("maxNumberOfNodesPerCluster", "body", int64(*m.MaxNumberOfNodesPerCluster), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *AccountPreference) validateUserTimeToLive(formats strfmt.Registry) error {

	if err := validate.Required("userTimeToLive", "body", m.UserTimeToLive); err != nil {
		return err
	}

	if err := validate.MinimumInt("userTimeToLive", "body", int64(*m.UserTimeToLive), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountPreference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPreference) UnmarshalBinary(b []byte) error {
	var res AccountPreference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
