package main

import "fmt"

// Define a struct User (always keep first letter capital in struct)
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Learning Structs")

	// Utilize the User struct
	soham := User{"soham", "soham@go.dev", true, 25}
	fmt.Println("soham details are:: ", soham) // Result is:: soham details are::  {soham soham@go.dev true 25}

	// Printing the soham details in more detailed using %+v and printf
	fmt.Printf("Soham detailed details are:: %+v\n", soham) // Result is:: {Name:soham Email:soham@go.dev Status:true Age:25}

	// Accessing particular value from struct
	fmt.Printf("User Name is %v and email is %v", soham.Name, soham.Email) // Result is:: User Name is soham and email is soham@go.dev 

}
