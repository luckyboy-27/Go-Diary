package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		x := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				x++
			}
		}
		if x == 1 {
			fmt.Printf("%d is prime\n", i)
		}
	}
}
