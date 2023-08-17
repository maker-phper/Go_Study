package main

import (
	"fmt"
	"unicode"
)

func removeDupSpace(b []byte) []byte {
	w := 0
	for _, s := range b {
		if unicode.IsSpace(rune(s)) {
			if w == 0 || (w > 0 && !unicode.IsSpace(rune(b[w-1]))) {
				b[w] = ' '
				w++
			}
			continue
		}
		b[w] = s
		w++
	}
	return b[:w]
}

func main() {
	b := []byte("abc\r  \n\rdef\n\r ghi")
	//fmt.Printf("%q\n", b)
	fmt.Printf("%q\n", string(removeDupSpace(b)))
}
