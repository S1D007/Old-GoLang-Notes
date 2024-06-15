package main

import "fmt"

type person struct {
	name string
	age  int
}

func createPerson(name string) *person {
	p := person{name: name}
	p.age = 17
	return &p
}

func main() {
	fmt.Println(person{"Bob Bhai", 20})
	fmt.Println(person{name: "Jack Bhai"})
	fmt.Println(person{name: "Oggy", age: 12})
	fmt.Println(&person{name: "Olive", age: 13})

	s := person{name: "Siddhant", age: 24}
	sp := &s
	fmt.Println(sp.age)

	ptrPerson := createPerson("siddhant")
	fmt.Println(ptrPerson)
}
