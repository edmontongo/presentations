package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"talk/lib"
)

func main() {
	lib.MainWrapperWithTime(Main3)
}
func Main3() error {
	// start cut
	key, err := rsa.GenerateKey(rand.Reader, 3072) // HL
	if err != nil {
		return err
	}

	pem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
	fmt.Printf("%s\n", pem)
	// end cut

	return nil
}
