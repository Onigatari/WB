package main

import (
	"fmt"
	"sync"
)

// Task 18

// Реализовать структуру-счетчик, которая будет
// инкрементироваться в конкурентной среде.
// По завершению программа значение счетчика должна выводить итоговое

type Counter struct {
	value int
}

// Метод для увеличение счетчика
func (c *Counter) increment(mutex *sync.Mutex) {
	mutex.Lock()
	c.value++ // Горутина увеличивает счетчик
	mutex.Unlock()
}

func main() {
	var mutex sync.Mutex

	wg := sync.WaitGroup{}
	wg.Add(6)
	counter := Counter{}

	for i := 0; i < 6; i++ {
		go func() {
			counter.increment(&mutex)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter.value)
}
