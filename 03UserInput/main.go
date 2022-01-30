package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating for out pizza")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating out Pizza, ", input)
}
