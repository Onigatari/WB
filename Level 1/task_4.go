package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Task 4
// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и выводят
// в stdout. Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
// способ завершения работы всех воркеров.

type Worker struct {
	num int
}

func (worker *Worker) Println(chanelInt chan int) {
	for {
		fmt.Println("Worker", worker.num, "| Value:", <-chanelInt)
		time.Sleep(time.Second)
	}
}

func main() {
	// Каналы обычный и для записи сигнала ОС
	chanelInt := make(chan int)
	osSignals := make(chan os.Signal, 1)

	// Если поступил SIGINT, генерируемый Ctrl+C, в канал будет записано значение
	signal.Notify(osSignals, syscall.SIGINT)

	fmt.Print("Введите кол-во воркеров: ")
	var workerCount int
	_, err := fmt.Scanln(&workerCount)
	// Если при вводе поймали ошибку или количество Воркеров меньше 1 завершаем программу
	if err != nil || workerCount < 1 {
		return
	}

	// Создаем объекты структуры Worker с номерами и запускаем у каждого в отдельной горутине функцию вывода
	for i := 1; i <= workerCount; i++ {
		worker := Worker{i}
		go worker.Println(chanelInt)
	}

	result := 1

	for {
		//Оператор select ожидает нескольких операций отправки или получения одновременно
		select {
		case chanelInt <- result:
			result++
		case <-osSignals:
			fmt.Println("Конец")
			os.Exit(0)
		}
	}
}
