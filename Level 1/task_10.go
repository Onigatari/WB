package main

import "fmt"

// Task 10

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
// 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
// градусов. Последовательность в подмножноствах не важна.

// Вывод групп из мапы
func printGroup(tmp *map[int][]float32) {
	for key, val := range *tmp {
		fmt.Printf("%d:\t%v\n", key, val)
	}
}

// Создание групп интервалов
func createTempGroups(temperature *[]float32) map[int][]float32 {
	groups := make(map[int][]float32)
	for _, val := range *temperature {
		// Считаем количество десяток, для определение группы
		key := int(val/10) * 10
		groups[key] = append(groups[key], val)
	}

	return groups
}

func main() {
	temperature := []float32{0, 10, 45, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 9}
	groups := createTempGroups(&temperature)
	printGroup(&groups)
}
