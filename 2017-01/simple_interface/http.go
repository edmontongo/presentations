package interfaceTalk

import (
	"net/http"
	"strings"
)

type Game struct {
	Actioner
	StateWriter
}

func (g Game) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	//Do Authentication
	switch {
	case strings.Contains(path, "action"):
		a, err := g.Action(r.Body)
		if err != nil {
			//Handle error
			return
		}
	// Log Action
	// Do other important stuff
	case strings.Contains(path, "state"):
		// Do important stuff
	default:
		// Return an error message
		return
	}

	if err := g.WriteState(w); err != nil {
		//Handle Error
		return
	}
}
