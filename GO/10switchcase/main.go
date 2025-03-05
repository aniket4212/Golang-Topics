package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Learning switch case statement")

	// Get random number every single time
	rand.Seed(time.Now().UnixNano())
	// Gettting the diceNumber beetween 1 to 6
	diceNumber := rand.Intn(6)+ 1
	fmt.Println("DiceNumber is:: ", diceNumber)

	// Use the switch case to print the random dice number(like Ludo)
	switch diceNumber{
	case 1:
		fmt.Println("diceNumber is 1 and you are open to play")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough  // fallthrough keyword will excute the next line also in this case case 5 also be executed
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll dice again")
	default:
		fmt.Println("what is that!")

	}

}
