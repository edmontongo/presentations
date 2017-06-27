type Error interface { // HL
	error
	Timeout() bool   // Is the error a timeout? // HL
	Temporary() bool // Is the error temporary? // HL
}
