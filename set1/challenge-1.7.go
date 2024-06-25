package set1

import (
	lib "cryptopals/library"
	b64 "encoding/base64"
	"fmt"
	"os"
)

// https://stackoverflow.com/questions/24072026/golang-aes-ecb-encryption

func Seven() {
	data64, err := os.ReadFile("set1/7.txt")
	lib.Check(err)
	// Don't need to close since we're just reading.

	Key := []byte("YELLOW SUBMARINE")

	data, _ := b64.StdEncoding.DecodeString(string(data64))

	decrypted := lib.AesECBDecrypt(data, Key)

	fmt.Println("Seven")
	fmt.Println(string(decrypted))

	fmt.Printf("Challenge 7\n")
	fmt.Printf("BEGIN DECODED TEXT\n")
	fmt.Printf("------------------\n")
	fmt.Printf("%s\n", decrypted)

}
