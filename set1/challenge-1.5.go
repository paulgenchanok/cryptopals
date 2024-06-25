package set1

import (
	lib "cryptopals/library"
	"encoding/hex"
	"fmt"
)

func Five() {
	data := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	data_b := []byte(data) // Now we have the byte array

	data_xor := lib.XorRepeatingKey(data_b, []byte("ICE"))

	data_hex := make([]byte, hex.EncodedLen(len(data_xor)))
	hex.Encode(data_hex, data_xor)

	fmt.Printf("Challenge 5\n")
	fmt.Printf("%s\n", string(data_hex))
	fmt.Printf("\n")

}
