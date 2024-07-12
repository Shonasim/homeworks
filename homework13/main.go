package main

import "fmt"

func main() {
	fileSize := 4096
	kilobytes := fileSize / 1024
	fmt.Println("Полных килобайтов:", kilobytes)
}
