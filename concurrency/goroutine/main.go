package main

import (
	"fmt"
	"time"
)

func main() {
	// to call a go routine all we need to do is put "go" in front of the
	// function call
	go greeter("aaron")
	greeter("steve")
}

func greeter(name string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Hello, %s\n", name)
	}
}
