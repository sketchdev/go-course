package main // main is a special main, not-importable

import (
	"fmt"
	"hello/pkg/picker"
)

func main() { // entry point

	students := picker.Random(5)

	for _, student := range students {
		fmt.Printf("%v %v\n", student.FirstName, student.LastName)
	}
}
