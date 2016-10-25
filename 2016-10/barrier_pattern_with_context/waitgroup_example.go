package main
import ("sync"; "time")

func doWork(duration int, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(duration) * time.Second)
	println("Worker taking", duration, "sec - done.")
	wg.Done() // HL
}

func main() {
	var wg sync.WaitGroup // HL
	var start = time.Now()

	for _, val := range []int{1, 2, 3, 4} {
		wg.Add(1) // HL
		go doWork(val, &wg)
	}

	wg.Wait() // HL
	println("All done after", time.Since(start).String())
}
