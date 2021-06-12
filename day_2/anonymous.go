package main

import "fmt"

func add(a, b int) int {
	return a + b

}

//anonymous
var addition = func(a, b int) int {
	return a + b
}

// multiple values to return
func Swap(a, b int) (int, int) {
	return b, a
}

// using ..., passing multiple arguments all in 1, using just like a slice.
// arr is a slice.
func sum(a int, arr ...int) int {
	for _, v := range arr {
		a += v
	}
	return a
}

func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main() {
	arr := []int{2, 4, 8, 1, 5}
	fmt.Println(addition(1, 2))
	fmt.Println(add(1, 2))
	fmt.Println(sum(2, 3))
	fmt.Println(getMax(10, 2, 6, 5)) // ... works on multiple variables to pass in function.
	fmt.Println(getMax(arr...))      // using ... suffix to pass a slice/array.
}
