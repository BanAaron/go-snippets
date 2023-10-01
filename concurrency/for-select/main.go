package main

import "fmt"

func main() {
	// create a buffered channel
	c := make(chan string, 3)
	// dealer our slice
	chars := []string{"a", "b", "c"}
	// loop through slice
	for _, v := range chars {
		// use select to insert into the channel
		select {
		case c <- v:
		}
	}

	// close the channel
	close(c)

	// because the channel is closed, we can loop over it
	for res := range c {
		fmt.Println(res)
	}
}
