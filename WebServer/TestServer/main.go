package main

import (
	"fmt"
	"net/http"
	"sync"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	var wg sync.WaitGroup
	http.HandleFunc("/", handleRequest)

	wg.Add(1)
	go func() {
		println("Server Running on Port [::]:8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println(err)
		}
	}()

	wg.Wait()
}
