package main

import "fmt"

type students struct {
	name string
	id   int
}

func main() {
	var a students
	//like c, using dot notation to access the struct if it is not pointer
	a.name = "david"
	a.id = 1213
	//assign
	//x := a.name
	//x = "trung"
	// way 2: initialize struct --> using "new" function --> make empty struct
	x := new(students) //zero struct
	fmt.Println(a)
	fmt.Println(x)

}
