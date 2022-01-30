package main

import "fmt"

func BaseToDec(value string, base int) int {
	ans := 0
	pow := 1
	n := len(value)
	for i := n - 1; i >= 0; i-- {
		var val int
		fmt.Sscanf(string(value[i]), "%X", &val)
		ans += val * pow
		pow *= base
	}
	return ans
}

func DecToBase(dec, base int) string {
	str := ""
	const charset = "0123456789ABCDEF"
	for dec > 0 {
		rem := dec % base
		str = string(charset[rem]) + str
		dec /= base
	}

	return str
}

func BaseToBase(value string, base, newBase int) string {
	dec := BaseToDec(value, base)
	newStr := DecToBase(dec, newBase)
	return newStr
}

func main() {
	fmt.Println(BaseToBase("E", 16, 2))          // "1110"
	fmt.Println(BaseToBase("8831A383B", 12, 16)) // DEADBEEF
	fmt.Println(BaseToBase("E", 16, 2))

}
