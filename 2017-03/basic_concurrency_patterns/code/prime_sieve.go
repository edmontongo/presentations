// From:  https://play.golang.org/p/9U22NfrXeq
// A concurrent prime sieve
package main

func Generate() <-chan int {
	var out = make(chan int)
	go func() { for i := 2; ; i++ { out <- i } }()
	return out
}

// Filter all numbers not divisible by prime from <-in into out<-
func FilterNotDivisible(in <-chan int, out chan<- int, prime int) { // HL
	for i := range in { // HL
		if i%prime != 0 {
			out <- i // HL
		}
	}
}

func main() {
	for i, inChan := 0, Generate(); i < 10; i++ { // a hack :-)
		prime := <-inChan // filtered from preceding stages // HL
		println(prime)
		outChan := make(chan int) // next stage // HL
		go FilterNotDivisible(inChan, outChan, prime)
		inChan = outChan // daisy-chain // HL
	}
}
