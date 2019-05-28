package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	plaintext, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}

	key := []byte(os.Args[1])
	if dat, err := ioutil.ReadFile(string(key)); err == nil { // if key is a path, read the key from this file
		key = dat
	}

	if len(plaintext) != len(key) {
		log.Fatalf("plaintext and key should have the same length (%d != %d)", len(plaintext), len(key))
	}

	// one-time-pad encryption/decryption
	ciphertext := []byte{}
	for i := 0; i < len(plaintext); i++ {
		ciphertext = append(ciphertext, plaintext[i]^key[i]) // xor
	}

	fmt.Print(string(ciphertext))
}
