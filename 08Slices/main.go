package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")
	var fruitList = []string{"Apple", "Orange", "Banana"}
	fmt.Printf("Type of fruitlist %T\n", fruitList)
	fmt.Println("Old Fruitlist ", fruitList)
	fruitList = append(fruitList, "Watermelon", "Cheekoo", "Mango", "Strawberry")
	fmt.Println("Updated Fruitlist ", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println("Updated Fruitlist ", fruitList)

	highScores := make([]int, 4)
	fmt.Printf("Type of highScores %T\n", highScores)
	fmt.Printf("len of highScores %d\n", len(highScores))

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 456
	highScores[3] = 867
	highScores = append(highScores, 555, 666, 345)

	fmt.Println("High Scores ", highScores)
	fmt.Printf("len of highScores %d\n", len(highScores))

	fmt.Println("Is High Scores slice sorted", sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println("Sorted High Scores ", highScores)
	fmt.Println("Is High Scores slice sorted", sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index

	courseList := []string{"ReactJs", "Javascript", "Ruby", "Swift", "Python"}

	fmt.Println("Courses", courseList)

	index := 2
	courseList = append(courseList[:index], courseList[index+1:]...)

	fmt.Println("Updated Courses", courseList)
}
