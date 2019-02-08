package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

var urls = []string{
	"https://www.google.com/humans.txt",
	"https://www.netflix.com/humans.txt",
	"https://medium.com/humans.txt",
}

func fetch(url string) (string, error) {
	r, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func fetchHumans(urls []string) ([]string, error) {
	var m sync.Mutex
	var wg sync.WaitGroup
	errChan := make(chan error, 1)

	resp := make([]string, len(urls))

	for index, url := range urls {
		wg.Add(1)
		go func(index int, url string) {
			defer wg.Done()

			s, err := fetch(url)
			if err != nil {
				select {
				case errChan <- err:
				default:
				}
				return
			}
			m.Lock()
			resp[index] = s
			m.Unlock()
		}(index, url)
	}
	wg.Wait()

	select {
	case err := <-errChan:
		return nil, err
	default:
	}
	return resp, nil
}

func handler(w http.ResponseWriter, req *http.Request) {
	data, err := fetchHumans(urls)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, "Bad gateway")
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, strings.Join(data, "\n\n"))
}

func main() {
	http.HandleFunc("/", handler)

	const addr = ":3000"
	log.Printf("Listening on %v", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
