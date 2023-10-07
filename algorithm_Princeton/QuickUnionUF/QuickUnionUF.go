package main

import "fmt"

var id []int

func main() {
	QuickUnionUF(10)
	union(3, 8)
	union(2, 4)
	union(4, 3)
	fmt.Println(id)
	fmt.Println(connected(2, 9))
}

func QuickUnionUF(N int) {
	for i := 0; i < N; i++ {
		id = append(id, i)
	}
}

func root(i int) int {
	for {
		if i == id[i] {
			break
		}
		i = id[i]
	}
	return i
}

func connected(p int, q int) bool {
	return root(p) == root(q)
}

func union(p int, q int) {
	var i, j int
	i = root(p)
	j = root(q)
	id[i] = j
}
