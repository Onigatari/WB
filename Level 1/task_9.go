package main

import (
	"fmt"
)

// Task 9
// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

func main() {
	inputChannel := make(chan int)
	outputChannel := make(chan int)

	var array [10]int
	for i := 0; i < len(array); i++ {
		array[i] = i + 1
	}

	// Горутина для записи значение из массива в первый канал
	go func() {
		for _, val := range array {
			inputChannel <- val
		}
		close(inputChannel)
	}()

	// Горутина для записи обработки значений из певого канала и записи во второй
	go func() {
		for {
			x, ok := <-inputChannel
			if !ok {
				close(outputChannel)
				break
			}
			outputChannel <- x * 2
		}
	}()

	// Вывод второго канала
	for i := 0; i < len(array); i++ {
		fmt.Println(<-outputChannel)
	}
}
