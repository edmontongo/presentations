package main

import (
	"fmt"
	"reflect"
)

func isZero(i interface{}) bool {
	// extract the Type struct from the argument (interface)
	iType := reflect.TypeOf(i)
	// obtain the (zero) Value struct for this type
	zeroValue := reflect.Zero(iType)

	// extract the Value struct for the argument (interface)
	iValue := reflect.ValueOf(i)

	// Can you guess why I can't write...
	// return zeroValue == iValue

	// How about...
	//return iValue.Interface() == zeroValue.Interface()

	return reflect.DeepEqual(
		zeroValue.Interface(),
		iValue.Interface())
}

func main() {
	var v []int
	fmt.Println(isZero(v))
}
