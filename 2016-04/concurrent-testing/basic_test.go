package main

import "testing"

func TestPass(t *testing.T) {
	t.Log("Logs are printed when a test fails")
	if 1+1 != 2 {
		t.Error("1 + 1 does not equal 2!")
	}
	// t.Fatal("Uncomment to fail tests")
}
