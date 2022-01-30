package main

import (
	"fmt"
)

func FindTwoThatSum(numbers []int, sum int) (int, int) {
	// solution 1
	/* sort.Ints(numbers)
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] == sum {
			break
		} else if numbers[i]+numbers[j] < sum {
			i++
		} else {
			j--
		}

	} */
	obj := map[int]int{}
	for _, num := range numbers {
		obj[num] = 1
	}
	var nm, flg int
	for _, num := range numbers {
		if sum-num != num && obj[sum-num] == 1 {
			nm = num
			flg = 1
			break
		}
	}
	if flg != 1 {
		return -1, -1
	} else {
		return sum - nm, nm
	}

}

func main() {
	fmt.Println(FindTwoThatSum([]int{1, 2, 3, 4}, 7)) // 3,4
	fmt.Println(FindTwoThatSum([]int{0, -1, 1}, 0))   // -1,1
	fmt.Println(FindTwoThatSum([]int{0, 1, 1}, 0))    // 0,0

	fmt.Println(FindTwoThatSum([]int{10, 1, 12, 3, 7, 2, 2, 1}, 4)) // 1,3
}
