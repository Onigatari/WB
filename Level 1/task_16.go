package main

import "fmt"

// Task 16

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// Time: O(n*log(n))
func quickSort(arr []int, lhs, rhs int) {
	if lhs < rhs {
		var pivot = partition(arr, lhs, rhs)
		quickSort(arr, lhs, pivot)
		quickSort(arr, pivot+1, rhs)
	}
}

func partition(arr []int, low, high int) int {
	var pivot = arr[low]
	var lhs = low
	var rhs = high

	for lhs < rhs {
		for arr[lhs] <= pivot && lhs < high {
			lhs++
		}
		for arr[rhs] > pivot && rhs > low {
			rhs--
		}

		if lhs < rhs {
			arr[lhs], arr[rhs] = arr[rhs], arr[lhs]
		}
	}

	arr[low] = arr[rhs]
	arr[rhs] = pivot

	return rhs
}

func main() {
	a := []int{1, -2, 3, 4, -5}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
