package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/Tolsi/vrf-lottery/tools"
	"github.com/Tolsi/vrf-lottery/vrf"
	"github.com/btcsuite/btcutil/base58"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	//region Read params

	participantsFile, err := ioutil.ReadFile(os.Args[1])
	PrintErrorAndExit(err)
	var participants []string
	err = json.Unmarshal(participantsFile, &participants)
	PrintErrorAndExit(err)

	blockHeight, err := strconv.ParseInt(os.Args[2], 10, 64)
	PrintErrorAndExit(err)

	skb := vrf.PrivateKey(base58.Decode(os.Args[3]))
	PrintErrorAndExit(err)

	pickN, err := strconv.ParseInt(os.Args[4], 10, 64)
	PrintErrorAndExit(err)

	//endregion

	//region Create proofs

	blockSignature, err := GetBlockSignature(uint(blockHeight))
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
