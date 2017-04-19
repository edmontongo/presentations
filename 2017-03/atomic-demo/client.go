package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	http.Post("http://localhost:8080/reset", "", nil)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				r, err := http.Get("http://localhost:8080")
				if err != nil {
					log.Println(err)
					return
				}
				defer r.Body.Close()
				io.Copy(ioutil.Discard, r.Body)
			}
		}()
	}
	wg.Wait()
	resp, err := http.Get("http://localhost:8080/status")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Counter is", string(buf))
}
