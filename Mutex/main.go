package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

func getStatusCode(url string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		panic("SOME ISSUE IN WEBSITE")
	} else {
		defer resp.Body.Close()
		mut.Lock()
		fmt.Println("DONE:", url)
		mut.Unlock()
	}
}

func main() {
	data := []string{
		"https://google.com",
		"https://gokapture.com",
		"https://gokapturehub.com",
	}

	for _, url := range data {
		wg.Add(1)
		go getStatusCode(url)
	}

	wg.Wait()
}
