package main

import (
	"fmt"
	"log"
	"task_2/unpack"
)

func main() {
	str := "a4bc2d5e"

	res, err := unpack.UnpackingString(str)
	if err != nil {
		log.Fatal("Error!")
	}

	fmt.Println(res)
}
