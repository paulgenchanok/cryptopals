package main

import (
	s1 "cryptopals/set1"
	s2 "cryptopals/set2"
)

func runSet1() {
	s1.One()
	s1.Two()
	s1.Three()
	s1.Four() // Re-work. Scoring is not returning the right string.
	s1.Five()
	s1.Six()
	s1.Seven()
	s1.Eight()
}

func runSet2() {
	s2.Nine()
	s2.Ten()
	s2.Eleven()
	s2.Twelve()
}

func main() {

	// runSet1()
	runSet2()
	s2.Thirteen()

}
