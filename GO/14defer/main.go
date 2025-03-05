package main

import "fmt"

func main() {

	// defer will execute at the end of function (imagine that defer statement tooks the line at last of function)
	// [ i.e. it work on principle last in first out ]
	defer fmt.Println("world") // this is 5th line to be executed
	defer fmt.Println("one")   // this is 4th line to be executed
	defer fmt.Println("Two")   // this is 3rd line to be executed
	fmt.Println("Hello")       // this is 1st line to be executed
	myDefer()                  // this is the 2nd line to be executed
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// result is ::
// Hello
// 4
// 3
// 2
// 1
// 0
// Two
// one
// world
