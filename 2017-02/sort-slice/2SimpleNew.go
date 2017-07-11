package main

import (
	"fmt"
	"sort"
)

// This file shows sorting a simple item by name, using sort.Slice.

type ItemThing struct {
	Name       string
	Fenceposts int
	FloatVal   float64
}

type ThingList []ItemThing

// SortByName now uses sort.Slice(), which slightly reduces the
// amount of code.
func (tl *ThingList) SortByName() {
	sort.Slice(*tl, func(i, j int) bool {
		return (*tl)[i].Name < (*tl)[j].Name
	})
}

/////////////////////////////////////////////////////////////////////
//// Everything below this point is the same as in 1SimpleOld.go ////
/////////////////////////////////////////////////////////////////////

func main() {
	tl := NewThingList()

	fmt.Println(tl)

	tl.SortByName()
	fmt.Println(tl)
}

// NewThingList creates a ThingList (a slice of ItemThing).
func NewThingList() ThingList {
	return ThingList{
		ItemThing{
			Name:       "Chocolat",
			Fenceposts: 23,
			FloatVal:   23.1,
		},
		ItemThing{
			Name:       "Applepie",
			Fenceposts: 8576,
			FloatVal:   -3920.86,
		},
		ItemThing{
			Name:       "Trusting",
			Fenceposts: 1,
			FloatVal:   0.0,
		},
		ItemThing{
			Name:       "Elephant",
			Fenceposts: 207,
			FloatVal:   -16.7,
		},
	}
}

// String is used to make the output print in a more readable way.
func (tl ThingList) String() string {
	s := ""
	for _, val := range tl {
		s += fmt.Sprintf("%v\n", val)
	}
	return s
}

// String is used to make the output print in a more readable way.
func (it ItemThing) String() string {
	return fmt.Sprintf("%s  %7d    %5.2f", it.Name, it.Fenceposts, it.FloatVal)
}
