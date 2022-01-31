package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
func (u *User) SetStatus(newStatus bool) {
	u.Status = newStatus
}
func (u *User) SetEmail(newEmail string) {
	u.Email = newEmail
}

func main() {
	fmt.Println("Methods in Golang")

	user1 := User{"Test", "testuser@gmail.com", true, 18}
	fmt.Println("Test User ", user1)
	fmt.Printf("Test User Details are %+v\n", user1)

	fmt.Printf("Test User Name: %v & email is %v \n", user1.Name, user1.Email)

	user1.GetStatus()
	user1.SetStatus(false)
	user1.SetEmail("testuser@go.dev")
	fmt.Println("Updated Test User ", user1)
	fmt.Printf("Updated Test User Details are %+v\n", user1)

	fmt.Printf("Updated Test User Name: %v & email is %v \n", user1.Name, user1.Email)
}
