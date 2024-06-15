package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	printNumbers()
}

func printNumbers() {
	for i := range 10 {
		defer fmt.Println(i)
	}
}
