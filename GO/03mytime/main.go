package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learning Time")

	// getting the current time
	presentTime := time.Now()

	fmt.Println("current time is: ", presentTime)

	//formatting the time using Format pkg
	formattedTime := presentTime.Format("Mon 01-02-2006 15:04:05")
	fmt.Println("formatted time is: ", formattedTime)

	// create the date using time.Date pkg
	createdDate := time.Date(2025, time.February, 11, 23, 4, 3, 5, time.UTC)
	fmt.Println("created date is: ", createdDate)

	//formatting the created date
	formattedDate := createdDate.Format("Mon 01-02-2006")
	fmt.Println("formatted date is: ", formattedDate)

}
