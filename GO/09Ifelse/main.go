package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Learning If Else")

	// Creating reader to get input from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter LoginCount: ")

	// get loginCount as input from user as a string (reader.Readstring function returns a string only)
	input, _ := reader.ReadString('\n')

	// Trim newline characters
	input = strings.TrimSpace(input)

	// convert the input string value into int value as loginCount
	loginCount, _ := strconv.Atoi(input)

	fmt.Println("Entered LoginCount is:: ", loginCount)
	var result string

	// Uses the if else statement to classify the user on the basis of loginCount
	if loginCount <= 10 {
		result = "Regular user"
	} else if loginCount > 10 && loginCount < 15 {
		result = "Imp user"
	} else {
		result = "VIMP user"
	}
	fmt.Println("Result is:: ", result)

}
