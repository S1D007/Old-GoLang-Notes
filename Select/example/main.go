package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("received", msg1)
		case msg2 := <-channel2:
			fmt.Println("received", msg2)
		}
	}
}
