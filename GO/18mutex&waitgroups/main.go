package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Learning mutex and waitgroups with Race Condition...")

	// create waitgroup globally
	wg := &sync.WaitGroup{}
	
	// Create mutex globally
	mut := &sync.Mutex{}  

	// There are also RWMutex for read and write data Basically lock the memory when we want to write something bcz read is faster
	rwmut := &sync.RWMutex{}


	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()  // this will lock memory
		score = append(score, 1)
		mut.Unlock()  // this will unlock memory
		wg.Done()   // this will decrease or marked the goroutine as finished 
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		rwmut.Lock()  // use rwmutex for lock memory becuase we just read from the memory
		fmt.Println(score)
		rwmut.Unlock() // this will unlock memory
		wg.Done()
	}(wg, rwmut)

	wg.Wait()
	fmt.Println(score)

}
