package main

import (
	"fmt"
)

func main() {
	i := 2
	tryPointer(&i)
}

func tryPointer(p *int) {
	fmt.Println(p)
}
