// Group is a sync.WaitGroup which tracks @err, but only @sync.Once // HL
type Group struct {
	wg      sync.WaitGroup
	err     error
	errOnce sync.Once

	cancel func() // aka context.CancelFunc // HL
}
