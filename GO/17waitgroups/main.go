package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Create a waitgroup globally
var wg sync.WaitGroup  // Usually this are pointers 

// There are three methods in waitgroupd 
 // 1 = wg.Add()    this will add a goroutine counter 
 // 2 = wg.Wait()   this will wait until goroutine finish
 // 3 = wg.Done     this will decrease the goroutine counter when goroutine finish

//================ Brief about waitgroups==================
// A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls [WaitGroup.Add] to set the number of goroutines to wait for. Then each of the goroutines runs and calls [WaitGroup.Done] when finished. At the same time, [WaitGroup.Wait] can be used to block until all goroutines have finished.

// declare the mutex Globally
var mut sync.Mutex  // Usually this are pointer
// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
// if one operation is in process at specific memory address and second one comes then mutex will lock the memory address until first one is completed.

var signals = []string{"test"}

func main() {
	fmt.Println("Learning Goroutines with waitgroups and Mutex...")

	// creating the website list
	websiteList := []string{
		"https://Lco.com",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.dev",
	}

	// interate through websiteList and call getstatuscode
	for _, web := range websiteList {
		
		go getStatusCode(web)
		wg.Add(1)   //this will add a one goroutine in wg
		
	}
	wg.Wait()   // this will wait for one goroutine to finish

}

// creating function which gets the statuscode of the provided endpoint or website
func getStatusCode(endpoint string) {
	defer wg.Done()  // this will decrease the count of goroutine in wg 
	res, err := http.Get(endpoint) // inbuilt Get will return statuscode and err

	if err != nil {
		fmt.Println("OOPs in endpoint")
	} else {
		mut.Lock()   // Here we lock the memory address for performing one operation at a time 
		signals = append(signals, endpoint)
		mut.Unlock() // Here we unlock the memory address
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

	}

}
