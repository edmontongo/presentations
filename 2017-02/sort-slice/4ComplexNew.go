package main

import (
	"fmt"
	"sort"
)

// This file shows sorting a simple item by name and two other fields,
// using sort.Slice.

type ItemThing struct {
	Name       string
	Fenceposts int
	FloatVal   float64
}

type ThingList []ItemThing

// SortByName is the same as in 2SimpleNew.go.
func (tl *ThingList) SortByName() {
	sort.Slice(*tl, func(i, j int) bool {
		return (*tl)[i].Name < (*tl)[j].Name
	})
}

// SortByFenceposts shows how much code is saved by the new API sort.Slice.
// In this version, there is no need to duplicate the implementations of
// Len and Swap.
func (tl *ThingList) SortByFenceposts() {
	sort.Slice(*tl, func(i, j int) bool {
		return (*tl)[i].Fenceposts < (*tl)[j].Fenceposts
	})
}

func (tl *ThingList) SortByFloatVal() {
	sort.Slice(*tl, func(i, j int) bool {
		return (*tl)[i].FloatVal < (*tl)[j].FloatVal
	})
}

/////////////////////////////////////////////////////////////////////
//// Everything below this point is the same as in 3ComplexOld.go ///
/////////////////////////////////////////////////////////////////////

func main() {
	tl := NewThingList()

	fmt.Println(tl)

	tl.SortByName()
	fmt.Println(tl)

	tl.SortByFenceposts()
	fmt.Println(tl)

	tl.SortByFloatVal()
	fmt.Println(tl)
}

/////////////////////////////////////////////////////////////////////
// Everything below this point is also the same as in 1SimpleOld.go /
/////////////////////////////////////////////////////////////////////

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
