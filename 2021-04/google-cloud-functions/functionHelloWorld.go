package google_cloud_functions

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) { // HL
	fmt.Println("Received Hello request from the world")
	fmt.Fprintln(w, "Hello, world!")
}
