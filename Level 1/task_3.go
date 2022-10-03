package main

import (
	"fmt"
	"sync"
)

// Task 3
// Дана последовательность чисел: 2, 4, 6, 8, 10. Найти сумму их
// квадратов(2^2 + 3^2 + 4^2 + ...) с использованием конкурентных вычислений.

func getSquareSum(nums []int) int {
	wg := sync.WaitGroup{}
	wg.Add(len(nums))

	sum := 0

	for _, val := range nums {
		go func(val int) {
			defer wg.Done()
			fmt.Println(val * val)
			sum += val * val
		}(val)
	}
	wg.Wait()

	return sum
}

func main() {
	a := []int{2, 4, 6, 8, 10}
	fmt.Println(fmt.Sprintf("Res: %d", getSquareSum(a)))
}
