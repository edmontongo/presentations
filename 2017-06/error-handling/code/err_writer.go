func (ew *errWriter) write(buf []byte) {
	if ew.err != nil { // no-op if earlier error // HL
		return
	} // HL
	_, ew.err = ew.w.Write(buf)
}

