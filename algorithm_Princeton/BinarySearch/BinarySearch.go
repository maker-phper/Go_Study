package main

import "fmt"

func main() {
	var data = []int{1, 12, 23, 34, 45, 56, 67}
	var key = 1
	fmt.Println(BinarySearch(key, data))
}

func BinarySearch(key int, data []int) bool {
	high := len(data) - 1
	for low := 0; low <= high; {
		mid := low + (high-low)/2
		if key < data[mid] {
			high = mid - 1
		} else if key > data[mid] {
			low = mid + 1
		} else {
			return true
		}
	}
	return false
}
