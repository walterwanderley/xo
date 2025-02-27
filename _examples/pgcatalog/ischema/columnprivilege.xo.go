package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"github.com/xo/xo/_examples/pgcatalog/pgtypes"
)

// ColumnPrivilege represents a row from 'information_schema.column_privileges'.
type ColumnPrivilege struct {
	Grantor       pgtypes.SQLIdentifier `json:"grantor"`        // grantor
	Grantee       pgtypes.SQLIdentifier `json:"grantee"`        // grantee
	TableCatalog  pgtypes.SQLIdentifier `json:"table_catalog"`  // table_catalog
	TableSchema   pgtypes.SQLIdentifier `json:"table_schema"`   // table_schema
	TableName     pgtypes.SQLIdentifier `json:"table_name"`     // table_name
	ColumnName    pgtypes.SQLIdentifier `json:"column_name"`    // column_name
	PrivilegeType pgtypes.CharacterData `json:"privilege_type"` // privilege_type
	IsGrantable   pgtypes.YesOrNo       `json:"is_grantable"`   // is_grantable

}
