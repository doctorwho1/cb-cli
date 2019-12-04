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

// StorageLocationBase storage location base
// swagger:model StorageLocationBase
type StorageLocationBase struct {

	// type
	// Enum: [ZEPPELIN_NOTEBOOK YARN_LOG HIVE_METASTORE_WAREHOUSE HIVE_METASTORE_EXTERNAL_WAREHOUSE HIVE_REPLICA_WAREHOUSE HBASE_ROOT RANGER_AUDIT PROFILER_SERVICE_FS_URI DEFAULT_FS]
	Type string `json:"type,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this storage location base
func (m *StorageLocationBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var storageLocationBaseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ZEPPELIN_NOTEBOOK","YARN_LOG","HIVE_METASTORE_WAREHOUSE","HIVE_METASTORE_EXTERNAL_WAREHOUSE","HIVE_REPLICA_WAREHOUSE","HBASE_ROOT","RANGER_AUDIT","PROFILER_SERVICE_FS_URI","DEFAULT_FS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageLocationBaseTypeTypePropEnum = append(storageLocationBaseTypeTypePropEnum, v)
	}
}

const (

	// StorageLocationBaseTypeZEPPELINNOTEBOOK captures enum value "ZEPPELIN_NOTEBOOK"
	StorageLocationBaseTypeZEPPELINNOTEBOOK string = "ZEPPELIN_NOTEBOOK"

	// StorageLocationBaseTypeYARNLOG captures enum value "YARN_LOG"
	StorageLocationBaseTypeYARNLOG string = "YARN_LOG"

	// StorageLocationBaseTypeHIVEMETASTOREWAREHOUSE captures enum value "HIVE_METASTORE_WAREHOUSE"
	StorageLocationBaseTypeHIVEMETASTOREWAREHOUSE string = "HIVE_METASTORE_WAREHOUSE"

	// StorageLocationBaseTypeHIVEMETASTOREEXTERNALWAREHOUSE captures enum value "HIVE_METASTORE_EXTERNAL_WAREHOUSE"
	StorageLocationBaseTypeHIVEMETASTOREEXTERNALWAREHOUSE string = "HIVE_METASTORE_EXTERNAL_WAREHOUSE"

	// StorageLocationBaseTypeHIVEREPLICAWAREHOUSE captures enum value "HIVE_REPLICA_WAREHOUSE"
	StorageLocationBaseTypeHIVEREPLICAWAREHOUSE string = "HIVE_REPLICA_WAREHOUSE"

	// StorageLocationBaseTypeHBASEROOT captures enum value "HBASE_ROOT"
	StorageLocationBaseTypeHBASEROOT string = "HBASE_ROOT"

	// StorageLocationBaseTypeRANGERAUDIT captures enum value "RANGER_AUDIT"
	StorageLocationBaseTypeRANGERAUDIT string = "RANGER_AUDIT"

	// StorageLocationBaseTypePROFILERSERVICEFSURI captures enum value "PROFILER_SERVICE_FS_URI"
	StorageLocationBaseTypePROFILERSERVICEFSURI string = "PROFILER_SERVICE_FS_URI"

	// StorageLocationBaseTypeDEFAULTFS captures enum value "DEFAULT_FS"
	StorageLocationBaseTypeDEFAULTFS string = "DEFAULT_FS"
)

// prop value enum
func (m *StorageLocationBase) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storageLocationBaseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StorageLocationBase) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageLocationBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageLocationBase) UnmarshalBinary(b []byte) error {
	var res StorageLocationBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}