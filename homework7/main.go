package main

import "fmt"

func main() {
	R := 5.0 // Примерное значение радиуса
	π := 3.14
	L := 2 * π * R
	S := π * R * R
	fmt.Println("Длина окружности:", L)
	fmt.Println("Площадь круга:", S)
}
