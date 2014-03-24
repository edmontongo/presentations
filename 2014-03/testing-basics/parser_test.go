package phoneparser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testsThatReturnErrors = []struct {
	input       string
	description string
}{
	{"", "parsing empty string"},
	{"78012345678", "number should be too long"},
	{"10000000000", "upper threshold"},
	{"780123456", "number should be too short"},
}

func TestStringsThatReturnErrors(t *testing.T) {
	for _, test := range testsThatReturnErrors {
		_, err := ParseString(test.input)
		assert.Error(t, err, test.description)
	}
}

func TestSimpleNumber(t *testing.T) {
	number, err := ParseString("7801234567")
	assert.NoError(t, err, "Parsing simple number")
	assert.Equal(t, number, PhoneNumber(7801234567), "Expected parsed number to equal 7801234567")
}

func ExamplePhoneNumber_String() {
	fmt.Println("Phone number:", PhoneNumber(7801234567))
	// Output:
	// Phone number: (780)123-4567
}
