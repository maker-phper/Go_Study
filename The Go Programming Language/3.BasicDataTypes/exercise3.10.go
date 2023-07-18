package main

import (
	"bytes"
	"fmt"
)

func commaUseBuffer(s string) string {
	var buf bytes.Buffer
	sNew := reversionStr(s)
	for i, v := range sNew {
		if i > 0 && i%3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(v))

	}
	return reversionStr(buf.String())
}

func reversionStr(s string) string {
	sLen := len(s)
	sNew := make([]byte, sLen, sLen)
	for i, n := range s {
		sNew[sLen-i-1] = byte(n)
	}
	return string(sNew)
}

func main() {
	fmt.Println(commaUseBuffer("1234567891"))
}
