package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 25

// Реализовать собственную функцию sleep.

func Sleep(d time.Duration) {
	if d <= 0 {
		return
	}

	// Горутина будет ожидать пока NewTimer не передаст текущее время в канал
	<-time.NewTimer(d).C
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		fmt.Println("Начало")
		Sleep(2 * time.Second)
		fmt.Println("Конец")
		wg.Done()
	}()

	wg.Wait()
}
