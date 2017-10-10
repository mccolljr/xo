// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/knq/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// RoleTableGrant represents a row from 'information_schema.role_table_grants'.
type RoleTableGrant struct {
	Grantor       pgtypes.SQLIdentifier `json:"grantor"`        // grantor
	Grantee       pgtypes.SQLIdentifier `json:"grantee"`        // grantee
	TableCatalog  pgtypes.SQLIdentifier `json:"table_catalog"`  // table_catalog
	TableSchema   pgtypes.SQLIdentifier `json:"table_schema"`   // table_schema
	TableName     pgtypes.SQLIdentifier `json:"table_name"`     // table_name
	PrivilegeType pgtypes.CharacterData `json:"privilege_type"` // privilege_type
	IsGrantable   pgtypes.YesOrNo       `json:"is_grantable"`   // is_grantable
	WithHierarchy pgtypes.YesOrNo       `json:"with_hierarchy"` // with_hierarchy
}
