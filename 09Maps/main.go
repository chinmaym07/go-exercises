package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Maps in Golang")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("List of all Languages: ", languages)

	fmt.Println("JS stands for: ", languages["JS"])
	fmt.Println("RB stabds for: ", languages["RB"])

	delete(languages, "JS")
	fmt.Println("List of all Languages: ", languages)

	// Looping in map
	for k, v := range languages {
		fmt.Printf("For Key: %v Value: %v\n", k, v)
	}
}
