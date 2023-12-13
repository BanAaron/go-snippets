package main

import "fmt"

func main() {
	byValue()
	byReference()
}

// byValue demonstrates that assignments are passed byValue and not be reference
func byValue() {
	var a []int
	b := a
	// appending to be does not update a because b contains the value of a and
	// is not a reference to it
	b = append(a, 1)
	fmt.Println(a, b)
}

// byReference shows that using pointers we can update a by reference
func byReference() {
	var a []int
	// b is a pointer to a
	b := &a
	// then when we append to b, a will be mutated
	*b = append(*b, 1)
	fmt.Println(a, *b)
}
