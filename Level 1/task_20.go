package main

import (
	"fmt"
	"strings"
)

// Task 20

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func reverseSentence(s string) string {
	arr := strings.Split(s, " ")
	n := len(arr)

	// Проход до середины строки
	for i := 0; 2*i < len(arr); i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}

	return strings.Join(arr, " ")
}

func main() {
	a := "snow dog sun"
	fmt.Println(reverseSentence(a))

	b := "аргентина манит негра"
	fmt.Println(reverseSentence(b))

	c := "外 人 外 人 12 46"
	fmt.Println(reverseSentence(c))
}
