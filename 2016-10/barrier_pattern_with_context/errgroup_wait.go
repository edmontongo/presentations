func (g *Group) Wait() error {
	g.wg.Wait() // HL
	if g.cancel != nil {
		g.cancel()
	}
	return g.err // HL
}
