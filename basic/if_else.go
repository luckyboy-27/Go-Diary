package main

import "fmt"

func main() {
	x := 28
	if x%3 == 0 {
		fmt.Printf("%d is a multiple of 3\n", x)
	} else {
		fmt.Printf("%d is not a multiple of 3\n", x)
	}
}
