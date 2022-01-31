package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("Days in a week", days)

	/* for _, val := range days {
		fmt.Println(val)
	} */
	/* for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	} */
	for i := range days {
		fmt.Println(days[i])
	}

	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 2 {
			goto tmo
		}
		if rougeValue == 5 {
			rougeValue++
			continue
		}
		fmt.Println("Rouge Value", rougeValue)
		rougeValue++
	}

tmo:
	fmt.Println("Welcome to world of Golang !!")
}
