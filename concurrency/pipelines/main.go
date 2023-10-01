package main

import "fmt"

func main() {
	nums := []int{69, 420, 42069, 69420}
	// stage 1
	dataChannel := sliceToChannel(nums)
	// stage 2
	finalChannel := sq(dataChannel)
	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
