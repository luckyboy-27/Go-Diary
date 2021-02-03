package main

import "fmt"

func main() {
	//"make" can be used to make a slice --> another way to make slice --> init slice in particular size
	// --> using make --> len = cap (slice)
	sli := make([]int, 10)      //10 is length (or capacity) --> it is the same
	sli1 := make([]int, 10, 15) // len and cap are separately different, we can define len and cap on our own. Len = 10, cap = 15
	// add stuffs to slice using append function, increase size if necessary, if the size is bigger than the array, it will create a new underlying array
	// slice is very comfortable, like list in python
	//example
	sli = append(sli, 100)
	sli1 = append(sli1, 20) // size is increase, to accomondate number we put in
	fmt.Println(sli)
	fmt.Println(sli1)
}
