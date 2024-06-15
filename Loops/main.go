package main

import "fmt"

func main() {
	var i int = 1

	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("Loop")
		break
	}

	for n := range 10 {
		fmt.Println(n)
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
