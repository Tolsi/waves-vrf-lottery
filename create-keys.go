package main

import (
	"bytes"
	"crypto/sha512"
	"flag"
	"fmt"
	. "github.com/Tolsi/vrf-lottery/tools"
	"github.com/Tolsi/vrf-lottery/vrf"
	"github.com/btcsuite/btcutil/base58"
	"os"
)

func main() {
	//region Read params

	seed := *flag.String("seed", "", "seed phrase used to generate keys")
	if seed == "" {
		flag.Usage()
		os.Exit(1)
	}

	//endregion

	//region Create keys

	hasher := sha512.New()
	hasher.Write([]byte(seed))
	private, err := vrf.GenerateKey(bytes.NewReader(hasher.Sum(nil)))
	PrintErrorAndExit(err)
	public, _ := private.Public()

	//endregion

	//region Result output

	fmt.Printf("Private key: %s\n", base58.Encode(private))
	fmt.Printf("Public key: %s\n", base58.Encode(public))

	//endregion
}
