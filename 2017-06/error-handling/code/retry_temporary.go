
for retry := 1; retry < 5; retry++ {
	if err := http.Get(url); err != nil {
		if err, ok := err.(temporaryer); ok && err.Temporary() { // HL
			time.Sleep(RETRY_INTERVAL)
			continue
		}
	}
	break // got it!
}
