package set1

import (
	lib "cryptopals/library"
	b64 "encoding/base64"
	"fmt"
	"os"
)

// Thanks to
// https://www.ramitmittal.com/blog/golang-hamming-distances/

func Six() {

	data64, err := os.ReadFile("set1/6.txt")
	lib.Check(err)
	// Don't need to close since we're just reading.

	data, _ := b64.StdEncoding.DecodeString(string(data64))

	KS, key := lib.XorDecryptRepeating(data, lib.FreqVowels)

	fmt.Printf("Challenge 6\n")
	fmt.Printf("Key size: %d\n", KS)
	fmt.Printf("Key: \"%s\"\n", key)

	fmt.Printf("BEGIN DECODED TEXT\n")
	fmt.Printf("------------------\n")
	fmt.Printf("%s\n", lib.XorRepeatingKey(data, key))

}
