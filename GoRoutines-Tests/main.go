package main

import (
	"fmt"
	"time"
)

func runCounter() {
	for i := range 5 {
		fmt.Println(i + 1)
	}
}

func main() {
	fmt.Println("First line inside the Main Function")
	go runCounter()
	time.Sleep(time.Second / 999)
	fmt.Println("Last line inside the Main Function")
}
