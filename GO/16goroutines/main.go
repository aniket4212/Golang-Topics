package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learning Goroutines and concurrency...")
	// call greeting function with goroutine
	go greeting("Hello") // go keyword indicates that this is a goroutine
	greeting("World")

}

// creating a function which is printing something 5 times

func greeting(s string) {
	for i := 0; i < 6; i++ {
		// sleep for a second
		time.Sleep(3 * time.Millisecond)  // sllep waits for goroutine to be complete
		fmt.Println(s)
	}
}
