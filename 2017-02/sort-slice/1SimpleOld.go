package main

import (
	"fmt"
	"sort"
)

// This file shows sorting a simple item by name, using sort.Sort.

type ItemThing struct {
	Name       string
	Fenceposts int
	FloatVal   float64
}

type ThingList []ItemThing

// SortByName uses sort.Sort(), which requires ThingList to implement
// Len, Less, and Swap.
func (list *ThingList) SortByName() {
	sort.Sort(list)
}

func (list ThingList) Len() int {
	return len(list)
}

func (list ThingList) Less(i, j int) bool {
	return list[i].Name < list[j].Name
}

func (list ThingList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

/////////////////////////////////////////////////////////////////////
//// Everything below this point is the same as in 2SimpleNew.go ////
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
