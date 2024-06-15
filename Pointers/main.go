package main

import "fmt"

func zeroval(val int) {
	val = 0
	fmt.Println("For This Scope of Function, Value is", val)
}

func zeroprt(val *int) {
	fmt.Println(val)
	*val = 0
}

func main() {
	i := 1
	fmt.Println(i)
	zeroval(i)
	fmt.Println(i)
	zeroprt(&i)
	fmt.Println(i)

	// some cool stuff
	value := 5
	var ptr *int = &value
	*ptr *= value
	fmt.Println(*ptr)
}
