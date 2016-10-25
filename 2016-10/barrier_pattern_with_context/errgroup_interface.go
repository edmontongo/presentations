// Async Op runner
func (g *Group) Go(f func() error) // HL

// Wait until all have finished
func (g *Group) Wait() error // HL

// This optional, but highly recommended:
func WithContext(ctx context.Context) (*Group, context.Context) // HL
