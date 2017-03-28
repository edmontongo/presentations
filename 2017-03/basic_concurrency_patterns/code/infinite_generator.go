package main

func Generate() <-chan int {
	var out = make(chan int)

	go func() {
		defer close(out) // Never happens! // HL

		for i := 2; ; i++ { // HL
			out <- i
		}
	}()
	return out
}

func main() {
	for i, ch := 0, Generate(); i < 20; i++ { // a hack! // HL
		println(<-ch)
	}
}
