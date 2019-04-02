package main

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/coniks-sys/coniks-go/crypto/vrf"
	"github.com/tolsi/vrf-lottery/tools"
	"io/ioutil"
	"math/big"
	"os"
)

// you can create the keys using create-keys.go
func main() {
	//region Read params

	m, err := ioutil.ReadFile(os.Args[1])
	tools.PrintErrorAndExit(err)
	skb := vrf.PrivateKey(base58.Decode(os.Args[2]))

	tools.PrintErrorAndExit(err)

	//endregion

	//region Create proofs

	tools.PrintErrorAndExit(err)
	vrfBytes, proof := skb.Prove(m)
	tools.PrintErrorAndExit(err)

	//endregion

	//region Result output

	vrfNumber := new(big.Int)
	vrfNumber.SetBytes(vrfBytes[:])

	fmt.Printf("message: '%s'\n", string(m))
	fmt.Printf("proof (base58): '%s'\n", base58.Encode(proof))
	fmt.Printf("vrf bytes: %s\n", vrfBytes)

	//endregion
}
