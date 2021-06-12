package main

import (
	"fmt"
	"math"
)

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

func (t Triangle) Check_right_triangle() string {
	sli := []float64{t.x, t.y, t.z}
	temp := sli[0]
	for i := range sli {
		if temp < sli[i] {
			temp = sli[i]
		}
	}

	if t.x == t.y && t.x == t.z {
		return "Equilateral triangle"
	} else {
		if t.x == temp && t.x*t.x == t.y*t.y+t.z*t.z {
			return "Right triangle"
		} else if t.y == temp && t.y*t.y == t.x*t.x+t.z*t.z {
			return "Right triangle"
		} else if t.z == temp && t.z*t.z == t.y*t.y+t.x*t.x {
			return "Right triangle"
		} else {
			return "Not a right triangle"
		}
	}
}

func main() {
	t := Triangle{3, 3, 3}
	fmt.Println("Area of triangle is:", t.Areas())
	fmt.Println(t.Check_right_triangle())
}
