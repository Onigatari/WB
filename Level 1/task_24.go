package main

import (
	point "Level_1/point"
	"fmt"
)

// Task 24

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x, y и
// конструктором.

func main() {
	p1 := point.CreatePoint(0, 0)
	p2 := point.CreatePoint(3, 4)
	fmt.Println(point.GetDistance(p1, p2))
}
