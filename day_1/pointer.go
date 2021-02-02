package main

import "fmt"

func main() {
	var x int = 1
	var y int
	var p *int

	p = &x
	x = 3
	//p = address of x, if we change x -> p, y changes
	//*p = y --> only y changes, p and x is the same
	y = *p

	ptr := new(int) // --> new function creates a variable but return a pointer points to that variable
	*ptr = 1
	x = *ptr // x, and p and ptr changes, y is stable

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*p)
	fmt.Println(*ptr)
}
