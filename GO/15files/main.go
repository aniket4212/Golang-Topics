package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Learn working of files using go..")
	// create the content that will be stored into the file
	content := "this needs to go in a file - Learncode online"

	// Create the file which has name myfirstfile.txt and also checking for error while creating the file.
	file, err := os.Create("./myfirstfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)  // calling these function instead of checking every time for nil err

	// write the above content into the file
	length, err := io.WriteString(file, content) // WriteString inbuild function will return the length of the file and err.
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)  // calling these function instead of checking every time for nil err

	fmt.Println("Length of file is ::", length)

	// Now just close the file reccomendation use the defer for closing the file
	defer file.Close()

	// call the readfile funtion and provide file path to it
	readfile("./myfirstfile.txt")
}

// creating the function for reading the file
func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename) // ReadFile inbuild function will read and return the data in byte format
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)  // calling these function instead of checking every time for nil err

	// prints the data which is in byte format
	fmt.Println("bytedata inside file is::: ", databyte)

	// prints the data in text format
	fmt.Println("Textdata inside the file is:: ", string(databyte))
}

// Create the function for checking nil err
func checkNilError(err error){
	if err != nil{
		panic(err)
	}
}