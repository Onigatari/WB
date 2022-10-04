package main

import (
	"fmt"
	"sync"
)

// Task 3
// Дана последовательность чисел: 2, 4, 6, 8, 10. Найти сумму их
// квадратов(2^2 + 3^2 + 4^2 + ...) с использованием конкурентных вычислений.

func getSquareSumWithMutex(nums []int) int {
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	wg.Add(len(nums))

	sum := 0

	for _, val := range nums {
		go func(val int, m *sync.Mutex) {
			defer wg.Done()
			fmt.Println(val * val)
			m.Lock()
			sum += val * val
			m.Unlock()
		}(val, &m)
	}
	wg.Wait()

	return sum
}

func getSquareSumWithChannel(nums []int) int {
	ch := make(chan int, 1)
	defer close(ch)

	wg := sync.WaitGroup{}
	wg.Add(len(nums))
	ch <- 0

	for _, val := range nums {
		go func(val int, w *sync.WaitGroup, c chan int) {
			defer wg.Done()
			fmt.Println(val * val)
			c <- (val * val) + <-c
		}(val, &wg, ch)
	}
	wg.Wait()

	return <-ch
}

func main() {
	a := []int{2, 4, 6, 8, 10}
	fmt.Println(fmt.Sprintf("Res: %d", getSquareSumWithMutex(a)))
	//fmt.Println(fmt.Sprintf("Res: %d", getSquareSumWithChannel(a)))
}
