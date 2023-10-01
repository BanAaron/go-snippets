package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)
	go count(10, channel)
	go count(10, channel)
	// we can iterate over a channel
	for v := range channel {
		fmt.Println(v)
	}
}

func count(number int, channel chan int) {
	for i := 0; i < number; i++ {
		time.Sleep(time.Second / 2)
		channel <- i + 1
	}
	// we close like this
	// when the channel is closed the range loop with also close
	close(channel)
}
