package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Learning the channels and deadlock....")
	// Channels are a way in which multiple goroutines are talk to each other
	// We can pass value in channel only if someone wants to access or someone reads from that channel(they are working only with goroutines)

	//create the channel of type int
	mychan := make(chan int) // this is unbuffered channel

	// mychan := make(chan int, 2)  // this is buffered channel with 2 buffer

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		// Accessing and printing value from channel
		fmt.Println(<-mychan) // here  <- sending value outside the channnel
		fmt.Println(<-mychan)
		wg.Done()
	}(mychan, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		// insert value  in channel
		mychan <- 15 // here  <- sending value inside channel
		mychan <- 20
		close(mychan) // Closing the channel
		wg.Done()
	}(mychan, wg)

	wg.Wait()

}
