package main

import "fmt"

func factorial(a float64) float64 {
	if a == 0 {
		return 1
	} else if a == 1 {
		return 1
	} else {
		return a * factorial(a-1)
	}
}

func main() {
	var x float64 = 0
	var y float64 = 0
	var i float64
	var index float64 = 5
	for i = 0; i <= index; i++ {
		//fmt.Println(i)
		x += i
		y = y + x/factorial(i)
	}
	fmt.Printf("Euler number with n = 5 is %f\n", y)
}
