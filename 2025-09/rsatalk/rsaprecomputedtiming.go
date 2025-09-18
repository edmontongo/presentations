package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"math/big"
	"talk/lib"
)

func main() {
	lib.MainWrapper(Main10)
}

type setup struct {
	bits   int
	trials int
}

func Main10() error {
	var privKey *rsa.PrivateKey
	var encrypted *big.Int
	var trials int

	message := "hi!"

	doSetup := func(s setup) error {
		trials = s.trials
		fmt.Printf("Generating %v-bit key\n", s.bits)
		lib.ShowTime(func() error {
			var err error
			privKey, err = rsa.GenerateKey(rand.Reader, s.bits)
			if err != nil {
				return err
			}
			return nil
		})
		pubKey := privKey.PublicKey

		m := new(big.Int).SetBytes([]byte(message))
		e := big.NewInt(int64(pubKey.E))

		// “A message is encrypted by … raising M to … e, and then taking the remainder
		// … when divided by … n”
		encrypted = new(big.Int).Exp(m, e, pubKey.N) // HL
		//fmt.Printf("encrypted = %s\n", encrypted)
		return nil
	}

	checkResuilt := func(r *big.Int) {
		if string(r.Bytes()) != message {
			log.Fatal("decryption failed")
		}
	}

	// start cut
	if err := doSetup(setup{trials: 100, bits: 3072}); err != nil {
		return err
	}

	lib.TimeRepeatedly(trials, func() {
		decrypted := new(big.Int).Exp(encrypted, privKey.D, privKey.N) // HL
		checkResuilt(decrypted)
	})

	lib.TimeRepeatedly(trials, func() {
		p := privKey.Primes[0]
		q := privKey.Primes[1]
		m_p := new(big.Int).Exp(encrypted, privKey.Precomputed.Dp, p) // HL
		m_q := new(big.Int).Exp(encrypted, privKey.Precomputed.Dq, q) // HL

		ret := new(big.Int)
		ret.Sub(m_p, m_q)
		ret.Mul(ret, privKey.Precomputed.Qinv)
		ret.Mul(ret, q)
		ret.Add(ret, m_q)
		ret.Mod(ret, privKey.N)
		checkResuilt(ret)
	})
	// end cut

	//fmt.Printf("Pre-computed decrypts as: %s = %v = %s\n", ret, ret.Bytes(), ret.Bytes())

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
