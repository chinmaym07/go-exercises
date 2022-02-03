package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files in Golang")

	content := "This needs to go in a file test.txt"

	filePtr, err := os.Create("./text.txt")
	checkNilErr(err)

	ln, err := io.WriteString(filePtr, content)
	checkNilErr(err)
	fmt.Println("Length is", ln)
	defer filePtr.Close()

	readFile("./text.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text Data inside file is \n", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
