package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/btcsuite/btcutil/base58"
	"github.com/coniks-sys/coniks-go/crypto/vrf"
	"github.com/tolsi/vrf-lottery/tools"
)

func main() {
	//region Read params

	m, err := ioutil.ReadFile(os.Args[1])
	tools.PrintErrorAndExit(err)

	vrfString := os.Args[2]
	vrfBytes := base58.Decode(vrfString)

	proofString := os.Args[3]
	proofBytes := base58.Decode(proofString)

	pkb := vrf.PublicKey(base58.Decode(os.Args[4]))
	tools.PrintErrorAndExit(err)

	//endregion

	//region Verify proofBytes

	verifyResult := pkb.Verify(m, vrfBytes, proofBytes)
	if !verifyResult {
		fmt.Printf("Proof verification was failed")
	}

	//endregion

	//region Result output

	fmt.Printf("message: '%s'\n", string(m))
	fmt.Printf("proof (base58): '%s'\n", proofString)
	fmt.Printf("vrf bytes: %s\n", base58.Encode(vrfBytes))

	//endregion
}
