// NewSocket allocates a new Socket using the REQ protocol
func NewSocket() (mangos.Socket, error) {
	return mangos.MakeSocket(&req{}), nil
}

type Socket interface {
	// Dial connects a remote endpoint to the Socket.
	Dial(addr string) error

	// Listen connects a local endpoint to the Socket.
	Listen(addr string) error

	// Close closes the open Socket.
	Close() error

	// Send puts the message on the outbound send queue.
	Send([]byte) error

	// Recv receives a complete message.
	Recv() ([]byte, error)
}
