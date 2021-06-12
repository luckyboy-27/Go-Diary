package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
	getName()
}

type Cow struct {
	name, food, locomotion, noise string
}

type Bird struct {
	name, food, locomotion, noise string
}

type Snake struct {
	name, food, locomotion, noise string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func (c Cow) getName() {
	fmt.Println(c.name)
}

func (b Bird) getName() {
	fmt.Println(b.name)
}

func (s Snake) getName() {
	fmt.Println(s.name)
}

func main() {
	var order string
	frame := map[string]Animal{
		name :   Cow{food: "grass", locomotion: "walk", noise: "moo"},
		name :  Bird{food: "worms", locomotion: "fly", noise: "peep"},
		name : Snake{food: "mice", locomotion: "slither", noise: "hsss"},
	}
	for true {
		fmt.Print(">")
		fmt.Scan(&order)
		fmt.Scan(&name)
		switch order {
		case "newanimal":
			var name, infor string
			var new_animal Animal
			fmt.Scan(&infor)
			if infor == "cow" {
				new_animal = Cow{name: name}
			} else if infor == "bird" {
				new_animal = Bird{name: name}
			} else {
				new_animal = Snake{name: name}
			}
			new_animal.getName()
			fmt.Println("Created it!")
		}
		case "query":
			var action string
			fmt.Scan(&action)
			if action == "eat" {
				frame[name].Eat()
			} else if action == "move" {
				frame[name].Move()
			} else {
				frame[name].Speak()
			}
		default:
			fmt.Println("Invalid input")

		}
	}
}
