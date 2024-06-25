package set1

import (
	"bufio"
	lib "cryptopals/library"
	"fmt"
	"os"
)

func Four() {

	file, err := os.Open("set1/4.txt")
	lib.Check(err)

	defer file.Close()

	fmt.Printf("Challenge 4\n")

	scanner := bufio.NewScanner(file)

	// keyFinal := 0
	greatestScore := 0
	var dataFinal []byte
	line := 0

	for scanner.Scan() {

		data := lib.HexToBytes(scanner.Text())

		_, dxor, score := lib.XorDecryptSingleByte(data, lib.FreqVowels)

		if score > greatestScore {
			greatestScore = score
			dataFinal = dxor
			copy(dataFinal, dxor)
			fmt.Printf("%d: %s\n", line, dataFinal)
		}

		line++

	}

	fmt.Printf("%s\n", dataFinal)

}
