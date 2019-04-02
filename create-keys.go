package main

import (
	"bytes"
	"crypto/sha512"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/coniks-sys/coniks-go/crypto/vrf"
	"os"
)

func main() {
	hasher := sha512.New()
	hasher.Write([]byte(os.Args[1]))
	private, _ := vrf.GenerateKey(bytes.NewReader(hasher.Sum(nil)))
	public, _ := private.Public()
	fmt.Printf("Private key: %s\n", base58.Encode(private))
	fmt.Printf("Public key: %s\n", base58.Encode(public))
}
