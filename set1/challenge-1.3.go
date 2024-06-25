package set1

import (
	lib "cryptopals/library"
	"fmt"
)

// var english_freqs = [26]float32{
// 	0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, // A-G
// 	0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749, // H-N
// 	0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758, // O-U
// 	0.00978, 0.02360, 0.00150, 0.01974, 0.00074} // V-Z

// var freq = []byte{' ', 'e', 't', 'a', 'o', 'i', 'n', 's', 'h', 'r', 'd', 'l', 'u'}

// func getChi2(decodedstr []byte) float32 {

// 	count := [52]int{0}
// 	ignored := 0

// 	for ind := 0; ind < len(decodedstr); ind++ {
// 		char := decodedstr[ind] // we laredy have ascii

// 		// uppercase A-Z
// 		if char >= 65 && char <= 90 {
// 			count[char-65]++

// 			// lowercase a-z
// 		} else if char >= 97 && char <= 122 {
// 			count[char-97]++

// 			// numbers and punctuation
// 		} else if char >= 32 && char <= 126 {
// 			ignored++

// 			// TAB, CR, LF
// 		} else if char == 9 || char == 10 || char == 13 {
// 			ignored++

// 		} else {
// 			return -1 // Error.
// 		}
// 	}

// 	// Compute Chi2
// 	//
// 	var chi2 float32 = 0
// 	length := len(decodedstr) - ignored

// 	for ind := 0; ind < 26; ind++ {

// 		observed := float32(count[ind])
// 		expected := float32(length) * english_freqs[ind]
// 		difference := observed - expected
// 		chi2 += difference

// 	}

// 	return chi2

// }

// type threeData struct {
// 	char     byte
// 	score    float32
// 	data_dec []byte
// }

// // A little more customizable if you want to rank order everything
// func xor_decode_complicated(dstr []byte) threeData {

// 	valid_ascii := make([]threeData, 0, 255)

// 	for char := 0; char <= 255; char++ {

// 		data_xor := lib.XorSingleByte(dstr, byte(char))
// 		chi2 := getChi2(data_xor)

// 		if chi2 != -1 {
// 			valid_ascii = append(valid_ascii, threeData{char: byte(char), score: chi2, data_dec: data_xor})
// 		}

// 	}

// 	slices.SortFunc(valid_ascii, func(a, b threeData) int {

// 		// Sort so the best score is at the top ... DESCENDING
// 		return cmp.Compare(b.score, a.score)

// 	})

// 	return valid_ascii[0]

// }

func Three() {

	data := lib.HexToBytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	var _, data_out, _ = lib.XorDecryptSingleByte(data, lib.FreqVowels)

	fmt.Printf("Challenge 3\n")
	fmt.Printf("%s\n", data_out)
	fmt.Printf("\n")

}
