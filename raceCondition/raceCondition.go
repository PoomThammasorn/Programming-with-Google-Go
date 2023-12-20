package main

/*
	Explanation:
	Race condition occurs when two or more goroutines access the same variable concurrently.
	In this case, the variable i is accessed by two goroutines.
	The first goroutine prints the value of i, while the second goroutine increments the value of i.
	The output of the program is unpredictable.
	It can be 3 or 4.
*/

import (
	"fmt"
	"time"
)

func main() {
	// race condition
	// To check race condition, run the following command:
	// go run -race PATH_TO_FILE/raceCondition.go
	i := 3

	go func() {
		fmt.Println(i)
	}()

	go func() {
		i += 1
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("done")
	fmt.Println("\nExplanation:")
	fmt.Println("condition occurs when two or more goroutines access the same variable concurrently.")
	fmt.Println("this case, the variable i is accessed by two goroutines.")
	fmt.Println("The first goroutine prints the value of i, while the second goroutine increments the value of i.")
	fmt.Println("The output of the program is unpredictable.")
	fmt.Print("It can be 3 or 4.\n\n\n")
}
