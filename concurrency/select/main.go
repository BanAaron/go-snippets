package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan string)
	b := make(chan string)

	go func() {
		a <- "hello"
	}()

	go func() {
		b <- "world"
	}()

	// we sleep for 100ms to make sure both channels are used
	time.Sleep(time.Millisecond * 100)

	// the select statement will wait on either channel
	// if both are ready it will choose one at random
	select {
	case messageA := <-a:
		fmt.Printf(messageA)
	case messageB := <-b:
		fmt.Printf(messageB)
	}
}
