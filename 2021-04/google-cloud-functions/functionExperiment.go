package google_cloud_functions

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var lock = sync.Mutex{}
var count = 0

func ExperimentBlockRequests(w http.ResponseWriter, req *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	count++
	time.Sleep(time.Second * 5)
	fmt.Fprintln(w, "count1:", count)
}

func ExperimentParallelRequests(w http.ResponseWriter, req *http.Request) {
	lock.Lock()
	count++
	lock.Unlock()

	time.Sleep(time.Second * 5)
	fmt.Fprintln(w, "count concurrent:", count)
	lock.Lock()
	count--
	lock.Unlock()
}

func ExperimentGoroutines(w http.ResponseWriter, req *http.Request) {
	wait := make(chan struct{})
	go func() {
		lock.Lock()
		count++
		lock.Unlock()
		close(wait)
		for n := range make([]int, 20) {
			fmt.Println("goroutine:", n)
			time.Sleep(time.Second * 10)
		}
		lock.Lock()
		count--
		lock.Unlock()
	}()
	<-wait

	fmt.Fprintln(w, "active goroutines:", count)
}
