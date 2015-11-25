package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Uhoh int
type Config struct {
	General struct {
		Site Uhoh `default:"1234"`
	}
	Daemon struct {
		Address string `default:"0.0.0.0"`
		Port    int    `default:"5678"`
	}
}

// fillDefaultValue accepts a Value (that must be a pointer to some
// type) and a default value: if the value is a pointer to a string
// or an integer, it sets the value to the default
func fillDefaultValue(v reflect.Value, d string) error {
	v = reflect.Indirect(v) // dereference the pointer
	t := v.Type()           // obtain Type for that value

	switch t.Kind() {
	case reflect.Struct:
		// iterate over the fields
		for i := 0; i < v.NumField(); i++ {
			// retrieve the tag while we have access to the struct
			// (they are *struct* tags, after all)
			tag := t.Field(i).Tag.Get("default")
			if tag == "" && v.Field(i).Kind() != reflect.Struct {
				// missing tag on non-struct => skip
				continue
			}

			// make a recursive call to fill in the default
			// value (handles nested structs)
			err := fillDefaultValue(v.Field(i).Addr(), tag)
			if err != nil {
				return err
			}
		}
	case reflect.Int:
		// convert the default (string)...
		val, err := strconv.Atoi(d)
		if err != nil {
			return errors.New("invalid integer (" + d + ")")
		}

		// ... and set it
		v.SetInt(int64(val))
	case reflect.String:
		// set the default (string)
		v.SetString(d)
	default:
		return errors.New("unsupported kind: " + v.Kind().String())
	}

	return nil
}

// setDefaults validates the argument's type and passes the interface's
// Value to fillDefaultValue to operate on the struct
func setDefaults(i interface{}) error {
	// check for a pointer
	value := reflect.ValueOf(i)
	if value.Kind() != reflect.Ptr {
		return errors.New("not a pointer")
	}

	// check for a pointer *to a struct*
	value = reflect.Indirect(value) // dereference the pointer
	if value.Kind() != reflect.Struct {
		return errors.New("not a pointer to a struct")
	}

	return fillDefaultValue(reflect.ValueOf(i), "")
}

func main() {
	c := Config{}
	if err := setDefaults(&c); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", c)
}
