package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	. "github.com/Tolsi/vrf-lottery/tools"
	"github.com/Tolsi/vrf-lottery/vrf"
	"github.com/btcsuite/btcutil/base58"
	"io/ioutil"
	"os"
)

func main() {
	//region Read params

	participantsFilename := *flag.String("participantsFile", "", "A path to file with participants. It should contains json array of strings")
	blockHeight := *flag.Uint("blockHeight", 0, "A waves block height, the signature of it will be used in provable message")
	privateKeyBase58 := *flag.String("privateKey", "", "A ed25519 private key in Base58 to prove the message")
	pickN := *flag.Uint("pickN", 1, "The number of winners to pick, it should be >= 1")

	flag.Parse()

	if participantsFilename == "" {
		flag.Usage()
		os.Exit(1)
	}
	if blockHeight == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if privateKeyBase58 == "" {
		flag.Usage()
		os.Exit(1)
	}
	if pickN < 1 {
		flag.Usage()
		os.Exit(1)
	}

	participantsFile, err := ioutil.ReadFile(participantsFilename)
	PrintErrorAndExit(err)
	var participants []string
	err = json.Unmarshal(participantsFile, &participants)
	PrintErrorAndExit(err)
	skb := vrf.PrivateKey(base58.Decode(privateKeyBase58))

	//endregion

	//region Create proofs

	blockSignature, err := GetBlockSignature(blockHeight)
	PrintErrorAndExit(err)

	participantsJson, err := json.Marshal(participants)
	PrintErrorAndExit(err)
	provableMessage := append(participantsJson, []byte("\n"+blockSignature)...)
	fileName := fmt.Sprintf("participants_and_%d_block_signature.txt", blockHeight)
	err = ioutil.WriteFile(fileName, provableMessage, 0644)
	fmt.Printf("Provable lottery data was saved to file '%s'\n", fileName)
	PrintErrorAndExit(err)

	vrfBytes, proof := skb.Prove(provableMessage)
	pk, _ := skb.Public()
	verifyResult, vrfBytes2 := pk.Verify(provableMessage, proof)
	if !verifyResult || bytes.Compare(vrfBytes, vrfBytes2) != 0 {
		fmt.Printf("Proof verification was failed")
		os.Exit(1)
	}
	PrintErrorAndExit(err)

	//endregion

	//region Result output

	winners := PickUniquePseudorandomParticipants(vrfBytes[:], int(pickN), participants)

	fmt.Printf("message: %s\n", string(provableMessage))
	fmt.Printf("proof (base58): %s\n", base58.Encode(proof))
	fmt.Printf("vrf bytes (base58): %s\n", base58.Encode(vrfBytes))
	fmt.Printf("winners are participants: %v\n", winners)

	//endregion
}
