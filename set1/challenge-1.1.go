package set1

import (
	lib "cryptopals/library"
	b64 "encoding/base64"
	"fmt"
)

func One() {

	data := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	data64 := b64.StdEncoding.EncodeToString(lib.HexToBytes(data))

	fmt.Printf("Challenge 1\n")
	fmt.Printf("%s\n", data64)
	fmt.Printf("\n")

}
