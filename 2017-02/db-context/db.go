package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "host=localhost sslmode=disable")

	fmt.Println("Going to perform a long query...")

	if _, err := db.Exec("SELECT pg_sleep(10)"); err != nil {
		log.Fatal("Couldn't do the long operation:", err)
	}

	fmt.Println("That sure was hard!")
}
