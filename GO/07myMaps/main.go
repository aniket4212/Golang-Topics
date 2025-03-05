package main

import "fmt"

func main() {
	fmt.Println("Learning Maps")

	// Creating a Map of type key string value also string
	languages := make(map[string]string)

	// Set the values in map using key
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	// prints the whole map
	fmt.Println("Map elements are: ", languages) //Result is:: Map elements are:  map[GO:Golang JS:Javascript PY:Python]

	// Accessing particular element from map on the basis of key
	fmt.Println("JS stands for::", languages["JS"]) // Result is::  JS stands for:: Javascript
	fmt.Println("PY stands for::", languages["PY"]) // Result is::  PY stands for:: Python

	// print the Map using the for loop
	for key, value := range languages {
		fmt.Printf("For key %v the value is %v\n", key, value) // this will print the maps element key and value one by one
	}

	// DELETE the particular value from map
	delete(languages, "GO")                                 // delete the key "GO" from languages map
	fmt.Println("latest map after deletion:: ", languages) // Result is::  latest map::  map[JS:Javascript PY:Python]

}
