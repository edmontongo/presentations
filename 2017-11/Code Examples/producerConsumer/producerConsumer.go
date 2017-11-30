package main

import "fmt"

//makes some numbers and stuffs them into a channel it returns.
//the call to go means that this function in executed ASYNC
func produce(kickoff chan int) chan int {
	out := make(chan int) //build the channel
	go func() {           //ASYNC
		<-kickoff
		for i := 1; i <= 8; i++ {
			out <- i //push numbers into the channel
		}
		close(out) //close it to let the next pipeline function know there will be no more input
		fmt.Println("Done Producing")
	}()
	return out //return the channel
}

//this function takes 2 integers from a channel and adds them.
//from above it recieves 1,2,3,4,5,6,7,8
//it will add 1+2, 3+4, 5+6, 7+8 and then stuff them back into an output channel
func add(in, out chan int) {
	go func() {
		for a := range in {
			b := <-in
			out <- (a + b)
		}
		close(out)
		fmt.Println("Done Adding")
	}()
}

//like p1, but with multiplication
//it recieves 3, 7, 11, 15 and multiples them together.
func multiply(in, out chan int) {
	go func() {
		for a := range in {
			out <- (a * <-in)
		}
		close(out)
		fmt.Println("Done Multiplying")
	}()
}

//consume just eats up the final output of the pipeline.
//The range loop will iterate until the channel is closed.
// in the example it adds 21 + 165 = 186 and spits out the answer
func consume(consume chan int, done chan int) {
	ans := 0
	for x := range consume {
		ans += x
	}
	fmt.Println("Done Consuming")
	done <- ans

}

//while completely async, the code is deterministic, race free, and will always result in an answer of 186.
func main() {
	addOutput := make(chan int)
	multiplyOutput := make(chan int)
	doneChan := make(chan int)
	kickoff := make(chan int)

	go add(produce(kickoff), addOutput)    // we produce the input into the add function
	kickoff <- 1                           //kickoff stops the producer from being able to close before the channel is created.
	go multiply(addOutput, multiplyOutput) // we pass the output from add into multiply.
	go consume(multiplyOutput, doneChan)   // we consume and finalize the result
	fmt.Println("SEE THIS FIRST!")
	fmt.Println(<-doneChan) // this will block until multiply sends on done
}
