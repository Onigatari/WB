package main

import (
	"fmt"
	"strings"
)

// Task 15

// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//    v := createHugeString(1 << 10)
//    justString = v[:100]
// }
//
// func main() {
//    someFunc()
// }

func createHugeString(num int) strings.Builder {
	var builder strings.Builder

	for i := 0; i < num; i++ {
		builder.WriteString("x")
	}

	return builder
}

// В итоге получается большая строка, которая могет не поместиться в памяти.
// Для работы со большими строками лучше исользовать strings.Builder.
// Т.к. эта структура подкопотом хранит символы в []byte
// За счет этого можно избавиться от полного копирования при изменении строки
func someFunc() {
	v := createHugeString(1 << 10)
	justString := v.String()[:100]
	fmt.Println(justString)
}

func main() {
	someFunc()
}
