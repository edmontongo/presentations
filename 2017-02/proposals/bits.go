// LeadingZerosN returns the number of leading zero bits in x.
func LeadingZeros(x uint) int

// TrailingZerosN returns the number of trailing zero bits in x.
func TrailingZeros(x uint) int

// OnesCountN returns the number of one bits ("population count") in x.
func OnesCount(x uint) int

// RotateLeftN returns the value of x rotated left by k bits; k must not be negative.
func RotateLeft(x uint, k int) uint
func RotateRight(x uint, k int) uint

// ReverseN returns the value of x with its bits in reversed order.
func Reverse(x uint) uint

// ReverseBytesN returns the value of x with its bytes in reversed order.
func ReverseBytes(x uint) uint

// LenN returns the minimum number of bits required to represent x.
func Len(x uint) int
