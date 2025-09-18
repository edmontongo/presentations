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
	lib.MainWrapper(Main20)
}

func Main20() error {
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

	message := "hi!"
	m := new(big.Int).SetBytes([]byte(message))
	e := big.NewInt(int64(pubKey.E))

	// “A message is encrypted by … raising M to … e, and then taking the remainder
	// … when divided by … n”
	encrypted := new(big.Int).Exp(m, e, pubKey.N) // HL
	fmt.Printf("encrypted = 0x[% x] = %s\n", encrypted.Bytes(), encrypted)

	//p := privKey.Primes[0]

	//fmt.Printf("Decrypts as: %s = %v = %s\n", decrypted, decrypted.Bytes(), decrypted.Bytes())

	p := privKey.Primes[0]
	q := privKey.Primes[1]
	// start cut
	decrypted := new(big.Int).Exp(encrypted, privKey.D, privKey.N) // HL
	fmt.Printf("encrypted ** d = %s ≡ %s (mod p), ≡ %s (mod q)\n",
		decrypted, new(big.Int).Mod(decrypted, p), new(big.Int).Mod(decrypted, q))

	m_p := new(big.Int).Exp(encrypted, privKey.Precomputed.Dp, p) // message mod p // HL
	m_q := new(big.Int).Exp(encrypted, privKey.Precomputed.Dq, q) // message mod q // HL
	fmt.Printf("encrypted ** dp ≡ %s (mod p), encrypted ** dq  ≡ %s (mod q)\n", m_p, m_q)

	ret := new(big.Int)
	// q * iqmp * (m_p - m_q) // HL
	ret.Sub(m_p, m_q)
	ret.Mul(ret, privKey.Precomputed.Qinv) // HL
	ret.Mul(ret, q)
	// ret + m_q (mod n) // HL
	ret.Add(ret, m_q)
	ret.Mod(ret, privKey.N)

	fmt.Printf("Pre-computed decrypts as: %s = 0x[% x] = %[2]s\n", ret, ret.Bytes())
	// end cut

	//
	//fmt.Printf("q * qInv %% p = %s\n",
	//	new(big.Int).Mod(
	//		new(big.Int).Mul(q, privKey.Precomputed.Qinv),
	//		p))

	//new(big.Int)

	//new(big.Int).Mul(
	//	privKey.Primes[1],
	//	new(big.Int).Mul(
	//		privKey.Precomputed.Qinv
	//new(big.Int).Sub(m_p, m_q)

	return nil
}
