package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "orange"

	fruitList[3] = "watermelon"
	fmt.Println("Fruitlist is: ", fruitList)

	fmt.Println("Fruitlist length is:", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", vegList)

	fmt.Println("Vegy list length is:", len(vegList))
}
