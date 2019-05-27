// Package p contains an HTTP Cloud Function.
package p

import (
	"fmt"
	"net/http"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloEdmonton(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello EdmontonGo!")
}
