package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha512"
	"flag"
	"fmt"
	. "github.com/Tolsi/vrf-lottery/tools"
	"github.com/Tolsi/vrf-lottery/vrf"
	"github.com/btcsuite/btcutil/base58"
)

func main() {
	//region Read params

	seed := flag.String("seed", "", "A base58 seed phrase used to generate keys, optional")

	flag.Parse()

	if *seed == "" {
		key := make([]byte, 64)

		_, err := rand.Read(key)
		if err != nil {
			panic(err)
		}
		tmp := base58.Encode(key)
		seed = &tmp
	}

	seedBytes := base58.Decode(*seed)

	//endregion

	//region Create keys

	hasher := sha512.New()
	hasher.Write(seedBytes)
	private, err := vrf.GenerateKey(bytes.NewReader(hasher.Sum(nil)))
	PrintErrorAndExit(err)
	public, _ := private.Public()

	//endregion

	//region Result output

	fmt.Printf("Seed: %s\n", *seed)
	fmt.Printf("Private key: %s\n", base58.Encode(private))
	fmt.Printf("Public key: %s\n", base58.Encode(public))

	//endregion
}
