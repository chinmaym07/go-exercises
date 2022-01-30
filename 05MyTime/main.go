package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println("Present Time", presentTime)

	fmt.Println("Formatted Present Time ", presentTime.Format("01-02-2006 Monday"))

	createdDate := time.Date(2021, time.September, 25, 23, 24, 0, 0, time.UTC)

	fmt.Println("Created Time", createdDate.Format("01-02-2006 Monday"))
}
