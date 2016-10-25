package main
import ("sync"; "time")

func main() {
	var wg sync.WaitGroup // HL
	var start = time.Now()

	for _, val := range []int{1, 2, 3, 4} {
		wg.Add(1) // HL
		go func(duration int, wg *sync.WaitGroup) {
			time.Sleep(time.Duration(duration) * time.Second)
			println("Worker taking", duration, "sec - done.")
			wg.Done() // HL
		}(val, &wg)
	}

	wg.Wait() // HL
	println("All done after", time.Since(start).String())
}
