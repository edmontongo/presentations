package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"
)

var t = template.Must(template.ParseFiles("escaping.html"))

func main() {
	fmt.Println("HTML Escaping demo")
	fmt.Println()

	fmt.Println("Contents will be passed through the escaping.html file, first as a string and then as their specific type. Casting as an html/template subtype allows one to prevent that string from being escaped in a specific context.")

	for _, i := range data {
		wait()
		fmt.Printf("%T: %#v", i, i)
		wait()
		if err := t.Execute(os.Stdout, fmt.Sprintf("%v", i)); err != nil {
			log.Println(err)
		}
		wait()
		if err := t.Execute(os.Stdout, i); err != nil {
			log.Println(err)
		}
		fmt.Println("\n")
	}
	wait()
}

var data = []interface{}{
	"safe",
	"blue",
	template.URL("O'Reilly: How are <i>you</i>?"), // unsafe
	template.HTML("<b>Hello!</b>"),
	template.HTMLAttr(`alttext="Alt"`),
	template.JS(`(x + y * z())`),
	template.CSS(`background-color: blue;`),
}

var buf = bufio.NewReader(os.Stdin)

func wait() {
	buf.ReadLine()
}
