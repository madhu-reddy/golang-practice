package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in golang")
	hitesh := User{"Hitesh", "hitesh@go.dev", true, 30}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)                        // hitesh details are: {Name:Hitesh Email:hitesh@go.dev Status:true Age:30}
	fmt.Printf("Name is %v and email is %v.\n", hitesh.Name, hitesh.Email) // Name is Hitesh and email is hitesh@go.dev.
	hitesh.GetStatus()
	hitesh.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", hitesh.Name, hitesh.Email) // doesn't change the original email
	hitesh.NewMailPointer()
	fmt.Printf("Name is %v and email is %v.\n", hitesh.Name, hitesh.Email) // changes the original email.
}

func (u User) GetStatus() {
	fmt.Println("is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

func (u *User) NewMailPointer() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
