package main

import "fmt"

//create a slice, try accessing the size, range of slice.
func main() {
	slice1 := []int{1, 2, 344, 3, 32, 21, 32}
	slice2 := slice1[2:5] // if we access out of the range --> error --> can not access freely like C
	slice2[1] = 10
	fmt.Println("SL2:", slice2)
	fmt.Println("SL1:", slice1)
}
