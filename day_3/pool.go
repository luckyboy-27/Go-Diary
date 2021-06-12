package main

type Shape interface {
	area() float64
}

type Triangle struct {
	height, width float64
}

func (t Triangle) area() {
	return (t.height + t.width) / 2
}

func FitIn(s Shape) bool {
	if s.area() > 100 {
		return true
	} else {
		return false
	}
}

func main() {
	var s Shape
	t := Triangle{3, 4}
	FitIn(t.area())

}
