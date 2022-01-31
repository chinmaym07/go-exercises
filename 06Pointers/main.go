package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	var ptr *int
	fmt.Println("Default Value of pointer is ", ptr)

	myNumber := 23
	newPtr := &myNumber
	fmt.Println("Value of actual myPtr", *newPtr)
	fmt.Println("Value of myPtr", newPtr)
	*newPtr = (*newPtr) * 2

	fmt.Println("Value of myNumber", myNumber)

}
