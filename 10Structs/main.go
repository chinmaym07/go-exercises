package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in Golang")
	// no inheritence in golang , no super or parent
	user1 := User{"Test", "testuser@gmail.com", true, 18}
	fmt.Println("Test User ", user1)
	fmt.Printf("Test User Details are %+v\n", user1)

	fmt.Printf("Test User Name: %v & email is %v \n", user1.Name, user1.Email)
}
