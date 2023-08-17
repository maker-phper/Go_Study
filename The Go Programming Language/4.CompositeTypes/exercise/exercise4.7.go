package main

import (
	"bytes"
	"fmt"
)

func reverse(b []byte) []byte {
	runes := bytes.Runes(b)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return []byte(string(runes))
}

func main() {
	b := []byte("abcde")
	s := reverse(b)
	fmt.Printf("%p, %s", s, s)
}
