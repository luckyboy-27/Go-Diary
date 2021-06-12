package main

import "fmt"

type member struct {
	name string
	id   int
}

func main() {
	// a is an array
	var a = [...]int{1, 2, 3}
	// b is a pointer to array a
	var b = &a
	// Print somes values in a and b
	fmt.Println(a[0], a[1])
	//print b according to a, and also change b inside loop.
	for index, value := range b {
		// changes element in b
		b[index] += 1
		fmt.Println("number", value, "at", "position", index)
	}
	//fmt.Println(a)
	//fmt.Println(*b)
	// a changes because b changes. --> print a below...
	fmt.Printf("\n")
	for index, value := range a {
		fmt.Println("number", value, "at", "position", index)
	}
	// array of a struct
	mem := [...]member{{name: "David", id: 1}, {name: "Vinh", id: 2}}
	fmt.Println(mem)
}
