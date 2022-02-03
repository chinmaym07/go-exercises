package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://jsonplaceholder.typicode.com/todos" // "https://lco.dev" //"http://jsonplaceholder.typicode.com/todos"

func main() {
	fmt.Println("Web requests in golang")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T\n", response)
	defer response.Body.Close() // caller's responsibilty to close the connection

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataBytes)
	fmt.Println("DataByte ", content)
}
