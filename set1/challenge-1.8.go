package set1

import (
	"bufio"
	lib "cryptopals/library"
	"encoding/hex"
	"fmt"
	"os"
)

func Eight() {

	file, err := os.Open("set1/8.txt")
	lib.Check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Assume a 16 byte keysize?
	// Score is going through the bytes, block by block
	// For each block, check if it has been seen before. The one with the most recurrences is most likely cbc

	ks := 16

	bestScore := 0
	bestData := []byte{}
	bestLineNumber := 0
	currLineNumber := 0

	for scanner.Scan() {

		lineData := lib.HexToBytes(scanner.Text()) // In Bytes now
		currLineNumber++

		_, hits := lib.DetectECB(lineData, ks)

		if hits > bestScore {
			bestScore = hits
			bestData = lineData
			bestLineNumber = currLineNumber

		}

	}

	fmt.Printf("Challenge 8\n")
	fmt.Printf("Best Line Number: %d. Score: %d\n", bestLineNumber, bestScore)
	fmt.Printf("Best Data: %s\n", hex.EncodeToString(bestData))

}
