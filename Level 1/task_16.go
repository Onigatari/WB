package main

import "fmt"

// Task 16

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// Time: O(n*log(n))
func QuickSort(arr []int) {
	// Берем индекс в качестве опорной элемента
	pivot := len(arr) - 1
	// Если мы имеем меньше 2‑х элементов, то выходим из рекурсии
	if pivot < 1 {
		return
	}
	// left будет в качестве стартового индекса
	left := 0
	// Все что меньше pivot будем класть в left
	// и переходить к следующему индексу, и так не доходя до pivot
	for i, _ := range arr {
		if arr[i] < arr[pivot] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	// перемещаем pivot на свое законное место на следующую ячейку от left
	arr[left], arr[pivot] = arr[pivot], arr[left]

	// Разделяем на рекурсию, на левую и правую часть, не трогая часть pivot
	QuickSort(arr[left+1:])
	QuickSort(arr[:left])
}

func main() {
	a := []int{1, -2, 3, 4, -5}
	QuickSort(a)
	fmt.Println(a)
}
