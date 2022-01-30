package main

import "fmt"

func Sum(numbers []int) int {
	s := 0
	for _, x := range numbers {
		s += x
	}
	return s
}
func SumRecur(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + SumRecur(numbers[1:])
}

func main() {
	fmt.Println("Sum using Sum Function")
	fmt.Println(Sum([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(Sum([]int{1, 2, 3, -4, 5, -1}))
	fmt.Println(Sum([]int{1, 2, -3, 4, -5, 6}))
	fmt.Println(Sum([]int{}))
	fmt.Println(Sum(nil))
	fmt.Println(Sum([]int{1, 2, 3, 3, 2, -4}))
	fmt.Println("Sum using SumRecurr")
	fmt.Println(SumRecur([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(SumRecur([]int{1, 2, 3, -4, 5, -1}))
	fmt.Println(SumRecur([]int{1, 2, -3, 4, -5, 6}))
	fmt.Println(SumRecur([]int{}))
	fmt.Println(SumRecur(nil))
	fmt.Println(SumRecur([]int{1, 2, 3, 3, 2, -4}))
}
