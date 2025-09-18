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
	lib.MainWrapper(Main9)
}

func Main9() error {
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
	e := big.NewInt(int64(pubKey.E))

	// start cut
	eTimesD := new(big.Int).Mul(e, privKey.D)
	fmt.Printf("e * d = %s * %s = %s\n", e, privKey.D, eTimesD)
	secretModulus := new(big.Int).Mul( // (p - 1) × (q - 1) // HL
		new(big.Int).Sub(privKey.Primes[0], big.NewInt(1)),
		new(big.Int).Sub(privKey.Primes[1], big.NewInt(1)))
	fmt.Printf("e * d ≡ %s (mod (p - 1)(q - 1))\n",
		new(big.Int).Mod(eTimesD, secretModulus))

	// Using built-in function to compute
	fmt.Printf("\ne ** -1 ≡ %s (mod (p - 1)(q - 1))\n",
		new(big.Int).ModInverse(e, secretModulus)) // HL
	// end cut

	return nil
}
