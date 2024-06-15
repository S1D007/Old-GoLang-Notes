package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[4] = 100

	fmt.Println(a)
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var c = [5]int{1, 2, 3, 4, 5}
	fmt.Println(c)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int

	for row := 0; row < len(twoD); row++ {
		fmt.Println("Row", row)
		for column := 0; column < len(twoD[row]); column++ {
			fmt.Println("Coulmn", column)
			/*
				[
					[] <- 0th
					[] <- 1th
				]
			*/
			twoD[row][column] = row + column
		}
	}

	fmt.Println(twoD)

	twoDArray := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println(twoDArray)
}
