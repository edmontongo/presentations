package main

func generateNumbers(done <-chan struct{}) <-chan int { // HL
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; ; i++ {
			select {
			case out <- i:
			case <-done: // HL
				return // HL
			}
		}
	}()
	return out
}
func main() {
	done := make(chan struct{}) // HL
	for number := range generateNumbers(done) {
		println(number)
		if number == 42 {
			close(done) // HL
		}
	}
}
