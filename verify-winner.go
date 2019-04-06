package main

import (
	"encoding/json"
	"flag"
	"fmt"
	. "github.com/Tolsi/vrf-lottery/tools"
	"github.com/Tolsi/vrf-lottery/vrf"
	"github.com/btcsuite/btcutil/base58"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//region Read params

	participantsFilename := *flag.String("participantsFile", "", "A path to file to validate. It should contains 2 lines: json array of the participants and a block signature")
	blockHeight := *flag.Uint("blockHeight", 0, "A waves block height, the signature of it will be used to validate the data")
	proofBase58 := *flag.String("proof", "", "A proof to validate the message")
	publicKeyBase58 := *flag.String("publicKey", "", "A ed25519 public key in Base58 to validate the message")
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
	if proofBase58 == "" {
		flag.Usage()
		os.Exit(1)
	}
	if publicKeyBase58 == "" {
		flag.Usage()
		os.Exit(1)
	}
	if pickN < 1 {
		flag.Usage()
		os.Exit(1)
	}

	participantsAndBlockSignature, err := ioutil.ReadFile(participantsFilename)
	PrintErrorAndExit(err)
	s := strings.Split(string(participantsAndBlockSignature), "\n")
	var participants []string
	err = json.Unmarshal([]byte(s[0]), &participants)
	PrintErrorAndExit(err)
	blockSignature := s[1]

	proofBytes := base58.Decode(proofBase58)

	pkb := vrf.PublicKey(base58.Decode(publicKeyBase58))
	PrintErrorAndExit(err)

	//endregion

	//region Create proofs

	loadedBlockSignature, err := GetBlockSignature(uint(blockHeight))
	PrintErrorAndExit(err)

	if loadedBlockSignature != blockSignature {
		fmt.Printf("Provable file contains different block signature at height %d: expected '%s', got '%s'\n", blockHeight, loadedBlockSignature, blockSignature)
		os.Exit(1)
	}

	verifyResult, vrfBytes := pkb.Verify(participantsAndBlockSignature, proofBytes)
	if !verifyResult {
		fmt.Printf("Proof verification was failed\n")
		os.Exit(1)
	}

	//endregion

	//region Result output

	winners := PickUniquePseudorandomParticipants(vrfBytes[:], int(pickN), participants)

	fmt.Printf("message: %s\n", string(participantsAndBlockSignature))
	fmt.Printf("proof (base58): %s\n", proofBase58)
	fmt.Printf("vrf bytes (base58): %s\n", base58.Encode(vrfBytes))
	fmt.Printf("winners are participants: %v\n", winners)

	//endregion
}
