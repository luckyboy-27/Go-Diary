package main

import "fmt"

func main() {
	var fruit = [...]string{
		"apricot",
		"banana",
		"pineapple",
		"apple",
		"persimmon",
		"pear",
		"blueberry",
	}
	var temp string
	for i := 1; i < len(fruit); i++ {
		for j := 0; j < len(fruit); j++ {
			if fruit[i] < fruit[j] {
				temp = fruit[i]
				fruit[i] = fruit[j]
				fruit[j] = temp
			}
		}
	}
	fmt.Println(fruit)
}
