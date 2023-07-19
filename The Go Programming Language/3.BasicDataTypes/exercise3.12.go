package main

import "fmt"

func main() {
	var s1, s2 string
	s1 = "abacdssa"
	s2 = "abacdsas"
	if compareAnagramsStrings(s1, s2) {
		fmt.Println("There are anagrams strings")
	} else {
		fmt.Println("There are not anagrams strings")
	}
	if isAnagrams(s1, s2) {
		fmt.Println("There are anagrams strings")
	} else {
		fmt.Println("There are not anagrams strings")
	}
}

func compareAnagramsStrings(s1, s2 string) bool { // best o(n), worst o(n^3)
	if (len(s1) != len(s2)) || (s1 == s2) {
		return false
	}
	for _, v1 := range s1 {
		for k2, v2 := range s2 {
			if v1 == v2 {
				s2 = s2[0:k2] + s2[k2+1:] //del character at s2
				break
			} else if k2 == (len(s2) - 1) { //Means s1 character is not in s2 string
				return false
			}
		}
	}
	return true
}

func isAnagrams(a, b string) bool { // better, Time complexity o(n)
	if (len(a) != len(b)) || (a == b) {
		return false
	}
	aFreq := make(map[rune]int)
	bFreq := make(map[rune]int)

	for _, v1 := range a {
		aFreq[v1]++
	}

	for _, v2 := range b {
		bFreq[v2]++
	}

	for k1, f1 := range aFreq {
		if bFreq[k1] != f1 {
			return false
		}
	}

	for k2, f2 := range bFreq {
		if aFreq[k2] != f2 {
			return false
		}
	}
	return true
}
