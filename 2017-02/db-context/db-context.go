package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second))

	db, _ := sql.Open("postgres", "host=localhost sslmode=disable")

	fmt.Println("Going to perform a long query...")

	if _, err := db.ExecContext(ctx, "SELECT pg_sleep(10)"); err != nil {
		log.Fatal("Couldn't do the long operation:", err)
	}

	fmt.Println("That sure was hard!")
}
