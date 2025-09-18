package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}

func Main() error {
	// START OMIT
	message := "secret message" // HL
	fmt.Printf("message = [% x]\n", []byte(message))

	key := make([]byte, len(message))         // HL
	if _, err := rand.Read(key); err != nil { // HL
		return err
	}
	fmt.Printf("key = [% x]\n", key)
	if err := ioutil.WriteFile("sample-symmetric-key", key, 0600); err != nil {
		return err
	}

	encrypted := make([]byte, len(message))
	for i, key_byte := range key { // HL
		encrypted[i] = key_byte ^ message[i] // HL
	}
	fmt.Printf("encrypted = [% x]\n", encrypted)

	if err := ioutil.WriteFile("sample-encrypted-message", encrypted, 0644); err != nil {
		return err
	}
	// END OMIT
	return nil
}
