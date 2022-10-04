package main

import (
	"fmt"
)

// Task 1
// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Human - абстрактная структура человека
type Human struct {
	name   string
	age    uint
	height uint
	weight uint
}

func (h *Human) GetHuman() string {
	return fmt.Sprintf("Имя: %s | Возраст: %d", h.name, h.age)
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) GetAge() uint {
	return h.age
}

func (h *Human) GetHeight() uint {
	return h.height
}

func (h *Human) GetWeight() uint {
	return h.weight
}

type Action struct {
	Human
}

func main() {
	h := Human{
		name:   "Sokrat",
		age:    21,
		height: 183,
		weight: 65,
	}

	a := Action{h}
	fmt.Printf("Type: %T | Value: %v\n", h, h)
	fmt.Printf("Type: %T | Value: %v\n", a, a)
	fmt.Println("========================")
	fmt.Println(h.GetHuman())
	fmt.Println(a.GetHuman())
	fmt.Println("========================")
	fmt.Println(h.GetName())
	fmt.Println(a.GetName())
	fmt.Println("========================")
	fmt.Println(h.GetHeight())
	fmt.Println(a.GetHeight())
}
