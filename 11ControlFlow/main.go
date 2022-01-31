package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 23
	result := ""
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 && loginCount < 20 {
		result = "Watch out"
	} else {
		result = "Vella User"
	}
	fmt.Println("Message: ", result)

	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}
}
