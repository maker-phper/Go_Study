package main

import "fmt"

func main() {
	var a [5]int
	a = [5]int{1, 2, 3, 4, 5}
	reverseArr(&a)
	fmt.Println(a)
}

func reverseArr(arr *[5]int) {
	for i := 0; i < len(arr)/2; i++ {
		end := len(arr) - i - 1
		arr[i], arr[end] = arr[end], arr[i]
	}
}
