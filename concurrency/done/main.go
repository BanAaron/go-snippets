package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("doing work")
		}
	}
}

func main() {
	done := make(chan bool)
	// create an infinite go routine
	go doWork(done)

	time.Sleep(time.Second)
	close(done)
}
