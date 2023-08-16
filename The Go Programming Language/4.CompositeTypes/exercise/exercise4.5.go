package main

import "fmt"

func main() {
	s := []string{"a", "a", "b", "b", "b", "c", "d", "d", "e"}
	res := eliminateDuplicatesElement(s)
	res_u := unique(s)
	fmt.Println(res)
	fmt.Println(res_u)
}
func eliminateDuplicatesElement(str []string) []string {
	last := str[len(str)-1]
	res := make([]string, len(str))
	for i := 0; i < len(str); i++ {
		ss := str[i]
		if i+1 < len(str) && ss == str[i+1] {
			copy(str[i:], str[i+1:])
			if last == str[i] {
				res = str[:i+1]
				break
			}
			i--
		} else {
			res[i] = ss
		}
	}
	return res
}

//https://github.com/torbiak/gopl/blob/master/ex4.5/unique.go
func unique(strs []string) []string {
	w := 0
	for _, s := range strs {
		if strs[w] == s {
			continue
		}
		w++
		strs[w] = s
	}
	return strs[:w+1]
}
