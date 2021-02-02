package main

import "fmt"

func foo() *int {
	x := 3
	return &x // very interesting
}

func main() {
	var y *int
	y = foo()
	fmt.Printf("%d\n", y)
}
