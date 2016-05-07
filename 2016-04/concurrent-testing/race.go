package main

import (
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	b := 1
	go func() { b += 2; wg.Done() }()
	go func() { log.Println(b); wg.Done() }()
	wg.Wait()
}
