package main

import "fmt"

var id []int
var sz []int

func main() {
	QuickUnionUF(10)
	//union(3, 8)
	//union(2, 4)
	//union(4, 3)

	quickUnion(3, 8)
	quickUnion(2, 4)
	quickUnion(4, 3)
	quickUnion(8, 1)
	fmt.Println(id)
	fmt.Println(sz)
	fmt.Println(connected(2, 9))
}

func QuickUnionUF(N int) {
	for i := 0; i < N; i++ {
		id = append(id, i)
		sz = append(sz, 1)
	}
}

func root(i int) int {
	for {
		if i == id[i] {
			break
		}
		id[i] = id[id[i]]
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

func quickUnion(p int, q int) {
	var i, j int
	i = root(p)
	j = root(q)
	if i == j {
		return
	}
	if sz[i] < sz[j] {
		id[i] = j
		sz[j] += sz[i]
	} else {
		id[j] = i
		sz[i] += sz[j]
	}
}
