package main

import "fmt"

func GenDisplaceFn(a, v, s_0 float64) func(float64) float64 {
	return func(t float64) float64 {
		s := (1/2)*a*t*t + v*t + s_0
		return s
	}
}

func main() {
	var a, v, s_0, t float64
	fmt.Printf("Enter aceleration: ")
	fmt.Scan(&a)
	fmt.Printf("Enter initial velocity: ")
	fmt.Scan(&v)
	fmt.Printf("Enter initial displacement: ")
	fmt.Scan(&s_0)
	fmt.Printf("Enter time: ")
	fmt.Scan(&t)
	fn := GenDisplaceFn(a, v, s_0)
	fmt.Println(fn(t))
}
