package phoneparser

import "testing"

func TestEmptyStringReturnsError(t *testing.T) {
	_, err := ParseString("")
	if err == nil {
		t.Error("Expected parsing empty string to return an error")
	}
}

func TestSimpleNumber(t *testing.T) {
	number, err := ParseString("7801234567")
	if err != nil {
		t.Errorf("Unexpected error '%s' when parsing simple number", err)
	}
	if number != 7801234567 {
		t.Errorf("Expected parsed number to equal 7801234567, got %d", number)
	}
}
