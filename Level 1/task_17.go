package main

import (
	"fmt"
	"sort"
)

// Task 17

// Реализовать бинарный поиск встроенными методами языка.

func binarySearch(arr []int, lhs, rhs int, item int) (int, bool) {
	mid := lhs + ((rhs - lhs) / 2)
	switch {
	// Элемент в правом под массиве
	case arr[mid] > item:
		return binarySearch(arr, lhs, mid, item)
	// Элемент в левом под массиве
	case arr[mid] < item:
		return binarySearch(arr, mid+1, rhs, item)
	case arr[mid] == item:
		return mid, true
	}
	return -1, false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort.Ints(arr)

	fmt.Println(binarySearch(arr, 0, len(arr), 5))
}
