package main

import "fmt"

const thisIsString string = "HEY THIS IS A STRING DECLARED AS A CONSTANT"

func main() {
	fmt.Println(thisIsString)

	const n = 500000000
	fmt.Println(n)

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
}
