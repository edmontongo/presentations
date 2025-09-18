package main

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"talk/lib"
)

func main() {
	lib.MainWrapper(Main7)
}

func Main7() error {
	// start cut
	pkcs1, err := ioutil.ReadFile("sample-small-rsa-key")
	if err != nil {
		return err
	}
	privKey, err := x509.ParsePKCS1PrivateKey(pkcs1)
	if err != nil {
		return err
	}
	fmt.Printf("N aka modulus=%s\n", privKey.N)         // HL
	fmt.Printf("E aka publicExponent=%d\n", privKey.E)  // HL
	fmt.Printf("D aka privateExponent=%s\n", privKey.D) // HL
	fmt.Printf("Primes=%s\n", privKey.Primes)           // HL
	fmt.Printf("Precomputed=%s\n", privKey.Precomputed) // HL

	pkix, err := ioutil.ReadFile("sample-small-rsa-key.pub")
	if err != nil {
		return err
	}
	pkixKey, err := x509.ParsePKIXPublicKey(pkix)
	if err != nil {
		return err
	}
	pubKey := pkixKey.(*rsa.PublicKey)
	fmt.Printf("\nPublic:\nN=%v\nE=%v\n", pubKey.N, pubKey.E)
	// end cut

	return nil
}
