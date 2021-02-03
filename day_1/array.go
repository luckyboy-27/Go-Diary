package main

import "fmt"

func main() {
	// array has fixed length series of elements of chosen type
	var x = [5]int{1, 2, 3, 4, 5}
	y := [5]int{1, 2, 3, 4, 5}
	k := y[1:4]
	var z [5]int
	//iteration --> print
	for i := 0; i < 5; i++ {
		z[i] = x[i] + y[i]
	}
	for j, v := range y {
		fmt.Printf("index %d, val %d\n", j, v)
	}
	fmt.Println(z[1:3])
	fmt.Println(k)
}
