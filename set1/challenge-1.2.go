package set1

import (
	lib "cryptopals/library"
	"encoding/hex"
	"fmt"
)

func Two() {

	data1_str := "1c0111001f010100061a024b53535009181c"
	data2_str := "686974207468652062756c6c277320657965"

	data1_dec := lib.HexToBytes(data1_str)
	data2_dec := lib.HexToBytes(data2_str)

	data_xor := lib.XorByteSlices(data1_dec, data2_dec)

	fmt.Printf("Challenge 2\n")
	fmt.Println(hex.EncodeToString(data_xor))
	fmt.Printf("\n")

}
