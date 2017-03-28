// From https://github.com/nomad-software/go-channel-compendium/blob/master/signalling/wait-for-an-event/main.go
/*
In this example, a goroutine is started, does some work (in this case waiting
for five seconds) then closes the channel. Unbuffered channels always halt the
current goroutine until communication can take place. Closing the channel
signals to the goroutine that it can continue because there is no more data to
be received. Closed channels never halt execution of the goroutine.
*/
package main

import (
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Ltime)
}

func main() {
	var ch = make(chan struct{}) // HL

	go func() {
		log.Printf("entering goroutine")
		// ... do some stuff
		time.Sleep(time.Second)
		ch <- struct{}{}
		//close(ch) // HL
	}()

	// Wait for 'done' signal
	<-ch

	log.Printf("done")
}
