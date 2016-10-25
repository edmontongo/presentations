package main
import ("fmt"; "time")

func generateNumbers() <-chan int {
	out := make(chan int)
	go func() { // HL
		for i := 0; ; i++ {
			time.Sleep(100 * time.Millisecond)
			out <- i
		}
		close(out) // HL
	}()
	return out // HL
}

func main() {
	for number := range generateNumbers() { // HL
		fmt.Println(number)
	}
}
