package main

import (
	"flag"
)

func main() {
	var someVal string
	flag.StringVar(&someVal, "someval", "default value", "description")
	flag.Parse()
}
