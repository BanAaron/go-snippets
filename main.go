package main

import (
	"fmt"
	"github.com/banaaron/gotils"
)

func main() {
	res := gotils.LeftPad("Hello, World!", 4)
	fmt.Println(res)

	var err error = nil
	gotils.HandleError(err)

	m := map[int]string{
		1: "aaron",
		2: "chris",
		3: "drew",
	}
	keys := gotils.OrderedKeys(m)
	fmt.Println(keys)
}
