package main

import (
	"fmt"
	"strings"
)

// Функции, не принимающие аргументы и ничего не возвращающие

func PrintGreeting() {
	fmt.Println("Hello, World!")
}

func DisplaySeparator() {
	fmt.Println(strings.Repeat("-", 20))
}

func NotifyStart() {
	fmt.Println("Program started")
}

// Функции, не принимающие аргументы, но возвращающие значение

func GetWelcomeMessage() string {
	return "Welcome!"
}

func GetPiValue() float64 {
	return 3.14
}

func IsServerActive() bool {
	return true
}

// Функции, принимающие аргументы, но ничего не возвращающие

func PrintNumber(number int) {
	fmt.Println(number)
}

func GreetUser(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func ToggleLight(state bool) {
	if state {
		fmt.Println("Light on")
	} else {
		fmt.Println("Light off")
	}
}

// Функции, принимающие аргументы и возвращающие значение

func Add(a, b int) int {
	return a + b
}

func Concat(str1, str2 string) string {
	return str1 + str2
}

func IsEven(number int) bool {
	return number%2 == 0
}

// Функции как переменные

var adder = func(a, b int) int {
	return a + b
}

var concatenator = func(str1, str2 string) string {
	return str1 + str2
}

var isPositive = func(number int) bool {
	return number > 0
}

// Функции как аргументы функции

func Calculate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func Execute(value bool, action func(bool)) {
	action(value)
}

func Apply(value int, operation func(int) int) int {
	return operation(value)
}

// Функции как возвращаемые значения (callback)

func Multiplier(factor int) func(int) int {
	return func(value int) int {
		return value * factor
	}
}

func StringRepeater(n int) func(string) string {
	return func(s string) string {
		return strings.Repeat(s, n)
	}
}

func ConditionalPrinter(condition bool) func(string) {
	return func(message string) {
		if condition {
			fmt.Println(message)
		}
	}
}

// Основная функция для демонстрации работы всех функций

func main() {
	PrintGreeting()
	DisplaySeparator()
	NotifyStart()
	fmt.Println(GetWelcomeMessage())
	fmt.Println(GetPiValue())
	fmt.Println(IsServerActive())
	PrintNumber(42)
	GreetUser("Alice")
	ToggleLight(true)
	ToggleLight(false)
	fmt.Println(Add(3, 5))
	fmt.Println(Concat("Hello, ", "World!"))
	fmt.Println(IsEven(4))
	fmt.Println(IsEven(5))
	fmt.Println(adder(7, 3))
	fmt.Println(concatenator("Go", "Lang"))
	fmt.Println(isPositive(10))
	fmt.Println(isPositive(-1))
	fmt.Println(Calculate(10, 20, adder))
	Execute(true, ToggleLight)
	Execute(false, ToggleLight)
	fmt.Println(Apply(5, func(x int) int { return x * x }))
	mult := Multiplier(3)
	fmt.Println(mult(10))
	repeat := StringRepeater(3)
	fmt.Println(repeat("Go"))
	printIfTrue := ConditionalPrinter(true)
	printIfTrue("This should print")
	printIfFalse := ConditionalPrinter(false)
	printIfFalse("This should not print")
}
