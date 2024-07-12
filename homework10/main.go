package main

import "fmt"

func main() {
	a := 6.0 // Примерное значение первого числа
	b := 3.0 // Примерное значение второго числа
	sum := a*a + b*b
	diff := a*a - b*b
	prod := a * a * b * b
	quotient := (a * a) / (b * b)
	fmt.Println("Сумма квадратов:", sum)
	fmt.Println("Разность квадратов:", diff)
	fmt.Println("Произведение квадратов:", prod)
	fmt.Println("Частное квадратов:", quotient)
}
