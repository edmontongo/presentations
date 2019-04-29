package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// START OMIT
func main() {
	resp, err := http.Get("https://edmontongo.org")
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println(len(body))
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// END OMIT
