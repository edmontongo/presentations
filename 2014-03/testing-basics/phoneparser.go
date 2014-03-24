// Contrived phoneparser package. Do not use
package phoneparser

import (
	"errors"
	"fmt"
	"strconv"
)

type PhoneNumber uint64

func (ph PhoneNumber) String() string {
	area := ph / 10000000
	local := (ph % 10000000) / 10000
	digits := ph % 10000
	return fmt.Sprintf("(%03d)%03d-%04d", area, local, digits)
}

// Parses string for a valid phone number
func ParseString(s string) (PhoneNumber, error) {
	number, err := strconv.ParseUint(s, 10, 64)

	if number >= 10000000000 {
		return 0, errors.New("number is too long")
	}
	if number < 1000000000 {
		return 0, errors.New("number is too short")
	}
	if number == 7807654321 {
		return 0, errors.New("arbitrarily invalid number")
	}

	return PhoneNumber(number), err
}
