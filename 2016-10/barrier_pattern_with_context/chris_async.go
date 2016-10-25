// Op is an asynchronous operation that may fail
type Op func() error
type Ops []Op

// Run executes the operation list within a go routine
func (a Ops) Run(doneChan chan bool, errChan chan error) { // HL
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(a.ops)) // HL
		for _, op := range a.ops {
			go func(op Op) {
				if err := op(); err != nil {
					errChan <- err // HL
					return
				}
				wg.Done()
			}(op)
		}
		wg.Wait()
		doneChan <- true // HL
	}()
}
