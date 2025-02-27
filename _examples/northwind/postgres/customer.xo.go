package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Customer represents a row from 'public.customers'.
type Customer struct {
	CustomerID   string         `json:"customer_id"`   // customer_id
	CompanyName  string         `json:"company_name"`  // company_name
	ContactName  sql.NullString `json:"contact_name"`  // contact_name
	ContactTitle sql.NullString `json:"contact_title"` // contact_title
	Address      sql.NullString `json:"address"`       // address
	City         sql.NullString `json:"city"`          // city
	Region       sql.NullString `json:"region"`        // region
	PostalCode   sql.NullString `json:"postal_code"`   // postal_code
	Country      sql.NullString `json:"country"`       // country
	Phone        sql.NullString `json:"phone"`         // phone
	Fax          sql.NullString `json:"fax"`           // fax
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Customer exists in the database.
func (c *Customer) Exists() bool {
	return c._exists
}

// Deleted returns true when the Customer has been marked for deletion from
// the database.
func (c *Customer) Deleted() bool {
	return c._deleted
}

// Insert inserts the Customer to the database.
func (c *Customer) Insert(ctx context.Context, db DB) error {
	switch {
	case c._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case c._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (basic)
	const sqlstr = `INSERT INTO public.customers (` +
		`customer_id, company_name, contact_name, contact_title, address, city, region, postal_code, country, phone, fax` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`)`
	// run
	logf(sqlstr, c.CustomerID, c.CompanyName, c.ContactName, c.ContactTitle, c.Address, c.City, c.Region, c.PostalCode, c.Country, c.Phone, c.Fax)
	if err := db.QueryRowContext(ctx, sqlstr, c.CustomerID, c.CompanyName, c.ContactName, c.ContactTitle, c.Address, c.City, c.Region, c.PostalCode, c.Country, c.Phone, c.Fax).Scan(&c.CustomerID); err != nil {
		return logerror(err)
	}
	// set exists
	c._exists = true
	return nil
}

// Update updates a Customer in the database.
func (c *Customer) Update(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case c._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.customers SET (` +
		`company_name, contact_name, contact_title, address, city, region, postal_code, country, phone, fax` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10` +
		`) WHERE customer_id = $11`
	// run
	logf(sqlstr, c.CompanyName, c.ContactName, c.ContactTitle, c.Address, c.City, c.Region, c.PostalCode, c.Country, c.Phone, c.Fax, c.CustomerID)
	if _, err := db.ExecContext(ctx, sqlstr, c.CompanyName, c.ContactName, c.ContactTitle, c.Address, c.City, c.Region, c.PostalCode, c.Country, c.Phone, c.Fax, c.CustomerID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Customer to the database.
func (c *Customer) Save(ctx context.Context, db DB) error {
	if c.Exists() {
		return c.Update(ctx, db)
	}
	return c.Insert(ctx, db)
}

// Upsert performs an upsert for Customer.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Customer) Upsert(ctx context.Context, db DB) error {
	switch {
	case c._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.customers (` +
		`customer_id, company_name, contact_name, contact_title, address, city, region, postal_code, country, phone, fax` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) ON CONFLICT (customer_id) DO UPDATE SET (` +
		`customer_id, company_name, contact_name, contact_title, address, city, region, postal_code, country, phone, fax` +
		`) = (` +
		`EXCLUDED.customer_id, EXCLUDED.company_name, EXCLUDED.contact_name, EXCLUDED.contact_title, EXCLUDED.address, EXCLUDED.city, EXCLUDED.region, EXCLUDED.postal_code, EXCLUDED.country, EXCLUDED.phone, EXCLUDED.fax` +
		`)`
	// run
	logf(sqlstr, c.CustomerID, c.CompanyName, c.ContactName, c.ContactTitle, c.Address, c.City, c.Region, c.PostalCode, c.Country, c.Phone, c.Fax)
	if _, err := db.ExecContext(ctx, sqlstr, c.CustomerID, c.CompanyName, c.ContactName, c.ContactTitle, c.Address, c.City, c.Region, c.PostalCode, c.Country, c.Phone, c.Fax); err != nil {
		return err
	}
	// set exists
	c._exists = true
	return nil
}

// Delete deletes the Customer from the database.
func (c *Customer) Delete(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return nil
	case c._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.customers WHERE customer_id = $1`
	// run
	logf(sqlstr, c.CustomerID)
	if _, err := db.ExecContext(ctx, sqlstr, c.CustomerID); err != nil {
		return logerror(err)
	}
	// set deleted
	c._deleted = true
	return nil
}

// CustomerByCustomerID retrieves a row from 'public.customers' as a Customer.
//
// Generated from index 'customers_pkey'.
func CustomerByCustomerID(ctx context.Context, db DB, customerID string) (*Customer, error) {
	// query
	const sqlstr = `SELECT ` +
		`customer_id, company_name, contact_name, contact_title, address, city, region, postal_code, country, phone, fax ` +
		`FROM public.customers ` +
		`WHERE customer_id = $1`
	// run
	logf(sqlstr, customerID)
	c := Customer{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, customerID).Scan(&c.CustomerID, &c.CompanyName, &c.ContactName, &c.ContactTitle, &c.Address, &c.City, &c.Region, &c.PostalCode, &c.Country, &c.Phone, &c.Fax); err != nil {
		return nil, logerror(err)
	}
	return &c, nil
}
