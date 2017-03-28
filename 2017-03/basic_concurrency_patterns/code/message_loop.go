import "context"

func Loop(ctx context.Context) {
	for {
		select {
		case msgIn := <-cqueues.Input: // HL
			go handleMsg(ctx, msgIn)

		case msgOut := <-cqueues.Output: // HL
			go sendRetry(ctx, msgOut)

		case <-ctx.Done(): // HL
			logger.Infof("done with message loop")
			return
		}
	}
}
