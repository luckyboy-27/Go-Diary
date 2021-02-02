package main

import "fmt"

func main() {
	//unicode package, string package
	//Compare(a, b) return 1 or -1
	//Contain (s, substring)
	//HasPrefix(s, prefix)
	//Index(s, substring)
	// string immutable, modified, returned
	//Replace(s, old, new, n)
	//TrimSpace(s) --> remove space
	//strconv package: Atoi(s) --> convert string to int, Itoi(s) --> conver int to string
	// can set many values for the constant data types
	const x = 1
	// use iota to represent set of const with difference
	const (
		y = 2
		z = "Hello"
	)
	type Grades int
	const (
		A Grades = iota
		B
		C
		D
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
