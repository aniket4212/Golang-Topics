package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to conversion input here: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("your input value is: ", input)

	//converting the string into the float value to add the values in it.(using Parsefloat) and use Trimspace to handle \n condition.
	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("your num rating is: ", numrating+1)
	}

}
