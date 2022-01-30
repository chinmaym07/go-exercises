package main

import (
	"fmt"
)

func BaseToDec(value string, base int) int {
	ans := 0
	pow := 1
	n := len(value)
	/* obj := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"A": 10,
		"B": 11,
		"C": 12,
		"D": 13,
		"E": 14,
		"F": 15,
	}
	for i := n - 1; i >= 0; i-- {
		ans += obj[string(value[i])] * pow
		pow *= base
	} */
	for i := n - 1; i >= 0; i-- {
		var val int
		fmt.Sscanf(string(value[i]), "%X", &val)
		ans += val * pow
		pow *= base
	}
	return ans
}

func main() {
	fmt.Println(BaseToDec("1", 2))
	fmt.Println(BaseToDec("10", 2))
	fmt.Println(BaseToDec("21", 3))
	fmt.Println(BaseToDec("1110", 2))
	fmt.Println(BaseToDec("E", 16))
	fmt.Println(BaseToDec("11", 16))
	fmt.Println(BaseToDec("11111", 2))
	fmt.Println(BaseToDec("D", 16))
	fmt.Println(BaseToDec("F", 16))
	fmt.Println(BaseToDec("1F", 16))
}
