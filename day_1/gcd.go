package main

import "fmt"

func gcd(a int, b int) int {
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	} else {
		if a > b {
			return gcd(b, a%b)
		} else {
			return gcd(a, b%a)
		}
	}
}

func main() {
	fmt.Println("Greatest Common Divisor of 2 numbers is", gcd(11212, 313213))
}
