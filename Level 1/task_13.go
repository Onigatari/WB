package main

import "fmt"

// Task 13

// Поменять местами два числа без создания временной переменной.

func SwapSum(x, y *int) {
	*x = *x + *y
	*y = *x - *y
	*x = *x - *y
}

func main() {
	var x, y int
	fmt.Println("Введите x и y:")
	_, err := fmt.Scan(&x, &y)
	if err != nil {
		return
	}

	fmt.Printf("x: %d | y: %d\n", x, y)
	SwapSum(&x, &y)
	fmt.Printf("x: %d | y: %d\n", x, y)
}
