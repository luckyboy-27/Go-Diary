package main

import "fmt"

func main() {
	//way 1
	//var a, b bool
	//a = True
	//b = False
	// can use := for fast declaration
	//var myname string
	myname := "Vinh"
	fmt.Println("Hello, ", myname)
	var num1 int = -2
	var num2 int = 3
	// use "type of variable"() --> convert to another variable
	var x int32
	var y int64
	var z complex128
	//x = y clearly this cause error
	x = int32(y) //no longer error
	fmt.Println("Sum is", num1+num2)
}
