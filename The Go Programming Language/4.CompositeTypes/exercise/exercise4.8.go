package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	runeTypeCounts := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int
	var invalid int

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar {
			invalid++
			continue
		}
		for catName, rangTable := range unicode.Properties {
			fmt.Println(catName)
			fmt.Println(rangTable)
			if unicode.In(r, rangTable) {
				runeTypeCounts[catName]++
			}
		}
		/*if unicode.IsLetter(r) {
			runeTypeCounts["letter"]++
		}
		if unicode.IsDigit(r) {
			runeTypeCounts["digit"]++
		}
		if unicode.IsSpace(r) {
			runeTypeCounts["space"]++
		}*/
		counts[r]++
		utflen[n]++
	}

	fmt.Println(counts)
	fmt.Println(utflen)
	fmt.Println(runeTypeCounts)
}
