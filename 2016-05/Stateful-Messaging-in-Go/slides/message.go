type Message struct {
    Header []byte	// routing information
    Port   Port		// endpoint information
    Body   []byte	// the wrapped message
}
