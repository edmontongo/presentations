package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	// BEGIN OMIT
	u, err := url.Parse("http://a b.com/")

	if err != nil {
		fmt.Println(err)         // parse http://a b.com/: invalid character " " in host name
		fmt.Printf("%#v\n", err) // &url.Error{Op:"parse", URL:"http://a b.com/", Err:" "}

		if e, ok := err.(*url.Error); ok {
			fmt.Println("Op: ", e.Op)
		}
		os.Exit(1)
	}
	// END OMIT
	fmt.Println(u)
}
