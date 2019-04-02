package main

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/coniks-sys/coniks-go/crypto/vrf"
	"github.com/tolsi/vrf-lottery/tools"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
)

// you can create the keys using:
// openssl ecparam -name prime256v1 -genkey -out p256-key.pem -noout
// openssl ec -in p256-key.pem -pubout -out p256-pubkey.pem
func main() {
	//region Read params

	m, err := ioutil.ReadFile(os.Args[1])
	tools.PrintErrorAndExit(err)
	skb := vrf.PrivateKey(base58.Decode(os.Args[2]))

	tools.PrintErrorAndExit(err)
	modulo, err := strconv.ParseInt(os.Args[3], 10, 64)
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

	moduloResult := new(big.Int)
	moduloBigint := new(big.Int)
	moduloBigint.SetInt64(modulo)
	moduloResult = moduloResult.Mod(vrfNumber, moduloBigint)

	fmt.Printf("message: '%s'\n", string(m))
	fmt.Printf("proof (base58): '%s'\n", base58.Encode(proof))
	fmt.Printf("vrf bytes: %s\n", vrfBytes)
	fmt.Printf("vrf as number: %d\n", vrfNumber)
	fmt.Printf("modulo: %d\n", moduloResult)

	//endregion
}
