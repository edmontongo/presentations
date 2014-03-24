// Contrived phoneparser package. Do not use
package phoneparser

import "strconv"

type PhoneNumber uint64

// Parses string for a valid phone number
func ParseString(s string) (PhoneNumber, error) {
	number, err := strconv.ParseUint(s, 10, 64)

	return PhoneNumber(number), err
}
