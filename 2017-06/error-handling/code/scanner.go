package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func init() { log.SetFlags(0) }

func main() {
	var count uint64

	const input = "Not every statement has to be about error handling"
	var scanner = bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() { // HL
		count++
		fmt.Printf("%d: %s\n", count, scanner.Text())
	} // HL

	if err := scanner.Err(); err != nil { // Error tested at end // HL
		log.Fatalf("failed to scan input: %s", err)
	}
	fmt.Printf("Counted %d words without error\n", count)
}
