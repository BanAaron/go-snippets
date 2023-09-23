package main

import (
	"fmt"
)

func main() {
	// create a channel like this
	channel := make(chan int)
	// call a go routine, passing it the channel
	go doubler(4, channel)
	go doubler(8, channel)
	// we can then read from the channel
	a, b := <-channel, <-channel
	fmt.Println(a, b)
}

func doubler(number int, channel chan int) {
	channel <- number * 2
}
