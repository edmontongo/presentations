package main

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"math/big"
	"talk/lib"
)

func main() {
	lib.MainWrapper(Main8)
}

func Main8() error {
	pkcs1, err := ioutil.ReadFile("sample-small-rsa-key") // HL
	if err != nil {
		return err
	}

	privKey, err := x509.ParsePKCS1PrivateKey(pkcs1) // HL
	if err != nil {
		return err
	}

	pkix, err := ioutil.ReadFile("sample-small-rsa-key.pub")
	if err != nil {
		return err
	}
	pkixKey, err := x509.ParsePKIXPublicKey(pkix)
	if err != nil {
		return err
	}
	pubKey := pkixKey.(*rsa.PublicKey)

	// start cut
	message := "hi!"
	m := new(big.Int).SetBytes([]byte(message))
	e := big.NewInt(int64(pubKey.E))

	// “A message is encrypted by … raising M to … e, and then taking the remainder
	// … when divided by … n”
	encrypted := new(big.Int).Exp(m, e, pubKey.N) // message ** e mod N // HL
	fmt.Printf("encrypted = %s = 0x%[1]x = 0x[% [2]x]\n", encrypted, encrypted.Bytes())

	decrypted := new(big.Int).Exp(encrypted, privKey.D, privKey.N) // HL
	fmt.Printf("Decrypts as: %s = 0x[% x] = %[2]s\n", decrypted, decrypted.Bytes())
	// end cut

	return nil
}
