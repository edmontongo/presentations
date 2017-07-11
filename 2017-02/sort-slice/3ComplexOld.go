package main

import (
	"fmt"
	"sort"
)

// This file shows sorting a simple item by name and two other fields,
// using sort.Sort.

type ItemThing struct {
	Name       string
	Fenceposts int
	FloatVal   float64
}

type ThingList []ItemThing

// SortByName is the same as before, requiring Len, Less, and Swap.
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

// How do I sort by a different field?
// Since sort.Sort requires Len, Less, and Swap, a new type is necessary so
// that they can be implemented, even though Len and Swap are the same.
type thingListByFenceposts ThingList

func (list thingListByFenceposts) Len() int {
	return len(list)
}

func (list thingListByFenceposts) Less(i, j int) bool {
	return list[i].Fenceposts < list[j].Fenceposts
}

func (list thingListByFenceposts) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list *ThingList) SortByFenceposts() {
	sort.Sort(thingListByFenceposts(*list))
}

// How do I sort by a yet another field?
// Again, a new type is necessary to implement Len, Less, and Swap.
type thingListByFloatVal ThingList

func (list thingListByFloatVal) Len() int {
	return len(list)
}

func (list thingListByFloatVal) Less(i, j int) bool {
	return list[i].FloatVal < list[j].FloatVal
}

func (list thingListByFloatVal) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list *ThingList) SortByFloatVal() {
	sort.Sort(thingListByFloatVal(*list))
}

/////////////////////////////////////////////////////////////////////
//// Everything below this point is the same as in 3ComplexNew.go ///
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
