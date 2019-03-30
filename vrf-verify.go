package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"

	"github.com/btcsuite/btcutil/base58"
	"github.com/google/keytransparency/core/crypto/vrf/p256"
	_ "github.com/tolsi/vrf-lottery/tools"
)

func main() {
	//region Read params

	m, err := ioutil.ReadFile(os.Args[1])
	PanicIfError(err)

	proofString := os.Args[2]
	proof := base58.Decode(proofString)

	pkb, err := ioutil.ReadFile(os.Args[3])
	PanicIfError(err)
	modulo, err := strconv.ParseInt(os.Args[4], 10, 64)
	PanicIfError(err)

	//endregion

	//region Verify proof

	verifier, err := p256.NewVRFVerifierFromPEM(pkb)
	PanicIfError(err)
	index2b, err := verifier.ProofToHash(m, proof)
	PanicIfError(err)

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
