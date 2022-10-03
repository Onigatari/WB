package main

import (
	"fmt"
	"math/big"
)

// Task 22

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.

func main() {
	// Для длинной арифметики используем math/big
	var x int64

	_, _ = fmt.Scan(&x)
	a := big.NewInt(x)

	_, _ = fmt.Scan(&x)
	b := big.NewInt(x)

	fmt.Printf("Сумма = %d \n", big.NewInt(0).Add(a, b))

	fmt.Printf("Вычитание = %d \n", big.NewInt(0).Sub(a, b))

	fmt.Printf("Умножение = %d \n", big.NewInt(0).Mul(a, b))

	fmt.Printf("Деление = %f \n", big.NewFloat(0.0).Quo(big.NewFloat(0.0).SetInt(a), big.NewFloat(0.0).SetInt(b)))

}
