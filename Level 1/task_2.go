package main

import (
	"fmt"
	"sync"
)

// Task 2
// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func squareWithWaitGroup(nums []int) {
	var wg sync.WaitGroup

	for _, val := range nums {
		wg.Add(1)
		// Возведение в квадрат с использованием WaitGroup
		go func(val int) {
			defer wg.Done()
			fmt.Println(fmt.Sprintf("Wait Group: %d", val*val))
		}(val)
	}

	wg.Wait()
}

func squareWithChanel(nums []int) {
	tempChanel := make(chan int, len(nums))
	defer close(tempChanel)

	// Перебор массива и создание горутина для каждого элемента
	for _, val := range nums {
		// Возведение в квадрат с использованием Каналов
		go func(number int, chanelInt chan int) {
			chanelInt <- number * number
		}(val, tempChanel)
	}

	for _, _ = range nums {
		fmt.Println(fmt.Sprintf("Chanel: %d", <-tempChanel))
	}
}

func main() {
	squareWithWaitGroup([]int{1, 2, 3, 4})
	squareWithChanel([]int{1, 2, 3, 4})
}
