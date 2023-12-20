package main

import (
	"fmt"
)

func main() {
	go fmt.Printf("Hello, World! from Go Routine\n")
	fmt.Printf("Hello, World! from Main\n")
}
