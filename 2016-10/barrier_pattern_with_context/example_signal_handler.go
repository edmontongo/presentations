// sigHandler broadcasts a termination signal if one of the @sigs is received.
func sigHandler(cancel context.CancelFunc, sigs ...os.Signal) {
	var sigChan = make(chan os.Signal, 1) // HL

	signal.Notify(sigChan, sigs...) // HL

	go func() {
		<-sigChan
		cancel()
	}()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // HL
	sigHandler(cancel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP)

	for ctx.Err() == nil {
		print(".")
	}

	<-ctx.Done() // HL
	println("\nMAIN DONE:", ctx.Err().Error())
}
