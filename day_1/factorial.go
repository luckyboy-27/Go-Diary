package main

import "fmt"

func factorial(a int) int {
	if a == 0 {
		return 1
	} else if a == 1 {
		return 1
	} else {
		return a * factorial(a-1)
	}
}

func main() {
	fmt.Println("Factorial of 5 is", factorial(5))
}
