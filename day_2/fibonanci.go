package main

import "fmt"

// write fibonanci by slice

func main() {
	slice1 := []int{1, 2}         //can increase automatically
	slice2 := make([]int, 11, 20) //have to arrange reasonable size
	n := 10

	for i := 2; i <= n; i++ {
		slice1 = append(slice1, slice1[i-1]+slice1[i-2])
	}
	fmt.Println(slice1)

	slice2[0], slice2[1] = 1, 2
	for i := 2; i <= n; i++ {
		slice2[i] = slice2[i-1] + slice2[i-2] //like an array
	}
	fmt.Println(slice2)
}
