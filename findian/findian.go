package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	str = strings.ToLower(str)
	if strings.HasPrefix(str, "i") && strings.Contains(str, "a") && strings.HasSuffix(str, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
