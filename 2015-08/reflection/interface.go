package main

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

func main() {
	var i interface{}

	i = os.Stdout
	// Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	// func NewFile(fd uintptr, name string) *File

	fmt.Println("type:", reflect.TypeOf(i))
	fmt.Println("value:", reflect.ValueOf(i))

	time.Sleep(100 * time.Millisecond)
}
