package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var SumsInt = []struct {
	m        map[string]int64
	expected int64
}{
	{map[string]int64{"first": 34, "second": 12}, 46},
	{map[string]int64{"first": 1111, "second": 2222}, 3333},
}

var SumsFloat64 = []struct {
	m        map[string]float64
	expected float64
}{
	{map[string]float64{"first": 34, "second": 12}, 46},
	{map[string]float64{"first": 1111, "second": 2222}, 3333},
}

func TestSumInts(t *testing.T) {
	for _, tc := range SumsInt {
		tc := tc
		t.Run("Int", func(t *testing.T) {
			t.Parallel()
			res := SumInts(tc.m)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

func TestSumFloats(t *testing.T) {
	for _, tc := range SumsFloat64 {
		tc := tc
		t.Run("Float64", func(t *testing.T) {
			t.Parallel()
			res := SumFloats(tc.m)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

func BenchmarkSumsInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := map[string]int64{"first": 34, "second": 12}
		_ = SumInts(a)
	}
}

func BenchmarkSumsFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := map[string]float64{"first": 34, "second": 12}
		_ = SumFloats(a)
	}
}
