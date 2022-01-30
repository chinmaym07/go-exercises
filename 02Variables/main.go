package main

import (
	"fmt"
)

func main() {
	var username string = "Chinmay"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)

	fmt.Printf("Variable is of type: %T\n", smallValue)

	var smallFloatValue float64 = 255.45678956789456789132
	fmt.Println(smallFloatValue)

	fmt.Printf("Variable is of type: %T\n", smallFloatValue)

	// default Values
	var anotherVariable float32
	fmt.Println(anotherVariable)

	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	website := "helloWorld"
	fmt.Println(website)

	fmt.Printf("Variable is of type: %T\n", website)
}
