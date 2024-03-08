package main

import "fmt"

func main() {
	arr := []int{5, 4, 6, 3, 1, 2, 7}
	fmt.Println(InsertSort(arr))
}

func InsertSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if CompareTo(arr[j], arr[minIndex]) == -1 {
				minIndex = j
			}
		}
		Swap(&arr[i], &arr[minIndex])
	}
	return arr
}

func Swap(i *int, j *int) {
	*i, *j = *j, *i
}

func CompareTo(i int, j int) int {
	if i > j {
		return 1
	} else if i < j {
		return -1
	} else {
		return 0
	}
}


package main

import "fmt"

func main() {
	arr := []int{5, 4, 6, 3, 1, 2, 7}
	fmt.Println(SelectionSort(arr))
}

func SelectionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if CompareTo(arr[j], arr[j-1]) == -1 {
				Swap(&arr[j], &arr[j-1])
			}
		}
	}
	return arr
}

func Swap(i *int, j *int) {
	*i, *j = *j, *i
}

func CompareTo(i int, j int) int {
	if i > j {
		return 1
	} else if i < j {
		return -1
	} else {
		return 0
	}
}
