package main

import "fmt"

func main() {
	numbers := []int8{1, 2, 3, 4, 5}
	fmt.Println(sum(numbers))

	floats := []float64{1.0, 2.5, 3.0}
	fmt.Println(sum(floats))
}

// non-generic

// this function will only work with slices of integers
// if we wanted to work with more types we would need to make more functions
// this is not very DRY
func _(numbers []int) (res int) {
	for _, element := range numbers {
		res += element
	}
	return
}

// generic version

// Number is an interface that accepts union types specified
// this basically says "Number will work with all of these types"
type Number interface {
	int | int8 | int16 | int32 | int64 | float64 | float32
}

// sum then takes the type constraint from Number
// previously we would need to do int | int8 | ... for each function
func sum[T Number](input []T) T {
	var res T
	for _, element := range input {
		res += element
	}
	return res
}
