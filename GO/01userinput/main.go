package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")

	// create a reader with inbuild packages 
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating: ")

	// readstring until the ends and ignore err with _ (comma, ok or comma error syntax)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)


}
