package main

import "fmt"

func main() {
	num := 27
	var user int

	for true {
		fmt.Printf("Guess a number: ")
		fmt.Scan(&user)
		if user > num {
			fmt.Println("It is smaller than that!")
			fmt.Printf("\n")
		} else if user < num {
			fmt.Println("It is bigger than that!")
			fmt.Printf("\n")
		} else {
			fmt.Println("You are correct!")
			break
		}
	}
}
