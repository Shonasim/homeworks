package main

import "fmt"

func main() {
	a := 3.0
	b := 4.0
	c := 5.0
	V := a * b * c
	S := 2 * (a*b + b*c + a*c)
	fmt.Println("Объем параллелепипеда:", V)
	fmt.Println("Площадь поверхности параллелепипеда:", S)
}