package interfaceTalk

import "net/http"

func (a Action) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a, err := c.Action(r.Body)
	if err != nil {
		//Handle error
		return
	}
	// Log Action
	// Do other important stuff
}

func (s StateWriter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := c.State(w); err != nil {
		//Handle Error
		return
	}
}
