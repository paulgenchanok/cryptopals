package set2

import (
	lib "cryptopals/library"
	"fmt"
	"os"
)

func Eleven() {

	keyLen := 16
	key := lib.RandBytes(keyLen)

	plainText, err := os.ReadFile("lyrics.txt")
	lib.Check(err)

	pre := lib.RandPad(5, 10)
	post := lib.RandPad(5, 10)

	plainTextBuffd := append(pre, []byte(plainText)...)
	plainTextBuffd = append(plainTextBuffd, post...)

	coinflip := lib.RandInt(0, 2)

	fmt.Printf("Challenge 11\n")

	var cipherText []byte

	if coinflip == 0 {
		// Encrypt with AES-128 EBC

		fmt.Printf("ECB Encrypting\n")
		cipherText = lib.AesECBEncrypt(plainTextBuffd, key)

	} else {
		// Encrypt with AES-128 CBC. Need random IV.
		fmt.Printf("CBC Encrypting\n")
		randIV := lib.RandBytes(keyLen)

		cipherText = lib.AesCBCEncrypt(plainTextBuffd, key, randIV)

	}

	isEBC, _ := lib.DetectECB(cipherText, keyLen)

	if isEBC {
		fmt.Printf("ECB Detected\n")
	} else {
		fmt.Printf("CBC Detected\n")
	}

	fmt.Printf("\n")

}
