package main

import (
	"github.com/hashicorp/vault/shamir"
	"crypto/aes"
	"fmt"
	"github.com/btcsuitereleases/btcutil/base58"
	"log"
	"crypto/rand"
)

func main() {

	parts := 5
	threshold := 3
	// Generate a 256bit key
	masterKey := make([]byte, 2*aes.BlockSize)
	_, err :=rand.Read(masterKey)
	if err != nil {
		log.Fatalf("Generate key failed\n%s", err)

	}
	fmt.Printf("Master key\t\t\t%s\n", base58.Encode(masterKey))

	sharedKeys, err := shamir.Split(masterKey, parts, threshold)
	if err != nil {
		log.Fatalf("Split master key failed\n%s", err)
	}

	for i := range sharedKeys {
		fmt.Printf("Shared key %d\t\t%s\n", i, base58.Encode(sharedKeys[i]))
	}

	shares := [][]byte{sharedKeys[4], sharedKeys[0], sharedKeys[2]}
	reconstructed, err := shamir.Combine(shares)
	if err != nil {
		log.Fatalf("Reconstruct shared parts failed\n%s", err)
	}
	fmt.Printf("Reconstructed key\t%s", base58.Encode(reconstructed))
}
