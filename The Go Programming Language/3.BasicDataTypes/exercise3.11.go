package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func commaUseBufferSupportFloat(s string) string {
	var fractionalPart string
	var sign byte

	sl := strings.Split(s, ".")
	s = sl[0]
	fractionalPart = "." + sl[1]

	if s[0] == '-' || s[0] == '+' {
		sign = s[0]
		s = s[1:]
	}
	sNew := reversionStr(s)
	var buf bytes.Buffer
	for i, v := range sNew {
		if i > 0 && i%3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(v))
	}
	return string(sign) + reversionStr(buf.String()) + fractionalPart
}

func reversionStr(s string) string {
	sLen := len(s)
	sNew := make([]byte, sLen, sLen)
	for i, n := range s {
		sNew[sLen-i-1] = byte(n)
	}
	return string(sNew)
}

func comma(s string) string {
	b := bytes.Buffer{}
	mantissaStart := 0
	if s[0] == '+' || s[0] == '-' {
		b.WriteByte(s[0])
		mantissaStart = 1
	}
	mantissaEnd := strings.Index(s, ".")
	if mantissaEnd == -1 {
		mantissaEnd = len(s)
	}
	mantissa := s[mantissaStart:mantissaEnd]
	pre := len(mantissa) % 3
	if pre > 0 {
		b.Write([]byte(mantissa[:pre]))
		if len(mantissa) > pre {
			b.WriteString(",")
		}
	}
	for i, c := range mantissa[pre:] {
		if i%3 == 0 && i != 0 {
			b.WriteString(",")
		}
		b.WriteRune(c)
	}
	b.WriteString(s[mantissaEnd:])
	return b.String()
}

func main() {
	var f float64 = -10116516.12311
	fmt.Println(commaUseBufferSupportFloat(strconv.FormatFloat(f, 'f', -1, 64)))
	fmt.Println(comma(strconv.FormatFloat(f, 'f', -1, 64)))
}
