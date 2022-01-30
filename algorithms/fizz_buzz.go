package main

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i < n; i++ {
		PrintFizzBuzzValue(i)
		fmt.Print(", ")
	}
	PrintFizzBuzzValue(n)
	fmt.Println("\n")
}

func PrintFizzBuzzValue(n int) {
	if n%15 == 0 {
		fmt.Print("Fizz Buzz")
	} else if n%3 == 0 {
		fmt.Print("Fizz")
	} else if n%5 == 0 {
		fmt.Print("Buzz")
	} else {
		fmt.Print(n)
	}
}
func main() {
	FizzBuzz(3)
	FizzBuzz(5)
	FizzBuzz(10)
	FizzBuzz(30)

}
