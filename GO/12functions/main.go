package main

import "fmt"

func main() {

	fmt.Println("Learning Functions")

	// call adder function
	result1 := adder(4, 6)
	fmt.Println("result1 is:: ", result1)

	// call proAdder function which receives infinite parameter
	proresult := proAdder(3, 4, 5, 6)
	fmt.Println("proresult is:: ", proresult)
}

// write the function to add two numbers
func adder(val1, val2 int) int {
	return val1 + val2
}

// writing proAdder fuction which receives number of arguments(infinite parameters)
func proAdder(values ...int) int { // (... means infinite)
	// declare result as 0
	result := 0

	// range through the vlaues sclice which we receive in function parameter
	for _, value := range values {
		result += value // add value in result one by one
	}

	return result
}

// we can also return multiple return values in function