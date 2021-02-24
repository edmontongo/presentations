package main

import (
	"errors"
	"fmt"
)

type DB struct {
	Host     string
	Port     int
	Database string
	Conn     int
}

// SSENTINEL OMIT
var (
	ErrInvalidCredentials = errors.New("username or password incorrect")
	ErrPermission         = errors.New("permission problem")
	ErrQueryError         = errors.New("query error")
	ErrEmptyQuery         = errors.New("empty query")
)

// ESENTINEL OMIT

func (d *DB) Connect(user User) error {
	// mock static credentials
	// authenticated user id = 1
	username := "postgres"
	password := "postgres"
	if user.Username == username && user.Password == password {
		d.Conn = 1
		return nil
	} else {
		return ErrInvalidCredentials
	}
}

func (d *DB) Exec(q string) error {
	// SEMPTY OMIT
	if q == "" {
		return ErrEmptyQuery // HL
	}
	// EEMPTY OMIT

	// SPERROR OMIT
	if q == "perror" {
		return fmt.Errorf("query error: %s %w", q, ErrPermission) // HL
	}
	// EPERROR OMIT

	// SQERROR OMIT
	if q == "qerror" {
		return &QueryError{Query: q, Err: ErrQueryError} // HL
	}
	// EQERROR OMIT

	// SQPERROR OMIT
	if q == "qperror" {
		return &QueryError{Query: q, Err: ErrPermission} // HL
	}
	// EQPERROR OMIT
	return nil
}

// SQTYPE OMIT
type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Error() string {
	return e.Query + ": " + e.Err.Error()
}

// EQTYPE OMIT
