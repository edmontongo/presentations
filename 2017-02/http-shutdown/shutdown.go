package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{Addr: ":8080", Handler: http.DefaultServeMux}

	go func() {
		fmt.Println("Press enter to shutdown server")
		fmt.Scanln()
		log.Println("Shutting down server...")
		if err := srv.Shutdown(context.Background()); err != nil { // HL
			log.Fatalf("could not shutdown: %v", err)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Happy Go 1.8'th")
	})
	if err := srv.ListenAndServe(); err != http.ErrServerClosed { // HL
		log.Fatalf("listen: %s\n", err)
	}
}
