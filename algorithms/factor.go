package main

import "fmt"

func Factor(num int) []int {
	var factorArr []int
	for i := 2; i < num; i++ {
		if num%i == 0 {
			factorArr = append(factorArr, i)
		}

	}
	return factorArr
}

func main() {
	fmt.Println(Factor(30))  // 2,3,5
	fmt.Println(Factor(28))  // 2,2,7
	fmt.Println(Factor(120)) //
	fmt.Println(Factor(720))
}
