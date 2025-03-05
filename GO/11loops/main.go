package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning Loops..")

	days := []string{"sun", "mon", "tue", "wed", "thur", "fri", "sat"}
	fmt.Println(days)

	// Loop Type 1
	// for i:= 0; i< len(days); i++{
	// 	fmt.Println(days[i])
	// }

	// Loop Type 2
	for i := range days {
		fmt.Println(days[i])
	}

	// Loop Type 3
	for index, day := range days { // we can also ignore the index by usng underscore(_) instead of index
		fmt.Printf("index is %v and value is %v\n", index, day)

	}

	rougueValue := 1

	for rougueValue < 10 {

		// use break to exit the loop on the basis of condition
		// if rougueValue == 5 {
		// 	break // break will exit from for loop
		// }

		// use continue to escape the value on the basis of condition
		if rougueValue == 6 {
			rougueValue++
			continue // continue will escape the value 6
		}

		// Use goto to jump to the different statement like below

		if rougueValue == 2 {
			goto lco  // goto keyword will jump to the lco statement which is defined below 
		}

		fmt.Println("value is:: ", rougueValue)
		rougueValue++
	}
// define this lco statement to access it using goto keyword
lco:
	fmt.Println("Jumping to learn code online")
}
