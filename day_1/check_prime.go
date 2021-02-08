package main

import (
	"fmt"
)

func main() {
	var num int
	x := 0
	fmt.Printf("Enter a number: ")
	fmt.Scan(&num)
	for i := 1; i < num; i++ {
		if num%i == 0 {
			x++
		}
	}
	if x == 1 {
		fmt.Printf("%d is prime\n", num)
	} else {
		fmt.Printf("%d is not prime\n", num)
	}
}
