package main

import (
	"fmt"
	"strconv"
)

func sortSlice(sli []int, size int) []int {
	//using buble sort
	var temp int
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if sli[i] > sli[j] {
				temp = sli[i]
				sli[i] = sli[j]
				sli[j] = temp
			}
		}
	}
	return sli
}

func main() {
	var text string //number to start the program
	size := 3
	sli := make([]int, size)
	for true {
		fmt.Printf("Enter an integer and i will put it in a sorted slice, or just enter x to get out of the program: ")
		fmt.Scan(&text)
		if text == "X" || text == "x" {
			break
		} else {
			num, e := strconv.Atoi(text)
			if e == nil {
				sli = append(sli, num)
				size++
				fmt.Println("Your new slice: ", sortSlice(sli, size))
			}
		}
	}
}
