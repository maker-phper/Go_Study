package main

import "fmt"

var id []int

func main() {
	QuickFindUF(5)
	Union(3, 4)
	fmt.Println(id)
}

func QuickFindUF(N int) {
	for i := 0; i < N; i++ {
		id = append(id, i)
	}
}

func Connected(p int, q int) bool {
	return id[p] == id[q]
}

func Union(p int, q int) {
	pid := id[p]
	qid := id[q]
	for i := 0; i < len(id); i++ {
		if id[i] == pid {
			id[i] = qid
		}
	}
}
