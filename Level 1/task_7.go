package main

import (
	"fmt"
	"sync"
)

// Task 7
// Реализовать конкурентную запись данных в map.

func main() {
	counters := make(map[int]int)
	mu := &sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int, mu *sync.Mutex) {
			for j := 0; j < 5; j++ {
				mu.Lock()
				counters[j]++
				mu.Unlock()
			}
		}(counters, i, mu)
		wg.Done()
	}

	wg.Wait()

	for key, val := range counters {
		fmt.Println(key, ": ", val)
	}
}
