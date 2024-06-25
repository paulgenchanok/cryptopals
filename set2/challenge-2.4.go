package set2

import (
	lib "cryptopals/library"
	b64 "encoding/base64"
	"fmt"
)

func Twelve() {

	secret, _ := b64.StdEncoding.DecodeString("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")

	ptSecret := lib.AesECBOracleDecrypt(secret)

	fmt.Println(string(ptSecret))

}
