package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b + c)

	var d = true
	fmt.Println(d)

	var e int = 0
	fmt.Println(e)

	// this is a shorthand for declaring and initializing a variable, and it is only available inside a function
	f := "apple"
	fmt.Println(f)
}
