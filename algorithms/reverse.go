package main

import "fmt"

func reverse(sent string) string {
	newWord := ""
	for _, i := range sent {
		newWord = string(i) + newWord
	}
	return newWord
}

func reverse_recurr(sent string) string {
	if len(sent) == 1 {
		return sent
	}

	return reverse_recurr(string(sent[1:])) + string(sent[0])
}

func main() {
	fmt.Println("Reverse a string using recursive code")
	fmt.Println(reverse_recurr("Hello World"))
	fmt.Println(reverse_recurr("Chinmay"))
	fmt.Println(reverse_recurr("Alok"))
	fmt.Println(reverse_recurr("Vaibhav"))

	fmt.Println("Reverse a string using Iteration code")
	fmt.Println(reverse("Hello World"))
	fmt.Println(reverse("Chinmay"))
	fmt.Println(reverse("Alok"))
	fmt.Println(reverse("Vaibhav"))
}
