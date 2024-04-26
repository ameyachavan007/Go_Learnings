package hello

import (
	"fmt"
	"os"
)

func Example() {
	a := 2
	b := 2.1
	/*
	   [1] means we can reuse the 1st paramter again instead of typing it     twice for example fmt.Printf("\na: %8T %v\n", a, a)
	   0th param is Printf
	*/
	fmt.Printf("\na: %8T %[1]v\n", a)
	fmt.Printf("b: %8T %[1]v\n", b)

	// cannot assign a to b without type casting
	a = int(b)
	fmt.Printf("\na: %8T %[1]v\n", a)
}

// TO RUN: go run ./cmd < nums.txt
func GetAverage() {
	var sum float64
	var n int

	for {
		var val float64

		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}
		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}
	fmt.Println("\nThe average is", sum/float64(n))
}
