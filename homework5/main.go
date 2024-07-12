package main

import "fmt"

func main() {
	a := 3.0 // Примерное значение длины ребра куба
	V := a * a * a
	S := 6 * a * a
	fmt.Println("Объем куба:", V)
	fmt.Println("Площадь поверхности куба:", S)
}
