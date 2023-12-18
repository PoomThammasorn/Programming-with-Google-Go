package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food, move, sound string
}

func (c Cow) Eat() {
	fmt.Println("Cow eats", c.food)
}

func (c Cow) Move() {
	fmt.Println("Cow moves by", c.move)
}

func (c Cow) Speak() {
	fmt.Println("Cow speaks", c.sound)
}

type Bird struct {
	food, move, sound string
}

func (b Bird) Eat() {
	fmt.Println("Bird eats", b.food)
}

func (b Bird) Move() {
	fmt.Println("Bird moves by", b.move)
}

func (b Bird) Speak() {
	fmt.Println("Bird speaks", b.sound)
}

type Snake struct {
	food, move, sound string
}

func (s Snake) Eat() {
	fmt.Println("Snake eats", s.food)
}

func (s Snake) Move() {
	fmt.Println("Snake moves by", s.move)
}

func (s Snake) Speak() {
	fmt.Println("Snake speaks", s.sound)
}

func main() {
	var animals = make(map[string]Animal)

	for {
		var command, name, info string
		fmt.Print("> ")
		fmt.Scan(&command, &name, &info)
		command = strings.ToLower(command)
		name = strings.ToLower(name)
		info = strings.ToLower(info)

		switch command {
		case "newanimal":
			switch info {
			case "cow":
				animals[name] = Cow{"grass", "walk", "moo"}
				fmt.Println("Created it!")
			case "bird":
				animals[name] = Bird{"worms", "fly", "peep"}
				fmt.Println("Created it!")
			case "snake":
				animals[name] = Snake{"mice", "slither", "hsss"}
				fmt.Println("Created it!")
			default:
				fmt.Println("Invalid animal type.")
			}
		case "query":
			if an, ok := animals[name]; ok {
				switch info {
				case "eat":
					an.Eat()
				case "move":
					an.Move()
				case "speak":
					an.Speak()
				default:
					fmt.Println("Invalid info request.")
				}
			} else {
				fmt.Println("Invalid animal name. Please create a new animal.")
			}
		default:
			fmt.Println("Invalid command.")
		}
	}
}
