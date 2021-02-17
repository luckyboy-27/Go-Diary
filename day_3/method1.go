package main

import (
	"fmt"
	"math"
)

//method with same name but different parameters
type Square struct {
	X float64
}

func (s Square) Areas() float64 {
	return s.X * s.X
}

type Rectangle struct {
	X, Y float64
}

func (r Rectangle) Areas() float64 {
	return r.X * r.Y
}

type Triangle struct {
	x, y, z float64
}

func (t Triangle) Areas() float64 {
	if t.x+t.y < t.z || t.y+t.z < t.x || t.x+t.z < t.y {
		fmt.Println("Not a triangle")
		return 0
	} else {
		p := (t.x + t.y + t.z) / 2
		return math.Sqrt(p * (p - t.x) * (p - t.y) * (p - t.z))
	}
}

func main() {
	s := Square{4}
	r := Rectangle{3, 4}
	t := Triangle{3, 4, 5}

	fmt.Println("Area of square is: ", s.Areas())
	fmt.Println("Area of rectangle is: ", r.Areas())
	fmt.Println("Area of triangle is: ", t.Areas())
	// note: the name of function is the same ("Areas"), it must have the different receiver (different struct)
	// because the method can be different with different data types but function itself must be named separately.x
}
