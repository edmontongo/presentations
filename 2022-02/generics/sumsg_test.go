package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumIntsOrFloats(t *testing.T) {
	for _, tc := range SumsInt {
		tc := tc
		t.Run("Int", func(t *testing.T) {
			t.Parallel()
			res := SumInts(tc.m)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

func BenchmarkSumIntsOrFloats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := map[string]int64{"first": 34, "second": 12}
		_ = SumIntsOrFloats(a)
	}
}

func TestSumNumbers(t *testing.T) {
	for _, tc := range SumsInt {
		tc := tc
		t.Run("Int", func(t *testing.T) {
			t.Parallel()
			res := SumNumbers(tc.m)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, res))
		})
	}
}

func BenchmarkSumNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := map[string]int64{"first": 34, "second": 12}
		_ = SumNumbers(a)
	}
}
