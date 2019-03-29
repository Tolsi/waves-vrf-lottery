package main

import (
	"github.com/google/keytransparency/core/crypto/vrf/p256"
)

func main() {
	m := []byte("data1")
	pk, pubk := p256.GenerateKey()
	_, proof := pk.Evaluate(m)
	_, err := pubk.ProofToHash(m, proof)
	if err != nil {
		print(err.Error())
	}
	print("Done!")
}
