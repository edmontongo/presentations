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
	lib.MainWrapper(Main6)
}

// not public
// decrypt performs an RSA decryption, resulting in a plaintext integer. If a
// random source is given, RSA blinding is used.
func decrypt(priv *rsa.PrivateKey, c *big.Int) (m *big.Int, err error) {
	// TODO(agl): can we get away with reusing blinds?
	if c.Cmp(priv.N) > 0 {
		err = rsa.ErrDecryption
		return
	}
	if priv.N.Sign() == 0 {
		return nil, rsa.ErrDecryption
	}

	var ir *big.Int

	if priv.Precomputed.Dp == nil {
		m = new(big.Int).Exp(c, priv.D, priv.N)
	} else {
		// We have the precalculated values needed for the CRT.
		m = new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
		m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
		m.Sub(m, m2)
		if m.Sign() < 0 {
			m.Add(m, priv.Primes[0])
		}
		m.Mul(m, priv.Precomputed.Qinv)
		m.Mod(m, priv.Primes[0])
		m.Mul(m, priv.Primes[1])
		m.Add(m, m2)

		for i, values := range priv.Precomputed.CRTValues {
			prime := priv.Primes[2+i]
			m2.Exp(c, values.Exp, prime)
			m2.Sub(m2, m)
			m2.Mul(m2, values.Coeff)
			m2.Mod(m2, prime)
			if m2.Sign() < 0 {
				m2.Add(m2, prime)
			}
			m2.Mul(m2, values.R)
			m.Add(m, m2)
		}
	}

	if ir != nil {
		// Unblind.
		m.Mul(m, ir)
		m.Mod(m, priv.N)
	}

	return
}

func Main6() error {
	// start cut
	pkcs1, err := ioutil.ReadFile("sample-small-rsa-key") // HL
	if err != nil {
		return err
	}

	privKey, err := x509.ParsePKCS1PrivateKey(pkcs1) // HL
	if err != nil {
		return err
	}

	encrypted, err := ioutil.ReadFile("sample-rsa-encrypted-message") // HL
	if err != nil {
		return err
	}
	fmt.Printf("encrypted = 0x[% x]\n", encrypted)

	// encrypt() method copied locally, is not public
	decrypted, err := decrypt(privKey, new(big.Int).SetBytes(encrypted)) // HL
	if err != nil {
		return err
	}

	fmt.Printf("%s = 0x[% x] = %[2]s\n", decrypted, decrypted.Bytes())
	// end cut

	return nil
}
