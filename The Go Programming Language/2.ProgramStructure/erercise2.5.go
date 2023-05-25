package main

func PopCount(x uint64) int {
	var count int
	for x > 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
