package main

import (
	"fmt"
	"testing"
)

func Test_Main(t *testing.T) {
	for x := 1; x < 10000; x++ {
		kickoff := make(chan int)
		addOutput := make(chan int)
		multiplyOutput := make(chan int)
		doneChan := make(chan int)

		add(produce(kickoff), addOutput) // we produce the input into the add function
		kickoff <- 1
		multiply(addOutput, multiplyOutput)  // we pass the output from add into multiply.
		go consume(multiplyOutput, doneChan) // we consume and finalize the result
		ans := <-doneChan                    // this will block until multiply sends on done                             // this will block until multiply sends on done
		if ans != 186 {
			fmt.Println(ans)
			t.Error("Uh Oh... Racy...")
		}
	}
}
