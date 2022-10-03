package main

import (
	"fmt"
	"strconv"
)

// Task 21

// Реализовать паттерн «адаптер» на любом примере.

type StringRequest interface {
	getValue() string
}

type OldCurrStruct struct{}

func (x OldCurrStruct) getValue() string {
	return "2077"
}

// ===========================================================

type IntRequest interface {
	getValue() int
}

type FewNewStruct struct{}

func AdapterInteger() int {
	var tmp OldCurrStruct
	res, _ := strconv.Atoi(tmp.getValue())
	return res
}

func (x FewNewStruct) getValue() int {
	return AdapterInteger()
}

// ===========================================================

func main() {
	var a OldCurrStruct
	var b FewNewStruct

	fmt.Printf("Value: %v\t| Type: %T\n", a.getValue(), a.getValue())
	fmt.Printf("Value: %v\t| Type: %T\n", b.getValue(), b.getValue())
}
