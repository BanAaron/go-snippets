package main

import (
	"fmt"
)

func main() {
	// the second argument is the buffer size
	channel := make(chan int, 2)
	// we can then store up to 2 values at a time
	channel <- 8
	channel <- 9
	// and then read out FIFO style
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	// we can add and take away in any order, as long as we do not go out of bounds
	channel <- 1
	channel <- 2
	fmt.Println(<-channel)
	channel <- 1
	fmt.Println(<-channel)
}
