package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Tolsi/vrf-lottery/curve"
	. "github.com/Tolsi/vrf-lottery/tools"
	"github.com/btcsuite/btcutil/base58"
	"os"
)

type OutputKeys struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

func main() {
	//region Read params

	printJson := flag.Bool("json", false, "Output JSON, not plain text")

	flag.Parse()

	//endregion

	//region Create keys

	private, public, err := curve.GenerateKeypair()
	PrintErrorAndExit(err)

	//endregion

	//region Result output

	if *printJson {
		err = json.NewEncoder(os.Stdout).Encode(OutputKeys{base58.Encode(private), base58.Encode(public)})
		PrintErrorAndExit(err)
	} else {
		fmt.Printf("Private key: %s\n", base58.Encode(private))
		fmt.Printf("Public key: %s\n", base58.Encode(public))
	}

	//endregion
}
