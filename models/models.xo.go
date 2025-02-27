// Package models contains generated code for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/csv"
	"fmt"
	"io"
	"regexp"
	"strings"
)

var (
	// logf is used by generated code to log SQL queries.
	logf = func(string, ...interface{}) {}
	// errf is used by generated code to log SQL errors.
	errf = func(string, ...interface{}) {}
)

// logerror logs the error and returns it.
func logerror(err error) error {
	errf("ERROR: %v", err)
	return err
}

// Logf logs a message using the package logger.
func Logf(s string, v ...interface{}) {
	logf(s, v...)
}

// SetLogger sets the package logger. Valid logger types:
//
//     io.Writer
//     func(string, ...interface{}) (int, error) // fmt.Printf
//     func(string, ...interface{}) // log.Printf
//
func SetLogger(logger interface{}) {
	logf = convLogger(logger)
}

// Errorf logs an error message using the package error logger.
func Errorf(s string, v ...interface{}) {
	errf(s, v...)
}

// SetErrorLogger sets the package error logger. Valid logger types:
//
//     io.Writer
//     func(string, ...interface{}) (int, error) // fmt.Printf
//     func(string, ...interface{}) // log.Printf
//
func SetErrorLogger(logger interface{}) {
	errf = convLogger(logger)
}

// convLogger converts logger to the standard logger interface.
func convLogger(logger interface{}) func(string, ...interface{}) {
	switch z := logger.(type) {
	case io.Writer:
		return func(s string, v ...interface{}) {
			fmt.Fprintf(z, s, v...)
		}
	case func(string, ...interface{}) (int, error): // fmt.Printf
		return func(s string, v ...interface{}) {
			_, _ = z(s, v...)
		}
	case func(string, ...interface{}): // log.Printf
		return z
	}
	panic(fmt.Sprintf("unsupported logger type %T", logger))
}

// DB is the common interface for database operations that can be used with
// types from schema 'public'.
//
// This works with both database/sql.DB and database/sql.Tx.
type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// Error is an error.
type Error string

// Error satisfies the error interface.
func (err Error) Error() string {
	return string(err)
}

// Error values.
const (
	// ErrAlreadyExists is the already exists error.
	ErrAlreadyExists Error = "already exists"
	// ErrDoesNotExist is the does not exist error.
	ErrDoesNotExist Error = "does not exist"
	// ErrMarkedForDeletion is the marked for deletion error.
	ErrMarkedForDeletion Error = "marked for deletion"
)

// ErrInsertFailed is the insert failed error.
type ErrInsertFailed struct {
	Err error
}

// Error satisfies the error interface.
func (err *ErrInsertFailed) Error() string {
	return fmt.Sprintf("insert failed: %v", err.Err)
}

// Unwrap satisfies the unwrap interface.
func (err *ErrInsertFailed) Unwrap() error {
	return err.Err
}

// ErrUpdateFailed is the update failed error.
type ErrUpdateFailed struct {
	Err error
}

// Error satisfies the error interface.
func (err *ErrUpdateFailed) Error() string {
	return fmt.Sprintf("update failed: %v", err.Err)
}

// Unwrap satisfies the unwrap interface.
func (err *ErrUpdateFailed) Unwrap() error {
	return err.Err
}

// ErrUpsertFailed is the upsert failed error.
type ErrUpsertFailed struct {
	Err error
}

// Error satisfies the error interface.
func (err *ErrUpsertFailed) Error() string {
	return fmt.Sprintf("upsert failed: %v", err.Err)
}

// Unwrap satisfies the unwrap interface.
func (err *ErrUpsertFailed) Unwrap() error {
	return err.Err
}

// ErrDecodeFailed is the decode failed error.
type ErrDecodeFailed struct {
	Err error
}

// Error satisfies the error interface.
func (err *ErrDecodeFailed) Error() string {
	return fmt.Sprintf("unable to decode: %v", err.Err)
}

// Unwrap satisfies the unwrap interface.
func (err *ErrDecodeFailed) Unwrap() error {
	return err.Err
}

// ErrInvalidStringSlice is the invalid StringSlice error.
const ErrInvalidStringSlice Error = "invalid StringSlice"

// StringSlice is a slice of strings.
type StringSlice []string

// Scan satisfies the sql.Scanner interface for StringSlice.
func (ss *StringSlice) Scan(v interface{}) error {
	buf, ok := v.([]byte)
	if !ok {
		return logerror(ErrInvalidStringSlice)
	}
	// change quote escapes for csv parser
	str := strings.Replace(quoteEscRE.ReplaceAllString(string(buf), `$1""`), `\\`, `\`, -1)
	str = str[1 : len(str)-1]
	// bail if only one
	if len(str) == 0 {
		return nil
	}
	// parse with csv reader
	r := csv.NewReader(strings.NewReader(str))
	line, err := r.Read()
	if err != nil {
		return logerror(&ErrDecodeFailed{err})
	}
	*ss = StringSlice(line)
	return nil
}

// quoteEscRE matches escaped characters in a string.
var quoteEscRE = regexp.MustCompile(`([^\\]([\\]{2})*)\\"`)

// Value satisfies the sql/driver.Valuer interface.
func (ss StringSlice) Value() (driver.Value, error) {
	v := make([]string, len(ss))
	for i, s := range ss {
		v[i] = `"` + strings.Replace(strings.Replace(s, `\`, `\\\`, -1), `"`, `\"`, -1) + `"`
	}
	return "{" + strings.Join(v, ",") + "}", nil
}

// PostgresSchema retrieves the current schema.
func PostgresSchema(ctx context.Context, db DB) (string, error) {
	// query
	const sqlstr = `SELECT CURRENT_SCHEMA()`
	// run
	logf(sqlstr)
	var schemaName string
	if err := db.QueryRowContext(ctx, sqlstr).Scan(&schemaName); err != nil {
		return "", logerror(err)
	}
	return schemaName, nil
}

// MysqlSchema retrieves the current schema.
func MysqlSchema(ctx context.Context, db DB) (string, error) {
	// query
	const sqlstr = `SELECT SCHEMA() AS schema_name`
	// run
	logf(sqlstr)
	var schemaName string
	if err := db.QueryRowContext(ctx, sqlstr).Scan(&schemaName); err != nil {
		return "", logerror(err)
	}
	return schemaName, nil
}

// Sqlite3Schema retrieves the current schema.
func Sqlite3Schema(ctx context.Context, db DB) (string, error) {
	// query
	const sqlstr = `SELECT REPLACE(file, RTRIM(file, REPLACE(file, '/', '')), '') AS schema_name FROM pragma_database_list()`
	// run
	logf(sqlstr)
	var schemaName string
	if err := db.QueryRowContext(ctx, sqlstr).Scan(&schemaName); err != nil {
		return "", logerror(err)
	}
	return schemaName, nil
}

// SqlserverSchema retrieves the current schema.
func SqlserverSchema(ctx context.Context, db DB) (string, error) {
	// query
	const sqlstr = `SELECT SCHEMA_NAME() AS schema_name`
	// run
	logf(sqlstr)
	var schemaName string
	if err := db.QueryRowContext(ctx, sqlstr).Scan(&schemaName); err != nil {
		return "", logerror(err)
	}
	return schemaName, nil
}

// OracleSchema retrieves the current schema.
func OracleSchema(ctx context.Context, db DB) (string, error) {
	// query
	const sqlstr = `SELECT LOWER(SYS_CONTEXT('USERENV', 'CURRENT_SCHEMA')) AS schema_name FROM dual`
	// run
	logf(sqlstr)
	var schemaName string
	if err := db.QueryRowContext(ctx, sqlstr).Scan(&schemaName); err != nil {
		return "", logerror(err)
	}
	return schemaName, nil
}
