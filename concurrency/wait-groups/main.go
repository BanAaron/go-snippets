package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// create a wait group
	wg := sync.WaitGroup{}
	// then add to the wait group and kick off a goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go count(i, 10, &wg)
	}
	// we then wait for each job in the group to complete
	wg.Wait()
}

// count takes a pointer to a wait group as an argument
func count(id int, number int, wg *sync.WaitGroup) {
	for i := 0; i < number; i++ {
		// sleep to simulate an intensive process
		time.Sleep(time.Second / 2)
		fmt.Println(id, i+1)
	}
	// use Done() when the activity for the job is done
	wg.Done()
}
