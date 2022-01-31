package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch & Case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1
	fmt.Println("Value of Dice is ", diceNum)

	// fallthrough is used when we want more than 1 case to be executed

	switch diceNum {
	case 1:
		fmt.Println("Dice Value is 1 & you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
	case 4:
		fmt.Println("you can move 4 spot")
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 6 spot & roll dice again")
	default:
		fmt.Println("What was that!")
	}

}
