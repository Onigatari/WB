package main

import (
	"context"
	"fmt"
	"time"
)

// Task 5
// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.

func Sender(ctx context.Context, chanelInt chan int) {
	tmp := 0
	for {
		select {
		case chanelInt <- tmp:
			// Отправка в канал
			tmp++
		case <-ctx.Done():
			// Пока в канал Done не получить что-то на вход
			close(chanelInt)
			return
		}
	}

}

func main() {
	chanelInt := make(chan int)
	var seconds time.Duration

	fmt.Print("Введите количество секунд: ")
	_, err := fmt.Scanln(&seconds)

	// Проверка ввода на ошибки и кол-ва секунд больше 0
	if err != nil || seconds < 0 {
		return
	}

	// Создаем контекст, в который передаем время,
	// канал контекста закрывается, когда
	// возвращается функция cancel или когда
	// закрывается канал Done родительского контекста,
	// в зависимости от того, что произойдет раньше.

	ctx, cancel := context.WithTimeout(context.Background(), seconds*time.Second)
	defer cancel()

	go Sender(ctx, chanelInt)

	for {
		tmp, err := <-chanelInt
		if !err {
			break
		}
		fmt.Println(tmp)
	}

	fmt.Println("Время вышло")
}
