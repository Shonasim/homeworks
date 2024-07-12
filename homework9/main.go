package main

import (
	"fmt"
	"math"
)

func main() {
	a := 4.0
	b := 9.0
	meanGeo := math.Sqrt(a * b)
	fmt.Println("Среднее геометрическое:", meanGeo)
}
