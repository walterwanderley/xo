package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"github.com/xo/xo/_examples/pgcatalog/pgtypes"
)

// PgTruetypmod calls the stored procedure 'information_schema._pg_truetypmod(pg_attribute, pg_type) integer' on db.
func PgTruetypmod(ctx context.Context, db DB, v0 pgtypes.PgAttribute, v1 pgtypes.PgType) (int, error) {
	// call information_schema._pg_truetypmod
	const sqlstr = `SELECT information_schema._pg_truetypmod($1, $2)`
	// run
	var ret int
	logf(sqlstr, v0, v1)
	if err := db.QueryRowContext(ctx, sqlstr, v0, v1).Scan(&ret); err != nil {
		return 0, logerror(err)
	}
	return ret, nil
}
