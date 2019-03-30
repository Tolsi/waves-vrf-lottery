package main

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/keytransparency/core/crypto/vrf/p256"
	_ "github.com/tolsi/vrf-lottery/tools"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
)

// openssl ecparam -name prime256v1 -genkey -out p256-key.pem -noout
// openssl ec -in p256-key.pem -pubout -out p256-pubkey.pem
func main() {
	m, err := ioutil.ReadFile(os.Args[1])
	PanicIfError(err)
	skb, err := ioutil.ReadFile(os.Args[2])

	PanicIfError(err)
	modulo, err := strconv.ParseInt(os.Args[3], 10, 64)
	PanicIfError(err)
	signer, err := p256.NewVRFSignerFromPEM(skb)
	PanicIfError(err)
	index1b, proof := signer.Evaluate(m)
	PanicIfError(err)

	////region Verify proof
	//index2b, err := verifier.ProofToHash(m, proof)
	//if err != nil {
	//	fmt.Print("Evaluated proof isn't valid\n")
	//}
	////endregion

	//region Result output
	index1 := new(big.Int)
	index1.SetBytes(index1b[:])
	//index2 := new(big.Int)
	//index2.SetBytes(index2b[:])
	//if bytes.Compare(index1b[:], index2b[:]) != 0 {
	//	fmt.Print("Got different indexes after evaluate proof\n")
	//}

	moduloResult := new(big.Int)
	moduloBigint := new(big.Int)
	moduloBigint.SetInt64(modulo)
	moduloResult = moduloResult.Mod(index1, moduloBigint)

	fmt.Printf("message: %s\n", string(m))
	fmt.Printf("proof: %s\n", base58.Encode(proof))
	fmt.Printf("index: %d\n", index1)
	fmt.Printf("modulo: %d\n", moduloResult)
}
