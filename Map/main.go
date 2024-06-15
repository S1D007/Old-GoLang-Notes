package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)

	value1 := m["one"]
	fmt.Println(value1)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

	// delete
	delete(n, "foo")

	// check if key exists
	_, ok := n["foo"]
	fmt.Println(ok)

	// iterate
	for key, value := range n {
		fmt.Println(key, value)
	}

	x1 := map[string]int{"foo": 1, "bar": 2}
	x2 := map[string]int{"foo": 1, "bar": 2}

	// compare
	fmt.Println(maps.Equal(x1, x2))
}
