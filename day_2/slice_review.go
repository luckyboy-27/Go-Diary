package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a[:3]
	c := []int{4, 5, 6, 7}
	// if we change b --> a changes
	//b[0] = 4
	// b is a slice, can be seen as a pointer points to array a, which have len = 2, cap = 3
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("Before appending:", len(b), cap(b))

	// a remains the same, while b changes
	// In case the capacity of slice can have more element, append function will automatically
	// allocate more memory in the slice, while the original array doesnt change.
	b = append(b, c...)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("After appending:", len(b), cap(b))

	// now the slice is in a different address, b changes --> a doesnt changes.
	// because the append function will return new slice.
	b[0] = 8
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//note: append function can not allocate new memory/create new slice if the cap is less than or equal to the inital cap.
}
