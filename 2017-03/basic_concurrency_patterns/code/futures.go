// From an idea by Dave Cheney (scraped from a gist)
package main

import (
	"fmt"
	"time"
)

// START OMIT
type Future func() Result // HL

type Result struct { // channels return 1 value at a time
	Val interface{}
	Err error
}

func NewFuture(f func() Result) Future { // HL
	var ch = make(chan Result, 1) // non-blocking // HL

	go func() {
		ch <- f() // HL
	}()
	return func() Result { return <-ch } // HL
}

// END OMIT

func main() {
	var future = NewFuture(func() Result {
		time.Sleep(4 * time.Second)  // pretend work
		return Result{Val: "bright"} // of course
	})

	fmt.Printf("The future is ... ")
	if res := future(); res.Err == nil { // HL
		fmt.Println(res.Val)
	}
}
