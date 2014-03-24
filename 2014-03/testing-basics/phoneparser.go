// Contrived phoneparser package. Do not use
package phoneparser

import (
	"errors"
	"strconv"
)

type PhoneNumber uint64

// Parses string for a valid phone number
func ParseString(s string) (PhoneNumber, error) {
	number, err := strconv.ParseUint(s, 10, 64)

	if number >= 10000000000 {
		return 0, errors.New("number is too long")
	}
	if number < 1000000000 {
		return 0, errors.New("number is too short")
	}

	return PhoneNumber(number), err
}
