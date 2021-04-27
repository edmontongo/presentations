package google_cloud_functions

import (
	"fmt"
	"net/http"
	"time"
)

// START OMIT
type LazyObject struct{ initialized bool }

func (lo *LazyObject) get() bool {
	if lo.initialized {
		return true
	}
	time.Sleep(time.Second * 3)
	lo.initialized = true
	return false
}

var lazyObject LazyObject

func LazyInit(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	initialized := lazyObject.get()
	fmt.Fprintln(w, "Duration:", time.Since(start))
	fmt.Fprintln(w, "Initialized:", initialized)
}

// END OMIT
