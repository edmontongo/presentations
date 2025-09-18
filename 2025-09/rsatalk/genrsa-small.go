package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"talk/lib"
)

// start cut2
func saveBinaryAndPem(filename string, data []byte, header string) error {
	if err := ioutil.WriteFile(filename, data, 0600); err != nil {
		return err
	}

	text := pem.EncodeToMemory(&pem.Block{Type: header, Bytes: data})
	if err := ioutil.WriteFile(filename+".pem", text, 0600); err != nil {
		return err
	}
	fmt.Printf("%s\n", text)
	return nil
}

// end cut2

func main() {
	lib.MainWrapper(Main4)
}

func Main4() error {
	// start cut
	key, err := rsa.GenerateKey(rand.Reader, 24) // HL
	if err != nil {
		return err
	}

	privPkcs1 := x509.MarshalPKCS1PrivateKey(key) // HL
	if err := ioutil.WriteFile("sample-small-rsa-key", privPkcs1, 0600); err != nil {
		return err
	}

	text := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privPkcs1})
	if err := ioutil.WriteFile("sample-small-rsa-key.pem", text, 0600); err != nil {
		return err
	}
	fmt.Printf("%s\n", text)

	pubPkix, err := x509.MarshalPKIXPublicKey(&key.PublicKey) // HL
	if err != nil {
		return err
	}
	err = saveBinaryAndPem("sample-small-rsa-key.pub", pubPkix, "PUBLIC KEY")
	if err != nil {
		return err
	}
	// end cut

	//err = saveBinaryAndPem("sample-small-rsa-key", privPkcs1, "RSA PRIVATE KEY")
	//if err != nil {
	//	return err
	//}

	//err = saveBinaryAndPem("sample-small-rsa-priv-key", privPkcs1, "RSA PRIVATE KEY")
	//if err != nil {
	//	return err
	//}
	//
	//pubPkcs1 := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	//err = saveBinaryAndPem("sample-small-rsa-pub-key", pubPkcs1, "RSA PUBLIC KEY")
	//if err != nil {
	//	return err
	//}

	return nil
}
