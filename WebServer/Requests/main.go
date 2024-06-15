package main

import (
	"io"
	"net/http"
	"os"
)

func checkError(err error) bool {
	if err != nil {
		return false
	}
	return true
}

func saveResults(path string, contents string) (done bool) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return false
	}
	checkError(err)

	_, err = io.WriteString(file, contents)
	checkError(err)

	return true
}

func main() {
	const URL = "http://gokapturehub.com"
	resp, err := http.Get(URL)
	checkError(err)
	defer resp.Body.Close()
	data, e := io.ReadAll(resp.Body)
	checkError(e)
	saveResults("./result.html", string(data))

}
