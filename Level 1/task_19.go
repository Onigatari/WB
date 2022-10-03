package main

import "fmt"

// Task 19

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func reversString(s string) string {
	arr := []rune(s)
	n := len(arr)

	// Проход до середины строки
	for i := 0; 2*i < len(arr); i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}

	return string(arr)
}

func main() {
	a := "аргентинаманитнегра"
	fmt.Println(reversString(a))

	b := "afafaasfasfaf"
	fmt.Println(reversString(b))

	c := "外人"
	fmt.Println(reversString(c))
}
