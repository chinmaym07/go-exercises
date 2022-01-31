package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")
	greeter()
	result := adder(3, 5)
	fmt.Println("Result is", result)
	result2, mes := proAdder(3, 5, 4, 6, 8)
	fmt.Println("Pro Result is", result2)
	fmt.Println("Pro Message is", mes)
}

func greeter() {
	fmt.Println("Namastey from Golang!")
}

func adder(num1, num2 int) int {
	return num1 + num2
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi from Pro Adder Function"
}
