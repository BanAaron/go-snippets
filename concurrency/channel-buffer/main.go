package main

import "fmt"

func main() {
	// the second argument is the buffer size
	channel := make(chan int, 2)
	// we can then store up to 2 values at a time
	channel <- 69
	channel <- 420
	close(channel)
	// and then read out FIFO order
	for v := range channel {
		fmt.Println(v)
	}
}
