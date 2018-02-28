package main

import "testing"

func TestCanDo(t *testing.T) {
	if !canDo() {
		t.Fail()
	}
}
