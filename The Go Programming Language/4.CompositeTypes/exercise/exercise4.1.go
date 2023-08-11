package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var a, b string
	a = "abcde"
	b = "qerttccdcssdsd"
	fmt.Println(ShaBitDiff([]byte(a), []byte(b)))
}

/**
计算byte数二进制中1的数量
*/
func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}
	return count
}

func ShaBitDiff(a, b []byte) int {
	shaA := sha256.Sum256(a)
	shaB := sha256.Sum256(b)
	return bitDiff(shaA[:], shaB[:])
}

func bitDiff(a, b []byte) int {
	count := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		switch {
		case i >= len(a):
			count += popCount(b[i])
		case i >= len(b):
			count += popCount(a[i])
		default:
			count += popCount(a[i] ^ b[i]) // a,b byte 做异或,得到a,b两个元素的二进制byte中相同位不同数的结果,例如111^101,得到10
		}
	}
	return count
}
