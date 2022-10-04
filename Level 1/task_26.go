package main

import "fmt"

// Task 26

// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

// Использую хэш-таблицу для предподсчета
// Time: O(n)
// Memory: O(n)
func IsUnique(str string) bool {
	cnt := make(map[rune]bool)
	for _, val := range str {
		if _, ok := cnt[val]; ok {
			return false
		}
		cnt[val] = true
	}
	return true
}

func main() {
	fmt.Println(IsUnique("abcd"))
	fmt.Println(IsUnique("acDefAaf"))
	fmt.Println(IsUnique("aabcd"))
	fmt.Println(IsUnique("外人外人"))
	fmt.Println(IsUnique("外人"))
}
