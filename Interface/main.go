package main

import "fmt"

type Say interface {
	myName() string
}

type Person struct {
	name string
}

func (p Person) myName() string {
	return "Hey, How are you " + p.name
}

func execute(ISay Say) {
	fmt.Println(ISay.myName())
}

func main() {
	person := Person{name: "Siddhant"}
	execute(person)
}
