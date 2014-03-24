// Contrived phoneparser package. Do not use
package phoneparser

import "errors"

type PhoneNumber uint64

// Parses string for a valid phone number
func ParseString(s string) (PhoneNumber, error) {
	return 0, errors.New("not implemented")
}
