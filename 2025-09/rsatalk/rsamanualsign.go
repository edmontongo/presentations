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
	lib.MainWrapper(Main18)
}

func Main18() error {
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

	decrypted := new(big.Int).Exp(m, privKey.D, pubKey.N) // message ** d mod N // HL
	fmt.Printf("decrypted = %s = 0x[% x]\n", decrypted, decrypted.Bytes())

	e := big.NewInt(int64(pubKey.E))
	verified := new(big.Int).Exp(decrypted, e, pubKey.N) // decrypted ** e mod N // HL
	fmt.Printf("decrypted ** e = %s = %#x = 0x[% [2]x] = %[2]s\n", verified, verified.Bytes())
	// end cut

	return nil
}
