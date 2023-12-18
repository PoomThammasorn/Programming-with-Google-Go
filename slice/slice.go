package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// Create an integer slice of size 3
	sli := make([]int, 3)

	// Prompt user to enter an integer or 'X' to exit
	var input string
	fmt.Scan(&input)

	// Initialize a counter to keep track of iterations
	count := 0

	// Continue looping until the user enters 'X' or 'x'
	for input != "X" && input != "x" {
		// Convert the input to an integer
		num, _ := strconv.Atoi(input)

		// Switch based on the count of iterations
		switch count {
		case 0:
			// If it's the first iteration, set the first element of the slice
			sli[0] = num
			fmt.Println(sli[:1])
		case 1:
			// If it's the second iteration, compare and update the second element
			if num < sli[0] {
				sli[1] = sli[0]
				sli[0] = num
			} else {
				sli[1] = num
			}
			fmt.Println(sli[:2])
		case 2:
			// If it's the third iteration, set the third element and sort the slice
			sli[2] = num
			sort.Ints(sli)
			fmt.Println(sli)
		default:
			// For subsequent iterations, append the number, sort the slice, and print
			sli = append(sli, num)
			sort.Ints(sli)
			fmt.Println(sli)
		}

		// Increment the counter
		count++

		// Prompt user for new input
		fmt.Scan(&input)
	}
}
