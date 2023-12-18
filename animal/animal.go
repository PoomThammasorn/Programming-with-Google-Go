package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	category string
	food     string
	move     string
	sound    string
}

func (a *Animal) Eat() {
	fmt.Println(a.category + " eats " + a.food + ".")
}

func (a *Animal) Move() {
	fmt.Println(a.category + " moves by " + a.move + "ing.")
}

func (a *Animal) Speak() {
	fmt.Println(a.category + " speaks " + a.sound + ".")
}

func main() {
	cow := Animal{"Cow", "grass", "walk", "moo"}
	bird := Animal{"Bird", "worms", "fly", "peep"}
	snake := Animal{"Snake", "mice", "slither", "hsss"}
	var input Animal
	for {
		fmt.Print("> ")
		var animal, info string
		fmt.Scan(&animal, &info)
		animal = strings.ToLower(animal)
		info = strings.ToLower(info)

		switch := animal {
		case "cow":
			input = cow
		case "bird":
			input = bird
		case "snake":
			input = snake
		default:
			fmt.Println("Invalid animal name.")
			continue
		}

		switch info {
		case "eat":
			input.Eat()
		case "move":
			input.Move()
		case "speak":
			input.Speak()
		default:
			fmt.Println("Invalid info request.")
			continue
		}
	}
}
