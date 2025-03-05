package main

import "fmt"

func main() {
	fmt.Println("Learning arrays")
	// array is of fixed size and we can not add elements into the array

	// declare the array of size 4
	var fruitarray [4]string

	//adding the elements in array w.r.t index
	fruitarray[0] = "apple"
	fruitarray[1] = "pineapple"
	fruitarray[3] = "mango"
	// index 2 is printed as whitespace
	fmt.Println("fruit array is: ", fruitarray) // output is fruit array is:  [apple pineapple  mango]

	//print the length and capacity of array
	fmt.Println("length of array is: ", len(fruitarray))   // length of array is:  4
	fmt.Println("capacity of array is: ", cap(fruitarray)) // capacity of array is:  4

	// declare and initialize the array in one line
	var vegList = [3]string{"potato", "beans", "tomato"}

	fmt.Println("vegList is: ", vegList)

}
