package main

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"math/big"
	"talk/lib"
)

// copied from crypto/rsa/rsa.go
// lightly edited too
func encrypt(pub *rsa.PublicKey, m *big.Int) *big.Int {
	c := new(big.Int)
	e := big.NewInt(int64(pub.E))
	c.Exp(m, e, pub.N)
	return c
}

func main() {
	lib.MainWrapper(Main5)
}
func Main5() error {
	// start cut
	pkcs1, err := ioutil.ReadFile("sample-small-rsa-key.pub") // HL
	if err != nil {
		return err
	}

	pubKey, err := x509.ParsePKIXPublicKey(pkcs1) // HL
	if err != nil {
		return err
	}

	m := new(big.Int).SetBytes([]byte("hi!")) // HL
	fmt.Printf("M = 0x[% x] = %#x = %s\n", m.Bytes(), m.Bytes(), m)
	// encrypt() method copied locally, is not public
	encrypted := encrypt(pubKey.(*rsa.PublicKey), m)                       // HL
	fmt.Printf("encrypted = %s = 0x[% x]\n", encrypted, encrypted.Bytes()) // HL

	err = ioutil.WriteFile("sample-rsa-encrypted-message", encrypted.Bytes(), 0644)
	if err != nil {
		return err
	}
	// end cut

	return nil
}
