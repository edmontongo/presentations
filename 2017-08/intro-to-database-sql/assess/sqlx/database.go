/*
This version of database.go uses jmoiron/sqlx and its
Get, Select, and NamedExec methods.
*/
package sqlx

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

const DataSource = "simplified.sqlite3"

type Database struct {
	*sqlx.DB
}

// OpenDatabase attempts to open the database specified by DataSource
// and return a handle to it
func OpenDatabase() (*Database, error) {
	db := Database{}
	var err error

	db.DB, err = sqlx.Open("sqlite3", DataSource)
	if err != nil {
		return nil, fmt.Errorf("Open (%v): %v", DataSource, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Ping: %v", err)
	}

	return &db, nil
}

// UpdateAssessments updates assessments that match 'address' with the
// specified 'value' and returns the number of updates
func (db *Database) UpdateAssessments(address string, value int) (int, error) {
	q := `UPDATE assessment
                 SET value = :value
               WHERE address LIKE UPPER(:address)`
	qArgs := struct {
		Value   int
		Address string
	}{value, address}

	res, err := db.NamedExec(q, qArgs)
	if err != nil {
		return 0, fmt.Errorf("NamedExec: %v", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("RowsAffected: %v", err)
	}

	return int(count), nil
}

// GetMinimum obtains the minimum assessed value for properties
// matching 'address' or -1 if the address matches zero properties
func (db *Database) GetMinimum(address string) (int, error) {
	// to simplify this first use of Get, do *not* use the
	// min aggregate function -- given 0 records, min produces a
	// NULL result
	q := `SELECT value
                FROM assessment
               WHERE address LIKE UPPER($1)
            ORDER BY value ASC
               LIMIT 1`
	var min int
	if err := db.Get(&min, q, address); err != nil {
		if err == sql.ErrNoRows {
			return -1, nil
		}

		return 0, fmt.Errorf("Get: %v", err)
	}

	return min, nil
}

// GetAverage obtains the average assessed value for properties
// matching 'address' or -1 if the address matches zero properties
func (db *Database) GetAverage(address string) (sql.NullFloat64, error) {
	q := `SELECT avg(value)
                FROM assessment
               WHERE address LIKE UPPER($1)`
	avg := sql.NullFloat64{} // or *float64
	if err := db.Get(&avg, q, address); err != nil {
		// the result of the aggregate function avg() is
		// always a floating-point value or NULL, i.e., we
		// don't need to worry about sql.ErrNoRows
		return sql.NullFloat64{}, fmt.Errorf("Get: %v", err)
	}

	return avg, nil
}

type Assessment struct {
	Account int    `db:"account"`
	Address string `db:"address"`
	Value   int    `db:"value"`
}

// GetAssessments obtains a slice of assessments that match 'address'
func (db *Database) GetAssessments(address string) ([]Assessment, error) {
	q := `SELECT account, address, value
                FROM assessment
               WHERE address LIKE UPPER($1)
            ORDER BY address`
	assessments := []Assessment{}

	if err := db.Select(&assessments, q, address); err != nil {
		return nil, fmt.Errorf("Select: %v", err)
	}

	return assessments, nil
}
