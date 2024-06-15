package main

func main() {
	num := [...]int{1, 2, 3, 4}
	for i, v := range num {
		println(i, v)
	}

	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		println(k, ":", v)
	}
}
