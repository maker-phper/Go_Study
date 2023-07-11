package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func commaUseBufferSupportFloat(s string, isFloat bool) string {
	var fractionalPart string
	if isFloat {
		sl := strings.Split(s, ".")
		s = sl[0]
		fractionalPart = "." + sl[1]
	}
	var buf bytes.Buffer
	for i, v := range s {
		if i > 0 && i%3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(v))
	}
	return buf.String() + fractionalPart
}

func main() {
	var f float64 = 1001231132.123
	fmt.Println(commaUseBufferSupportFloat(strconv.FormatFloat(f, 'f', -1, 64), true))
}
