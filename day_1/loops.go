package main

import "fmt"

func main() {
	x := 0
	y := 0
	for i := 0; i < 10; i++ {
		x += i
		if i%2 == 0 {
			y += i
		}
	}
	fmt.Println("Sum of number from 0 to 10 is", x)
	fmt.Println("Sum of even number from 0 to 10 is", y)

	for j := 0; j < 30; j++ {
		if j%3 == 0 {
			//continue --> in other word, REMOVE the number which mod 3 = 0
			fmt.Println(j)
		}
		//fmt.Println(j)
	}
}
