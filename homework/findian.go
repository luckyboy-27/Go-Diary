package main

import (
	"fmt"
	"strings"
)

func findIan(a string) string {
	x := 0
	for i := 0; i < len(a); i++ {
		if a[i] == 97 {
			x++
		}
	}
	if a[0] == 105 && a[len(a)-1] == 110 && x > 1 {
		return "Found!"
	} else if strings.Compare(a, "ian") == 0 {
		return "Found!"
	}
	return "Not Found!"
}

func main() {
	var ian string
	fmt.Printf("Enter a string: ")
	fmt.Scan(&ian)
	// i = 105, n = 110, a = 97
	fmt.Println(findIan(strings.ToLower(ian)))
}
