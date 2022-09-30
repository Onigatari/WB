package main

import "fmt"

// Task 14

// Разработать программу, которая в рантайме способна определить
// тип переменной: int, string, bool, channel из переменной типа interface{}.

// Используем пустой интерфейс, для использвания объектов любого типа
func getType(v interface{}) string {
	// %T представление типа значения в Go
	return fmt.Sprintf("%T", v)
}

func main() {
	a := make(chan int)
	fmt.Println("Тип переменной:", getType(a))

	var b float32
	fmt.Println("Тип переменной:", getType(b))

	var c map[string][]struct{}
	fmt.Println("Тип переменной:", getType(c))
}
