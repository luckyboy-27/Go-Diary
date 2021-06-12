package main

import "fmt"

func main() {
	defer fmt.Println("World")
	// operation will be pushed into stack and activate when the around operations end.
	// using defer to "release resources"
	fmt.Println("Hello")
	i := 1
	defer fmt.Println(i + 1) // stil 2, not change to 3
	i++
	fmt.Println("stop")
	//defer fmt.Println("Stop!")
}
