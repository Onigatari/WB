package main

import (
	"context"
	"fmt"
	"time"
)

// Task 6
// Реализовать все возможные способы остановки выполнения горутины.

// Использование отдельного канала для отбработки завершения горутины
// Этот сигнальный канал используется для ввода значения, когда вы хотите,
// чтобы подпрограмма остановилась. Горутина регулярно опрашивает этот канал.
// Как только он обнаруживает сигнал, он завершает работу текущей горутины.
func goroutine1() {
	quit := make(chan bool)
	defer close(quit)

	output := make(chan string)
	defer close(output)

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Println("Горутина 1 — работает")
			}
		}
	}()

	time.Sleep(time.Millisecond)
	quit <- true

	fmt.Println("Горутина 1 — все")
}

// Когда на канале выполняется операция приема, мы проверяем,
// закрыт ли канал или нет, и выходим из горутины, если канал закрыт.
func goroutine2() {
	quit := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				quit <- true
				return
			default:
				fmt.Println("Горутина 2 — работает")
			}
		}
	}(ctx)

	time.Sleep(time.Millisecond)
	cancel()

	<-quit
	fmt.Println("Горутина 2 — все")
}

func main() {
	goroutine1()
	//goroutine2()
}
