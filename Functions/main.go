package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

// consecutive paramater with same type (int)
func threeSome(a, b, c int) int {
	return a + b + c
}

// return multiple values
func values() (int, int, int) {
	return 1, 2, 3
}

// Variadic Function
func sumAll(nums ...int) int { // similar to rest operator in Javacript
	fmt.Println(nums)
	i := 0
	for _, num := range nums {
		i += num
	}
	return i
}

// Closures in Go
//
// ~~~~--------vvvvvvvvvv <- is the return value
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	result := sum(1, 2)
	fmt.Println(result)

	result = threeSome(1, 2, 3)
	println(result)

	a, b, d := values()
	println(a, b, d)

	nums := []int{1, 2, 3, 4}
	ans := sumAll(nums...) // simiar to spread operatoe in Javascript
	println(ans)

	generator := intSeq()
	generator()
	generator()

	// now generator will have 3
	fmt.Println(generator())
	fmt.Println(fact(120))

	// fibonacci
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
}
