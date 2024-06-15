package main

import (
	"fmt"
)

func worker(done chan bool) {
	fmt.Print("working...")
	fmt.Println("done")
	done <- true
}

func sendOnly(ch chan<- string) {
	ch <- "Hello"
}

func receiveOnly(ch <-chan string) {
	msg := <-ch
	fmt.Println(msg)
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done

	messages := make(chan string)
	go sendOnly(messages)
}
