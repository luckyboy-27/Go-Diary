package main

import "fmt"

func main() {
	//slice is a "window" on an underlying array
	//can increase the size of the slice, while array is fixed
	// pproperties: 1. pointer, point to the first element which contains the slice
	// 2. length is the # of elements in slice
	// 3. capacity is maximum number of element in the slice
	arr := [...]string{"a", "b", "c", "d", "e", "f"}
	//slice literals --> can be used to initialize a slice
	//--> create underlying array ([]) and create a slice reference to it
	//--> slice points to the start of the array, length is capacity
	sli := []{1,2,3,4,5} // the compiler understand this is a slice, so it makes an underlying array for that slice to point 
	s1 := arr[1:3]
	s2 := arr[2:5]
	fmt.Println(arr)
	fmt.Println(sli)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(len(arr))
	fmt.Println(cap(s1))
	fmt.Println(cap(s2))

}
