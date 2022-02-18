package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var SortTestsInt = []struct {
	a        []int
	expected []int
}{
	{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{[]int{-2, 45, 0, 11, -9}, []int{-9, -2, 0, 11, 45}},
}

var SortTestsFloat64 = []struct {
	a        []float64
	expected []float64
}{
	{[]float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 4, 5}},
	{[]float64{-2, 45, 0, 11, -9}, []float64{-9, -2, 0, 11, 45}},
}

var SortTestsString = []struct {
	a        []string
	expected []string
}{
	{[]string{"1", "2", "3", "4", "5"}, []string{"1", "2", "3", "4", "5"}},
	{[]string{"2", "4", "0", "1", "9"}, []string{"0", "1", "2", "4", "9"}},
}

func TestBubbleSortInt(t *testing.T) {
	for _, tc := range SortTestsInt {
		tc := tc
		t.Run("Int", func(t *testing.T) {
			t.Parallel()
			BubbleSortInt(&tc.a)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, tc.a))
		})
	}
}

func TestBubbleSortFloat64(t *testing.T) {
	for _, tc := range SortTestsInt {
		tc := tc
		t.Run("Int", func(t *testing.T) {
			t.Parallel()
			BubbleSortInt(&tc.a)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, tc.a))
		})
	}
}

func TestBubbleSortString(t *testing.T) {
	for _, tc := range SortTestsString {
		tc := tc
		t.Run("Int", func(t *testing.T) {
			t.Parallel()
			BubbleSortString(&tc.a)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, tc.a))
		})
	}
}

func BenchmarkBubbleSortInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []int{-2, 45, 0, 11, -9}
		BubbleSortInt(&a)
	}
}

func BenchmarkBubbleSortFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []float64{-2, 45, 0, 11, -9}
		BubbleSortFloat64(&a)
	}
}

func BenchmarkBubbleSortString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []string{"p", "e", "t", "e", "r"}
		BubbleSortString(&a)
	}
}
