package library

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"slices"
	"strconv"
	"strings"
)

// Percent commonality of english letters
var english_freqs = [26]float32{
	0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, // A-G
	0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749, // H-N
	0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758, // O-U
	0.00978, 0.02360, 0.00150, 0.01974, 0.00074} // V-Z

// ETAOIN SHRDL Vowel frequency array for comparison
var FreqVowels = []byte{' ', 'e', 't', 'a', 'o', 'i', 'n', 's', 'h', 'r', 'd', 'l', 'u'}

// Numeric frequency for comparison
var FreqNumeric = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func HexToBytes(hexstr string) []byte {
	hexstr_bytes := []byte(hexstr) // Operate on bytes not encoding strings

	dst := make([]byte, hex.DecodedLen(len(hexstr_bytes))) // Go's allocator.
	n, err := hex.Decode(dst, hexstr_bytes)                // Decode the bytes?

	if err != nil {
		log.Fatal(err)
	}

	return dst[:n]
}

func XorByteSlices(b1 []byte, b2 []byte) []byte {

	b1_len := len(b1)

	if b1_len != len(b2) {
		panic("XorByteSlices: different sizes")
	}

	b3 := make([]byte, b1_len)

	for ind := range b1_len {

		b3[ind] = b1[ind] ^ b2[ind]

	}

	return b3
}

// XORs all bytes in data by one byte and returns the result
func XorSingleByte(data []byte, char byte) []byte {

	data_len := len(data)
	data_xor := make([]byte, data_len)

	for ind := 0; ind < data_len; ind++ {
		data_xor[ind] = data[ind] ^ char
	}

	return data_xor

}

func xorAssignScore(dstr []byte, freq []byte) int {
	score := 0

	for ind := 0; ind < len(dstr); ind++ {
		if slices.Contains(freq, dstr[ind]) {
			score += 1
		}

		if dstr[ind] <= 31 || dstr[ind] >= 127 {
			// Non printable character reached. Assuming the output is ASCII printable
			//
			score = 0
			break
		}
	}

	return score

}

// Decrypt data assuming it's a single byte XOR encrypted ASCII string
// Freq is the frequency array. What to SCORE against
func XorDecryptSingleByte(data []byte, freq []byte) (int, []byte, int) {

	data_len := len(data)

	key := -1                          // The key
	data_out := make([]byte, data_len) // Decoded data

	greatest_score := 0

	for char := 0; char <= 255; char++ {

		data_xor := XorSingleByte(data, byte(char))
		score := xorAssignScore(data_xor, freq)

		if score > greatest_score {
			greatest_score = score
			key = char
			data_out = data_xor
		}

	}

	return key, data_out, greatest_score

}

func checkABA(data []byte) bool {
	// Okay. We have byte. That might be the ASCII reprsentation of numbers.
	dataLen := len(data)

	var aba [9]int

	for ind := 0; ind < dataLen; ind++ {
		abat, err := strconv.Atoi(string(data[ind]))

		if err != nil {
			return false
		}

		aba[ind] = abat
	}

	check := 3*(aba[0]+aba[3]+aba[6]) + 7*(aba[1]+aba[4]+aba[7]) + aba[2] + aba[5] + aba[8]

	check = check % 10

	return (check == 0)
}

// Decrypt an ABA routing ... I think ...
func XorDecryptSingleByteABA(data []byte) {

	for char := 0; char <= 255; char++ {

		data_xor := XorSingleByte(data, byte(char))
		abaCheck := checkABA(data_xor)

		if abaCheck {
			fmt.Printf("key: %d. data: %s\n", char, data_xor)
		}

	}

}

// XOR encrypts data with the provided keyh
func XorRepeatingKey(data []byte, key []byte) []byte {

	key_len := len(key)
	data_len := len(data)

	if key_len > data_len {
		panic("XorRepatingKey: key longer than data")
	}

	// Use modulo for the good stuff here. Easy enough.

	data_out := make([]byte, data_len)

	for ind := 0; ind < data_len; ind++ {

		data_out[ind] = data[ind] ^ key[ind%key_len]

	}

	return data_out

}

// Begin break repeating XOR functions

var oneSetBitByteSlice = []byte{
	0b00000001,
	0b00000010,
	0b00000100,
	0b00001000,
	0b00010000,
	0b00100000,
	0b01000000,
	0b10000000,
}

func countSetBits(b1 []byte) int {

	hammingDist := 0

	// b is a byte in b1
	for _, b := range b1 {
		for _, oneSetBitByte := range oneSetBitByteSlice {
			if (b & oneSetBitByte) > 0 {
				hammingDist += 1
			}
		}
	}

	return hammingDist

}

// Assuming data is repeating XOR encrypted, find the most likely key size
func xorRepeatingFindKeySize(data []byte, ksStart int, ksEnd int) int {

	data_len := len(data)

	hbest := float32(data_len) // Initial upper bound for hnorm
	KS := 0

	for ks := ksStart; ks <= ksEnd; ks++ {

		hammingDist := 0
		num_iters := data_len / ks

		for ind := 0; ind < num_iters; ind++ {

			// Take two blocks next to each other all the way across
			//
			start := ind * ks
			mid := start + ks
			end := mid + ks

			blk1 := data[start:mid]
			blk2 := data[mid:end]

			x1 := XorByteSlices(blk1, blk2)
			hammingDist += countSetBits(x1)

		}

		hbar := float32(hammingDist) / float32(1+num_iters)
		hnorm := hbar / float32(ks)

		if hnorm < hbest {

			hbest = hnorm
			KS = ks

		}

	}

	return KS

}

// Decrypt data assuming data is repeating XOR encyrpted
func XorDecryptRepeating(data []byte, freq []byte) (keysize int, keyBytes []byte) {

	KS := xorRepeatingFindKeySize(data, 2, 40)

	var key = make([]byte, KS)

	transpose := make([][]byte, KS)
	for i := range transpose {
		transpose[i] = make([]byte, 0)
	}

	for ind := 0; ind < len(data); ind++ {
		transpose[ind%KS] = append(transpose[ind%KS], data[ind])
	}

	for ind := 0; ind < KS; ind++ {
		// Solve the single character XORs
		//
		k, _, _ := XorDecryptSingleByte(transpose[ind], freq)
		key[ind] = byte(k)
	}

	return KS, key

}

func AesECBDecrypt(data []byte, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)

	dataLen := len(data)

	decrypted := make([]byte, dataLen)
	size := len(key)

	// They work in pairs. This inits bs to 0, be to size
	for bs, be := 0, size; bs < dataLen; bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	decrypted = PKCS7UnPad(decrypted)

	return decrypted
}

func AesECBEncrypt(plainText []byte, key []byte) []byte {

	keyLen := len(key)
	plainText = PKCS7Pad(plainText, keyLen)

	PtLen := len(plainText)

	cipher, _ := aes.NewCipher(key)

	cipherText := make([]byte, PtLen)

	for bs, be := 0, keyLen; bs < PtLen; bs, be = bs+keyLen, be+keyLen {
		cipher.Encrypt(cipherText[bs:be], plainText[bs:be])
	}

	return cipherText

}

// TODO: Simplify this. Don't need a full map. And limited to 16 byte keys
func DetectECB(cipherText []byte, keyLen int) (bool, int) {

	CtLen := len(cipherText)

	uniqueBlocks := make(map[[16]byte]int)

	hits := 0
	isEBC := false

	for bs, be := 0, keyLen; bs < CtLen; bs, be = bs+keyLen, be+keyLen {

		blk := [16]byte(cipherText[bs:be])

		_, prs := uniqueBlocks[blk]

		// Check if it's present or not. If it is present ... ..
		if prs {
			hits++
			uniqueBlocks[blk]++
			isEBC = true // This is our criterica for true ...
		} else {
			// Add it
			uniqueBlocks[blk] = 1
		}

	}

	return isEBC, hits

}

// END SET 1 FUNCTIONS

// BEGIN SET 2 FUNCTIONS

func PKCS7Pad(data []byte, ks int) []byte {

	dataLen := len(data)
	paddedLen := ((dataLen / ks) + 1) * ks

	padded := make([]byte, paddedLen)

	copy(padded, data)

	pad := paddedLen - dataLen

	for ind := dataLen; ind < paddedLen; ind++ {
		padded[ind] = byte(pad)
	}

	return padded

}

func PKCS7UnPad(data []byte) []byte {

	dataLen := len(data)
	newLen := dataLen

	padChar := data[dataLen-1]

	for data[newLen-1] == padChar {
		newLen--
	}

	unPadded := make([]byte, newLen)
	copy(unPadded, data[:newLen])

	return unPadded

}

func AesCBCDecrypt(data []byte, key []byte, InitializationVector []byte) []byte {

	keyLen := len(key)
	dataLen := len(data)
	cipher, _ := aes.NewCipher(key)

	//TODO: Input error checking. Mature libraries do this.

	decrypted := make([]byte, dataLen)
	xordBlock := make([]byte, keyLen)
	prevBlock := InitializationVector

	for bs, be := 0, keyLen; bs < dataLen; bs, be = bs+keyLen, be+keyLen {

		// Decrypt to get the PT XOR'd with the IV or Previous Block
		cipher.Decrypt(xordBlock, data[bs:be])

		// (un)XOR the result with prevBlock and save the PT
		copy(decrypted[bs:be], XorByteSlices(xordBlock, prevBlock))

		// Set the CT prevBlock to the current CT block
		prevBlock = data[bs:be]

	}

	decrypted = PKCS7UnPad(decrypted)

	return decrypted

}

func AesCBCEncrypt(data []byte, key []byte, InitializationVector []byte) []byte {

	// Pad the data before further processing
	keyLen := len(key)
	data = PKCS7Pad(data, keyLen)

	dataLen := len(data)
	cipher, _ := aes.NewCipher(key)

	encrypted := make([]byte, dataLen)
	prevBlock := InitializationVector

	for bs, be := 0, keyLen; bs < dataLen; bs, be = bs+keyLen, be+keyLen {

		// Encrypt to get the PT XOR'd with the IV or Previous Block
		cipher.Encrypt(encrypted[bs:be], XorByteSlices(data[bs:be], prevBlock))

		// Set the CT prevBlock to the current CT block
		prevBlock = encrypted[bs:be]

	}

	return encrypted

}

// In range [min, max)
func RandInt(min int, max int) int {
	diff := max - min

	rbig, _ := rand.Int(rand.Reader, big.NewInt(int64(diff)+1))
	rint, _ := strconv.Atoi(rbig.String()) // KEEP IT SIMPLE

	return (min + rint)
}

func RandBytes(len int) []byte {

	rBytes := make([]byte, len)

	_, err := rand.Read(rBytes)

	if err != nil {
		panic("Eleven creating random key")
	}

	return rBytes
}

func RandPad(min int, max int) []byte {

	padLen := RandInt(min, max)

	return RandBytes(padLen)

}

// AES ECB Oracle Attack

var OracleKey = RandBytes(16)

func AesECBOracleEncrypt(myData []byte, secret []byte) []byte {

	paddedPlainText := slices.Concat(myData, secret)
	oracle := AesECBEncrypt(paddedPlainText, OracleKey)

	return oracle

}

func repeatByteA(n int) []byte {
	// Repeats the letter A as many times as needed

	return []byte(strings.Repeat("A", n))
}

func aesECBOracleBlocksize(secret []byte) int {

	initLen := len(AesECBOracleEncrypt([]byte("A"), secret)) // The initLen. Just one A

	oracleLen := initLen
	var oracle []byte

	for ind := range len(secret) {

		oracle = AesECBOracleEncrypt(repeatByteA(ind), secret)
		oracleLen = len(oracle)

		if oracleLen != initLen {
			break
		}

	}

	// TODO: Catch nil oracle type

	return oracleLen - initLen

}

func aesECBOracleSecretMaxLen(blockSize int, secret []byte) int {

	oracle := AesECBOracleEncrypt(repeatByteA(blockSize), secret)
	oracleLen := len(oracle)

	return oracleLen - blockSize

}

func AesECBOracleDecrypt(secret []byte) []byte {
	// Attack ECB

	blockSize := aesECBOracleBlocksize(secret)

	maxSecretLen := aesECBOracleSecretMaxLen(blockSize, secret)

	aData := repeatByteA(maxSecretLen)

	var ptSecret []byte

	for ind := maxSecretLen - 1; ind > 0; ind-- {

		trueBlock := AesECBOracleEncrypt(aData[:ind], secret)[:maxSecretLen]

		for char := range 256 {
			block := slices.Concat(aData[:ind], ptSecret, []byte{byte(char)})

			oracleBlock := AesECBOracleEncrypt(block, secret)[:maxSecretLen]

			if bytes.Equal(trueBlock, oracleBlock) {
				ptSecret = append(ptSecret, byte(char))
				break
			}

		}

	}

	return PKCS7UnPad(ptSecret)

}
