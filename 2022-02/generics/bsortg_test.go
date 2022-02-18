package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortAlphaNum(t *testing.T) {
	for _, tc := range SortTestsInt {
		tc := tc
		t.Run("Int", func(t *testing.T) {
			t.Parallel()
			BubbleSortAlphaNum(&tc.a)
			assert.True(t, assert.ObjectsAreEqualValues(tc.expected, tc.a))
		})
	}
}

func BenchmarkBubbleSortAlphaNumInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []int{-2, 45, 0, 11, -9}
		BubbleSortAlphaNum(&a)
	}
}

func BenchmarkBubbleSortAlphaNumFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []float64{-2, 45, 0, 11, -9}
		BubbleSortAlphaNum(&a)
	}
}

func BenchmarkBubbleSortAlphaNumString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []string{"p", "e", "t", "e", "r"}
		BubbleSortAlphaNum(&a)
	}
}
