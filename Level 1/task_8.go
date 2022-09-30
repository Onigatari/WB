package main

import (
	"fmt"
)

// Task 8
// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

// При замене i-ого бита на 1
// 01010110
//      ^ вот этот
// mask = 00000001 << 3  = 00001000
//
// 01010110  or
// 00001000
// --------
// 01011110

// При замене i-ого бита на 0
// mask = 00000001 << 3 = 00001000
//
// 01011110  xor
// 00001000
// --------
// 01010110

func Sign(x int64) int64 {
	if x < 0 {
		return -1
	}
	return 1
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func setBit(number int64, bitPos int, bitValue int) int64 {
	// Проверка на знак числа
	var sign = Sign(number)
	number = Abs(number)

	// Проверки на то нужно ли менять соответстующий бит у числа
	if int((number&(1<<bitPos))>>bitPos) != bitValue {
		if bitValue == 1 {
			number |= 1 << bitPos
		} else {
			number ^= 1 << bitPos
		}
	}

	return sign * number
}

func main() {
	var number int64
	var bitPos, bitValue int

	fmt.Println("Введи через пробел: число, номер бита и значение, которое хочешь поместить в бит")
	_, err := fmt.Scanf("%d %d %d", &number, &bitPos, &bitValue)
	if err != nil {
		fmt.Println("Input error")
		return
	}

	fmt.Printf("Result: %v\n", number)
	fmt.Printf("Binary: %b\n", number)

	number = setBit(number, bitPos, bitValue)

	fmt.Printf("Result: %v\n", number)
	fmt.Printf("Binary: %b\n", number)
}
