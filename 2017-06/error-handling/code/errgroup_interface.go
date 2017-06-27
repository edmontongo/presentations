// Like WaitGroup.Wait() 
func (g *Group) Wait() error // HL

// Run @f asynchronously and check error return
func (g *Group) Go(f func() error) // HL

// This is optional, but highly recommended:
func WithContext(ctx context.Context) (*Group, context.Context) // HL
