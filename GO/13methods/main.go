package main

import "fmt"

// Define a struct User (always keep first letter capital in struct)
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// Create a method GetStatus which receives a User struct
func (u User) GetStatus() {
	fmt.Println("User status is: ", u.Status)
}

// update the actual status of user by using pointer
func (u *User) UpdateStatus() {
	u.Status = false // make status of user false
}

func main() {
	fmt.Println("Learning Methods")

	// Utilize the User struct
	soham := User{"soham", "soham@go.dev", true, 25}
	fmt.Println("soham details are:: ", soham) // Result is:: soham details are::  {soham soham@go.dev true 25}

	// Printing the soham details in more detailed using %+v and printf
	fmt.Printf("Soham detailed details are:: %+v\n", soham) // Result is:: {Name:soham Email:soham@go.dev Status:true Age:25}

	// Accessing particular value from struct
	fmt.Printf("User Name is %v and email is %v\n", soham.Name, soham.Email) // Result is:: User Name is soham and email is soham@go.dev

	// Accessing the users Status using method
	soham.GetStatus() // result is :: User status is: true

	// Accessing Updatestatus which is update the status of user
	soham.UpdateStatus()

	// Accessing the status after updating it.
	soham.GetStatus() // result is :: User status is: false
}
