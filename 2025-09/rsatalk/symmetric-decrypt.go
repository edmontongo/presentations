package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	if err := Main2(); err != nil {
		log.Fatal(err)
	}
}

func Main2() error {
	// START OMIT
	key, err := ioutil.ReadFile("sample-symmetric-key") // HL
	if err != nil {
		return err
	}

	encrypted, err := ioutil.ReadFile("sample-encrypted-message") // HL
	if err != nil {
		return err
	}

	message := make([]byte, len(encrypted))
	for i, key_byte := range key {
		message[i] = key_byte ^ encrypted[i] // HL
	}

	fmt.Printf("%s\n", message)
	// END OMIT
	return nil
}
