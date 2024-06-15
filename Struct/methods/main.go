package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	r.height = 1
	return r.height * r.width
}

func (r rect) perim() int {
	return 2*r.height + 2*r.width
}

func main() {
	r := rect{width: 10, height: 5}
	rp := &r
	fmt.Println(rp.perim())
	fmt.Println(r)
	fmt.Println("Area", r.area())
	fmt.Println(r)
}
