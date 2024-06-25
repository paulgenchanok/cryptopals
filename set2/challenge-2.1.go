package set2

import (
	lib "cryptopals/library"
	"fmt"
)

func Nine() {

	data := []byte("YELLOW SUBMARINE")

	padded := lib.PKCS7Pad(data, 20)

	fmt.Printf("Challenge 9\n")

	fmt.Println(padded)

	unPadded := lib.PKCS7UnPad(padded)

	fmt.Println(unPadded)

	fmt.Printf("\n")

}
