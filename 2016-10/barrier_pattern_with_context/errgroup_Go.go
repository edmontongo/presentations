// Go wraps the Op in an async goroutine
func (g *Group) Go(Op func() error) { // HL
	g.wg.Add(1)

	go func() { // HL
		defer g.wg.Done()

		if err := Op(); err != nil {
			g.errOnce.Do(func() {
				g.err = err // return first error only // HL
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
	}()
}
