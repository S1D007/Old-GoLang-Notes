package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	x1 := []string{"1", "2", "3"}
	x2 := []string{"1", "2", "3"}

	if slices.Equal(x1, x2) {
		fmt.Println("x and x2 are equal")
	}

	c := make([]string, len((s)))

	copy(c, s)
	fmt.Println(c)
}
