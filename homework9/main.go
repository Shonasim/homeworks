package main

import (
	"fmt"
	"math"
)

func main() {
	a := 4.0 // Примерное значение первого числа
	b := 9.0 // Примерное значение второго числа
	meanGeo := math.Sqrt(a * b)
	fmt.Println("Среднее геометрическое:", meanGeo)
}
