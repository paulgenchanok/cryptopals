package set2

import (
	lib "cryptopals/library"
	b64 "encoding/base64"
	"fmt"
	"os"
)

func Ten() {

	data64, err := os.ReadFile("set2/10.txt")
	lib.Check(err)

	data, _ := b64.StdEncoding.DecodeString(string(data64))

	key := []byte("YELLOW SUBMARINE")

	InitializationVector := make([]byte, 16)

	// All zero IV
	//
	for i := range InitializationVector {
		InitializationVector[i] = '\x00'
	}

	decrypted := lib.AesCBCDecrypt(data, key, InitializationVector)

	fmt.Printf("Challenge 10\n")

	fmt.Println(string(decrypted))

}
