package main

import (
	"fmt"
)

func main() {
	var floatNum float32
	var intNum int32

	fmt.Printf("Enter floating point number: ")
	fmt.Scan(&floatNum)
	intNum = int32(floatNum)
	fmt.Printf("Int number is %d\n", intNum)
}
