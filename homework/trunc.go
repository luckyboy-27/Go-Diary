package main

import "fmt"

func trunc(a float64) int64 {
	return int64(a)
}
func main() {
	var a float64
	fmt.Printf("Enter a float number: ")
	fmt.Scan(&a)
	fmt.Printf("Number: %d\n", trunc(a))

}
