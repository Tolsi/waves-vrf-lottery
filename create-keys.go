package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha512"
	"encoding/json"
	"flag"
	"fmt"
	. "github.com/Tolsi/vrf-lottery/tools"
	"github.com/Tolsi/vrf-lottery/vrf"
	"github.com/btcsuite/btcutil/base58"
	"os"
)

type OutputKeys struct {
	Seed       string `json:"seed"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

func main() {
	//region Read params

	seed := flag.String("seed", "", "A base58 seed phrase used to generate keys, optional")
	printJson := flag.Bool("json", false, "Output JSON, not plain text")

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

	if *printJson {
		err = json.NewEncoder(os.Stdout).Encode(OutputKeys{*seed, base58.Encode(private), base58.Encode(public)})
		PrintErrorAndExit(err)
	} else {
		fmt.Printf("Seed: %s\n", *seed)
		fmt.Printf("Private key: %s\n", base58.Encode(private))
		fmt.Printf("Public key: %s\n", base58.Encode(public))
	}

	//endregion
}
