package main

import "fmt"

// START OMIT
type Foo struct{ Bar, Baz string }
type OptionFunc func(f *Foo)

func New(options ...OptionFunc) *Foo {
	f := &Foo{Bar: "default bar", Baz: "default baz"} // HL
	for _, opt := range options {
		opt(f)
	}
	return f
}

func WithBar(bar string) OptionFunc {
	return func(f *Foo) { f.Bar = bar }
}

func main() {
	f := New(WithBar("the bar"), // HL
		func(f *Foo) { f.Baz = "what's this?" }) // HL
	fmt.Printf("%+v", f)
}

// END OMIT
