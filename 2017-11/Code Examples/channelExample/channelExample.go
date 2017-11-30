package main

import (
	"fmt"
	"time"
)

func leader(gimmie chan string, quit chan int, letter chan rune) {
	quitAgain := make(chan int)
	go follower(letter, quitAgain)
	myResponse := "Gimmie an "
	for {
		select {
		//when there is someone to talk to, talk
		case gimmie <- myResponse:
		//if someone tells us to knock it off... go home.
		case <-quit:
			//if we receive anything on the quit channel we stop talking.
			fmt.Println("What does it spell?")
			quitAgain <- 0
			break
		}
	}
}

func follower(letter chan rune, quit chan int) {
	myResponse := 'A'
	for {
		select {
		//if we get something in we send something out our response, then add to it.
		case letter <- myResponse:
			myResponse++
		case <-quit:
			fmt.Println("Nothing... Its the Alphabet")
			return
		}
	}
}

func main() {
	gimmie := make(chan string)
	quit := make(chan int)
	letter := make(chan rune)
	go leader(gimmie, quit, letter)
	for i := 0; i < 26; i++ {
		fmt.Print(<-gimmie)
		fmt.Printf("%q\n", <-letter)
	}
	quit <- 0
	time.Sleep(time.Second)
}
