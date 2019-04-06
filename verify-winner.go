package main

import (
	"encoding/json"
	"fmt"
	. "github.com/Tolsi/vrf-lottery/tools"
	"github.com/Tolsi/vrf-lottery/vrf"
	"github.com/btcsuite/btcutil/base58"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	//region Read params

	participantsAndBlockSignature, err := ioutil.ReadFile(os.Args[1])
	PrintErrorAndExit(err)
	s := strings.Split(string(participantsAndBlockSignature), "\n")
	var participants []string
	err = json.Unmarshal([]byte(s[0]), &participants)
	PrintErrorAndExit(err)
	blockSignature := s[1]
	PrintErrorAndExit(err)

	blockHeight, err := strconv.ParseInt(os.Args[2], 10, 64)
	PrintErrorAndExit(err)

	proofString := os.Args[3]
	proofBytes := base58.Decode(proofString)

	pkb := vrf.PublicKey(base58.Decode(os.Args[4]))
	PrintErrorAndExit(err)

	pickN, err := strconv.ParseInt(os.Args[5], 10, 64)
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
	fmt.Printf("proof (base58): %s\n", proofString)
	fmt.Printf("vrf bytes (base58): %s\n", base58.Encode(vrfBytes))
	fmt.Printf("winners are participants: %v\n", winners)

	//endregion
}
