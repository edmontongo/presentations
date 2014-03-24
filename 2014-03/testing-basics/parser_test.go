package phoneparser

import "testing"

func TestEmptyStringReturnsError(t *testing.T) {
	_, err := ParseString("")
	if err == nil {
		t.Error("Expected parsing empty string to return an error")
	}
}
