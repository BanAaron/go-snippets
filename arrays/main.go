package main

import (
	"fmt"
	"slices"
)

func main() {
	// create a slice of integers
	numbers := []int{1, 2, 3}
	// arrays have a length
	names := [3]string{"aaron", "chris", "drew"}
	// we can get the min and max of a slice
	fmt.Println(slices.Max(numbers))
	// the same cannot be done with arrays
	fmt.Println(names)
}
