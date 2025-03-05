package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Learn the slices")

	// Declare and intialize the fruitlist
	var fruitlist = []string{"apple", "banana", "grapes"}

	//print the type of fruitlist
	fmt.Printf("Type of fruitlist is %T\n", fruitlist) // result is: Type of fruitlist is []string

	// Now append the elements to the fruilist slice using append
	fruitlist = append(fruitlist, "chiku", "amla")

	fmt.Println("fruitlist is: ", fruitlist)

	// Gettting the specific element from slice
	fruitlist = fruitlist[1:3] // this will start from index 1 upto the index 2 it will not consider index 3(index 3 is exclusive)
	fmt.Println(fruitlist)

	fruitlist = fruitlist[:4] // this will start from starting point which is index 0 upto the index 3 it will not consider index 4(index 4 is exclusive)
	fmt.Println(fruitlist)

	// Declare highscore of size 4  using make func
	highScore := make([]int, 4)
	fmt.Printf("type before is %T\n", highScore) // result is type before is []int

	highScore[0] = 275
	highScore[1] = 987
	highScore[2] = 345
	highScore[3] = 856

	highScore = append(highScore, 555, 777, 888) // append can realocate the extra memory to the slice
	fmt.Println("highScore elements are: ", highScore)
	fmt.Printf("type after is: %T\n", highScore) // result is type after is []int

	// Sort the slice in asc order
	sort.Ints(highScore)
	fmt.Println("sorted slice is: ", highScore) // result is sorted slice is:  [275 345 555 777 856 888 987]

	// check whether slice is sorted or not in bool
	fmt.Println(sort.IntsAreSorted(highScore)) // result is true

	fmt.Println("=============================")

	// How to remove values from Slice based on index
	var course = []string{"react", "javascript", "python", "golang", "rubby"}

	// consider we have to delete index 2 that is python
	var index int = 2

	course = append(course[:index], course[index+1:]...) // this will exclude the index 2 and continue for next
	fmt.Println("course after deletion is: ", course) // result is course after deletion is:  [react javascript golang rubby]

}
