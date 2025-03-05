package main

import "fmt"

func main() {
	// pointer will work on actual value present in specific memory address not on the copy of that value.
	fmt.Println("learning the pointers")

	// declare a pointer
	// var ptr *int  // ptr is a pointer that holds integer type values
	// fmt.Println("pointer value: ", ptr)

	myNumber := 23

	var ptr = &myNumber // here ptr is referencing to the myNumbers address

	fmt.Println("ptr value is: ", ptr) // this will print the memory address

	// print by dereferencing the ptr by add * to front 
	fmt.Println("ptr value is: ", *ptr) // this will print the actual value present at that address


	// get the actual ptr value by *ptr and multiply it by 2 and check for result 
	*ptr =  *ptr * 2 
	fmt.Println("new value is: ", myNumber) 


}
