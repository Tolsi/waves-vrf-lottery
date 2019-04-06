package main

import (
	"bytes"
	"crypto/sha512"
	"fmt"
	. "github.com/Tolsi/vrf-lottery/tools"
	"github.com/Tolsi/vrf-lottery/vrf"
	"github.com/btcsuite/btcutil/base58"
	"os"
)

func main() {
	hasher := sha512.New()
	hasher.Write([]byte(os.Args[1]))
	private, err := vrf.GenerateKey(bytes.NewReader(hasher.Sum(nil)))
	PrintErrorAndExit(err)
	public, _ := private.Public()
	fmt.Printf("Private key: %s\n", base58.Encode(private))
	fmt.Printf("Public key: %s\n", base58.Encode(public))
}
