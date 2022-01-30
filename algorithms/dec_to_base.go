package main

import (
	"fmt"
)

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

func main() {
	fmt.Println(DecToBase(15, 2))
	fmt.Println(DecToBase(10, 2))
	fmt.Println(DecToBase(11, 2))
	fmt.Println(DecToBase(8, 2))
	fmt.Println(DecToBase(31, 2))
	fmt.Println(DecToBase(15, 16))
	fmt.Println(DecToBase(31, 16))
}
