package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	i := 42

	fmt.Println("type:", reflect.TypeOf(i))
	fmt.Println("value:", reflect.ValueOf(i))

	p := &i

	fmt.Println("type:", reflect.TypeOf(p))
	fmt.Println("value:", reflect.ValueOf(p))

	time.Sleep(100 * time.Millisecond)
}
