package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var user, pass, nenter string
	fmt.Print("username: ")
	user = "postgres"
	//fmt.Scanln(&user)
	fmt.Print("password: ")
	//fmt.Scanln(&pass)
	pass = "postgres"

	// create user
	userone := User{
		Username: user,
		Password: pass,
	}

	// create db
	db := DB{
		Host: "pgdb",
		Port: 5432,
	}

	// connect to database
	err := db.Connect(userone)
	if err != nil {
		fmt.Printf("an error occured: %v \n", err)
		os.Exit(1)
	}

	////////////////////////////
	// error checking
	// go v1.12 and previous
	////////////////////////////
	fmt.Printf("//////////\ngo v1.12 and previous\n//////////\n")

	////////////////////////////
	// normal error check
	////////////////////////////
	// S12_ERR_CHECK OMIT
	fmt.Println("\n/// normal error check")
	err = db.Exec("query")
	if err != nil { // HL
		fmt.Printf("an error occured: %v \n", err)
	}
	// E12_ERR_CHECK OMIT
	fmt.Scanln(&nenter)

	////////////////////////////
	// error: empty query
	// compare to sentinel value
	////////////////////////////
	// S12_SENT_CHECK OMIT
	fmt.Println("\n/// sentinel error value check")
	err = db.Exec("")
	if err == ErrEmptyQuery { // HL
		fmt.Printf("SENTINEL an error occured: %v \n", err)
	}
	// E12_SENT_CHECK OMIT
	fmt.Scanln(&nenter)

	////////////////////////////
	// error: permission error
	// compare to sentinel value
	////////////////////////////
	// S12_SENTF_CHECK OMIT
	fmt.Println("\n/// sentinel error value check with non sentinel value")
	err = db.Exec("perror")
	if err == ErrPermission { // HL
		fmt.Printf("SENTINEL an error occured: %v \n", err)
	}
	// E12_SENTF_CHECK OMIT
	fmt.Scanln(&nenter)

	////////////////////////////
	// error: query error type
	// checking error type QueryError
	////////////////////////////
	// S12_TYPE_CHECK OMIT
	fmt.Println("\n/// error type check")
	err = db.Exec("qerror")
	// requires type assertion
	if e, ok := err.(*QueryError); ok { // HL
		fmt.Printf("an error occured: %v \n", e)
	}
	// E12_TYPE_CHECK OMIT
	fmt.Scanln(&nenter)

	////////////////////////////
	// error: query error type
	// checking error type QueryError
	// checking sentinel value ErrPermission
	////////////////////////////
	// S12_TYPE_SENT OMIT
	err = db.Exec("qperror")
	if e, ok := err.(*QueryError); ok && e.Err == ErrQueryError { // HL
		fmt.Printf("QueryError and ErrQueryError error: %v \n", e)
	}
	if e, ok := err.(*QueryError); ok && e.Err == ErrPermission { // HL
		fmt.Printf("QueryError and ErrPermission error: %v \n", e)
	}
	// E12_TYPE_SENT OMIT
	fmt.Scanln(&nenter)

	////////////////////////////
	// wrapping errors
	////////////////////////////
	fmt.Println("\n/// wrapping errors")

	// S12_WRAP OMIT
	// only wraps text and discards error value
	err = fmt.Errorf("access denied 1: %v", ErrPermission)
	err = fmt.Errorf("access denied 2: %v", err)
	fmt.Println(err)
	if errors.Is(err, ErrPermission) { // HL
		fmt.Println("no-wrap: ErrPermission", err)
	}
	// E12_WRAP OMIT
	fmt.Scanln(&nenter)

	////////////////////////////
	// error checking
	// go v1.13+
	////////////////////////////
	fmt.Printf("//////////\ngo v1.13+\n//////////\n")

	////////////////////////////
	// normal error check
	////////////////////////////
	// S13_ERR_CHECK OMIT
	fmt.Println("\n/// normal error check")
	err = db.Exec("query")
	if err != nil { // HL
		fmt.Printf("an error occured: %v \n", err)
	}
	// E13_ERR_CHECK OMIT
	fmt.Scanln(&nenter)

	////////////////////////////
	// error: empty query
	// compare to sentinel value
	////////////////////////////
	// S13_SENT_CHECK OMIT
	fmt.Println("\n/// sentinel error value check with non sentinel value")
	err = db.Exec("")
	if err == ErrEmptyQuery { // HL
		fmt.Printf("SENTINEL an error occured: %v \n", err)
	}
	// E13_SENT_CHECK OMIT
	fmt.Scanln(&nenter)

	////////////////////////////
	// error: permission error
	// compare to sentinel value
	////////////////////////////
	// S13_SENTF_CHECK OMIT
	fmt.Println("\n/// sentinel error value check with non sentinel value")
	err = db.Exec("perror")
	if errors.Is(err, ErrPermission) { // HL
		fmt.Printf("SENTINEL an error occured: %v \n", err)
	}
	// E13_SENTF_CHECK OMIT
	fmt.Scanln(&nenter)

	////////////////////////////
	// error: query error type
	// checking error type QueryError
	////////////////////////////
	// S13_TYPE_CHECK OMIT
	fmt.Println("\n/// error type check")
	err = db.Exec("qerror")
	if errors.Is(err, ErrQueryError) { // HL
		fmt.Printf("IS an error occured: %v \n", err)
	}
	var e *QueryError
	if errors.As(err, &e) { // HL
		// err is a query error and e is set to the errors value
		fmt.Printf("AS an error occured: %v %v\n", e.Query, e.Err.Error())
	}
	// E13_TYPE_CHECK OMIT
	fmt.Scanln(&nenter)

	////////////////////////////
	// error: query error type
	// checking error type QueryError
	// checking sentinel value ErrPermission
	////////////////////////////
	// S13_TYPE_SENT OMIT
	fmt.Println("\n/// error type check with sentinel")
	err = db.Exec("qperror")
	if errors.As(err, &e) && e.Err == ErrQueryError { // HL
		fmt.Printf("an error occured: %v %v\n", e.Query, e.Err.Error())
	}
	if errors.As(err, &e) && e.Err == ErrPermission { // HL
		fmt.Printf("an error occured: %v %v\n", e.Query, e.Err.Error())
	}
	if errors.As(err, &e) && errors.Is(e.Err, ErrPermission) { // HL
		fmt.Printf("an error occured: %v %v\n", e.Query, e.Err.Error())
	}
	// E13_TYPE_SENT OMIT
	fmt.Scanln(&nenter)

	////////////////////////////
	// wrapping errors
	////////////////////////////
	fmt.Println("\n/// wrapping errors")

	// S13_WRAP_V OMIT
	// only wraps text and discards error value
	err = fmt.Errorf("access denied 1: %v", ErrPermission) // HL
	err = fmt.Errorf("access denied 2: %v", err)           // HL
	fmt.Printf("err: %v \n", err)
	if errors.Is(err, ErrPermission) {
		fmt.Println("no-wrap: ErrPermission ", err)
	}
	// E13_WRAP_V OMIT
	fmt.Scanln(&nenter)

	// S13_WRAP_W OMIT
	// with %w wraps the full error value
	err = fmt.Errorf("access denied 1: %w", ErrPermission) // HL
	err = fmt.Errorf("access denied 2: %w", err)           // HL
	fmt.Printf("err: %v \n", err)
	if errors.Is(err, ErrPermission) { // HL
		fmt.Println("wrapped: ErrPermission ", err)
	}
	// E13_WRAP_W OMIT
	fmt.Scanln(&nenter)

	////////////////////////////
	// wrapping errors depth
	////////////////////////////
	// S13_WRAP_DEPTH OMIT
	fmt.Println("\n/// wrapping errors depth")
	e1 := fmt.Errorf("level 1: %w", &QueryError{Query: "e1", Err: ErrQueryError})
	e2 := fmt.Errorf("level 2: %w", e1)
	err = fmt.Errorf("level 3: %w", e2)
	fmt.Printf("err: %v \n", err)
	if errors.As(err, &e) { // HL
		fmt.Println("wrapped: Err ", err)
	}
	// E13_WRAP_DEPTH OMIT
}
