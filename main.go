package main

import (
	"fmt"
)

func main() {
	operations2 := [...]int64{34, 235, -98, 234, 244, 981}

	for i := 0; i < len(operations2); i++ {
		fmt.Println(operations2[i])
	}

	operations1 := [3]int64{54, 22, -113}

	for _, operation := range operations1
	{
		fmt.Println(operation)
	}

	var operations[3]int64

	operations[0] = -12
	operations[1] = 23
	operations[2] = 112

	for index := range operations {
		operation := operations[index]
		fmt.Println(operation)
	}

	last := operations[2]
	fmt.Println(last)
}