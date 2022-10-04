package main

import "fmt"

// Task 23

// Удалить i-ый элемент из слайса.

// Склеиваем два среза. s[start, index-1] + s[index+1, end].
func removeItemSlice(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	var num int
	fmt.Printf("Введите индекс для удаления: ")
	_, _ = fmt.Scan(&num)
	fmt.Println(removeItemSlice(arr, num))
}
