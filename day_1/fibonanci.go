package main

import "fmt"

func fibonanci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonanci(n-1) + fibonanci(n-2)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fibonanci(i))
	}
	x, y := 0, 1
	var z int
	for i := 0; i < 10; i++ {
		x = y
		y = z
		z = x + y
		fmt.Println(z)
	}
	a, b := 1, 1
	for i := 0; i < 10; i++ {
		a, b = b, a+b
		fmt.Println(b)
	}
}
