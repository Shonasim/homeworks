package main

import (
	"fmt"
)

func main() {
	var A, B int
	A = 17
	B = 5

	if A > B {

		x := A % B
		fmt.Printf("Длина незанятой части отрезка A: %d\n", x)
	} else {
		fmt.Println("A должно быть больше B")
	}
}