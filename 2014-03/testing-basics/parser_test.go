package phoneparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyStringReturnsError(t *testing.T) {
	_, err := ParseString("")
	assert.Error(t, err, "parsing empty string")
}

func TestSimpleNumber(t *testing.T) {
	number, err := ParseString("7801234567")
	assert.NoError(t, err, "Parsing simple number")
	assert.Equal(t, number, PhoneNumber(7801234567), "Expected parsed number to equal 7801234567")
}
