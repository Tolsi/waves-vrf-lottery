package main

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/keytransparency/core/crypto/vrf/p256"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"tools"
)

func main() {
	//region Read params

	m, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	proofString := os.Args[2]
	proof := base58.Decode(proofString)

	pkb, err := ioutil.ReadFile(os.Args[3])
	if err != nil {
		panic(err)
	}
	modulo, err := strconv.ParseInt(os.Args[4], 10, 64)
	panicIfError(err)

	//endregion

	//region Verify proof

	verifier, err := p256.NewVRFVerifierFromPEM(pkb)
	if err != nil {
		panic(err)
	}
	index2b, err := verifier.ProofToHash(m, proof)
	if err != nil {
		fmt.Print("Evaluated proof isn't valid\n")
		os.Exit(1)
	}

	//endregion

	//region Result output

	//index1 := new(big.Int)
	//index1.SetBytes(index1b[:])
	index2 := new(big.Int)
	index2.SetBytes(index2b[:])
	//if bytes.Compare(index1b[:], index2b[:]) != 0 {
	//	fmt.Print("Got different indexes after evaluate proof\n")

	moduloResult := new(big.Int)
	moduloBigint := new(big.Int)
	moduloBigint.SetInt64(modulo)
	moduloResult = moduloResult.Mod(index2, moduloBigint)

	fmt.Printf("message: %s\n", string(m))
	fmt.Printf("proof: %s\n", proofString)
	fmt.Printf("index: %d\n", index2)
	fmt.Printf("modulo: %d\n", moduloResult)

	//endregion
}
