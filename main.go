package main

import (
	"fmt"
	"github.com/banaaron/gotils"
)

func main() {
	res := gotils.LeftPad("Hello, World!", 4)
	var err error = nil
	gotils.HandleError(err)
	fmt.Printf(res)
}
