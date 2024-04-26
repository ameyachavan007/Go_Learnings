package main

import (
	"fmt"
	"hello"
	"os"
)

func main() {
	fmt.Printf(hello.Say(os.Args[1:]))
	hello.Example()
	hello.GetAverage()
}
