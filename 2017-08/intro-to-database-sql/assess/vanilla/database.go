/*
This version of database.go uses only database/sql for interacting
with the database.
*/
package vanilla

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const DataSource = "simplified.sqlite3"

type Database struct {
	*sql.DB
}

// OpenDatabase attempts to open the database specified by DataSource
// and return a handle to it
func OpenDatabase() (*Database, error) {
	db := Database{}
	var err error

	db.DB, err = sql.Open("sqlite3", DataSource)
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
                 SET value = $1
               WHERE address LIKE UPPER($2)`
	res, err := db.Exec(q, value, address)
	if err != nil {
		return 0, fmt.Errorf("Exec: %v", err)
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
	// to simplify this first use of QueryRow, do *not* use the
	// min aggregate function -- given 0 records, min produces a
	// NULL result
	q := `SELECT value
                FROM assessment
               WHERE address LIKE UPPER($1)
            ORDER BY value ASC
               LIMIT 1`
	var min int
	if err := db.QueryRow(q, address).Scan(&min); err != nil {
		if err == sql.ErrNoRows {
			return -1, nil
		}

		return 0, fmt.Errorf("QueryRow: %v", err)
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
	if err := db.QueryRow(q, address).Scan(&avg); err != nil {
		// the result of the aggregate function avg() is
		// always a floating-point value or NULL, i.e., we
		// don't need to worry about sql.ErrNoRows
		return sql.NullFloat64{}, fmt.Errorf("QueryRow: %v", err)
	}

	return avg, nil
}

type Assessment struct {
	Account int
	Address string
	Value   int
}

// GetAssessments obtains a slice of assessments that match 'address'
func (db *Database) GetAssessments(address string) ([]Assessment, error) {
	q := `SELECT account, address, value
                FROM assessment
               WHERE address LIKE UPPER($1)
            ORDER BY address`
	rows, err := db.Query(q, address)
	if err != nil {
		return nil, fmt.Errorf("Query: %v", err)
	}
	defer rows.Close()

	assessments := []Assessment{}
	for rows.Next() {
		a := Assessment{}
		if err := rows.Scan(&a.Account, &a.Address, &a.Value); err != nil {
			return nil, fmt.Errorf("Scan: %v", err)
		}

		assessments = append(assessments, a)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Next: %v", err)
	}

	return assessments, nil
}
