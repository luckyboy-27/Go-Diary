package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func (p Point) Distance() float64 {
	return float64(math.Sqrt(math.Pow(float64(p.x), 2) + math.Pow(float64(p.y), 2)))
}

// using pointer --> pass by reference.
func (p *Point) Create(a, b int) {
	p.x = a
	p.y = b
}

func (p *Point) Scale(v int) {
	p.x = p.x * v
	p.y = p.y * v
}

func (p Point) PrintMe() {
	fmt.Println(p.x, p.y)
}

// no need pointer because we do not pass anything into PrintX/PrintY func
func (p Point) PrintX() {
	fmt.Println(p.x)
}

func (p Point) PrintY() {
	fmt.Println(p.y)
}

func main() {
	p := Point{1, 2}
	var t Point
	fmt.Println("Distance is:", p.Distance())
	t.Create(5, 3)
	t.Scale(2)
	t.PrintMe()
	t.PrintX()
	t.PrintY()
}
