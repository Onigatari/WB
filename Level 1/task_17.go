package main

import (
	"fmt"
	"sort"
)

// Task 17

// Реализовать бинарный поиск встроенными методами языка.

func BinarySearch(arr []int, item int) (int, bool) {
	lhs := 0
	rhs := len(arr) - 1
	for lhs <= rhs {
		mid := lhs + ((rhs - lhs) / 2)
		if arr[mid] == item {
			return mid, true
		} else if arr[mid] < item {
			lhs = mid + 1
		} else if arr[mid] > item {
			rhs = mid - 1
		}
	}

	return -1, false
}

func main() {
	arr := []int{1, 2, 3, 4, 6, 7, 8, 9}
	sort.Ints(arr)

	fmt.Println(BinarySearch(arr, 5))
}
