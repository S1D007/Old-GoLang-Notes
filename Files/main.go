package main

import (
	"fmt"
	"io"
	"os"
)

func ReadFileLength(filepath string) (length int, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return
	}
	return int(info.Size()), nil
}

func main() {
	content := "Hey My name is Siddhant"
	file, err := os.Create("./name.txt")
	if err != nil {
		panic("Some Issues")
	}
	io.WriteString(file, content)
	length, _ := ReadFileLength("./name.txt")
	fmt.Println(length)
	defer file.Close()
}
