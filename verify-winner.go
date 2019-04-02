package main

import (
	"encoding/json"
	"fmt"
	"github.com/Tolsi/vrf-lottery/tools"
	"github.com/btcsuite/btcutil/base58"
	"github.com/coniks-sys/coniks-go/crypto/vrf"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	//region Read params

	participantsAndBlockSignature, err := ioutil.ReadFile(os.Args[1])
	tools.PrintErrorAndExit(err)
	s := strings.Split(string(participantsAndBlockSignature), "\n")
	var participants []string
	err = json.Unmarshal([]byte(s[0]), &participants)
	tools.PrintErrorAndExit(err)
	blockSignature := s[1]
	tools.PrintErrorAndExit(err)

	blockHeight, err := strconv.ParseInt(os.Args[2], 10, 64)
	tools.PrintErrorAndExit(err)

	vrfString := os.Args[3]
	vrfBytes := base58.Decode(vrfString)

	proofString := os.Args[4]
	proofBytes := base58.Decode(proofString)

	pkb := vrf.PublicKey(base58.Decode(os.Args[5]))
	tools.PrintErrorAndExit(err)

	//endregion

	//region Create proofs

	loadedBlockSignature, err := tools.GetBlockSignature(uint(blockHeight))
	tools.PrintErrorAndExit(err)

	if loadedBlockSignature != blockSignature {
		fmt.Printf("Provable file contains different block signature at height %d: expected '%s', got '%s'\n", blockHeight, loadedBlockSignature, blockSignature)
		os.Exit(1)
	}

	verifyResult := pkb.Verify(participantsAndBlockSignature, vrfBytes, proofBytes)
	if !verifyResult {
		fmt.Printf("Proof verification was failed")
		os.Exit(1)
	}

	//endregion

	//region Result output

	vrfNumber := new(big.Int)
	vrfNumber.SetBytes(vrfBytes[:])

	moduloBigint := big.NewInt(int64(len(participants)))
	moduloResult := new(big.Int).Mod(vrfNumber, moduloBigint)

	winner := participants[moduloResult.Int64()]

	fmt.Printf("message: %s\n", string(participantsAndBlockSignature))
	fmt.Printf("proof (base58): %s\n", proofString)
	fmt.Printf("vrf bytes (base58): %s\n", base58.Encode(vrfBytes))
	fmt.Printf("vrf as number: %d\n", vrfNumber)
	fmt.Printf("modulo: %d\n", moduloResult)
	fmt.Printf("winner is participant #%d: %s\n", moduloResult, winner)

	//endregion
}
