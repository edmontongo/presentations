package main
import ("time"; "context")

func generateNumbers(ctx context.Context) <-chan int { // HL
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; ; i++ {
			select {
			case out <- i:
			case <-ctx.Done(): // HL
				return
			}
		}
	}()
	return out
}
func main() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Millisecond) // HL
	for number := range generateNumbers(ctx) {
		println(number)
	}
}
