package main

import "fmt"

type Point struct {
	X, Y float64
}

func (p *Point) Increase(dx, dy float64) {
	// if we did not use pointer, it will only change value in Increase func not in main func --> using pass by refference.
	p.X = p.X + dx
	p.Y = p.Y + dy
	//fmt.Println(p.X, p.Y)
}

//method with which receiver is a pointer
func main() {
	p := Point{3, 4}
	fmt.Println("Point p = ", p)

	p.Increase(7, 9)
	fmt.Println("After increasing, p = ", p)

	// We can write a method that receives a ptr and changes the value in the address that it points to
	//For example, a point in 2D coordinate x, y --> increase x, y by adding dx and dy
	// We need to use ptr to point to the address of x and y --> it can change value of x and y.
}
