package main

import "fmt"

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	var name, infor string
	frame := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	for true {
		fmt.Printf("> ")
		fmt.Scan(&name)
		fmt.Scan(&infor)
		if infor == "eat" {
			frame[name].Eat()
		} else if infor == "move" {
			frame[name].Move()
		} else if infor == "speak" {
			frame[name].Speak()
		} else {
			fmt.Println("Try again!")
		}
	}
}
