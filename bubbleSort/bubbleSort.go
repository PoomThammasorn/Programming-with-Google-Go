package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]

}

func bubbleSort(sli []int) {
	var isSwap bool
	size := len(sli)
	for c := 0; c < size*(size-1)/2; c++ {
		isSwap = false
		for i := 0; i < size-1; i++ {
			if sli[i] > sli[i+1] {
				swap(sli, i)
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
}

func parseInput(input string) []int {
	inputStr := strings.Fields(input)
	// fmt.Println("inputStr", inputStr)
	result := make([]int, len(inputStr))

	for i, numStr := range inputStr {
		num, _ := strconv.Atoi(numStr)
		result[i] = num
	}
	// fmt.Println("perse", result)
	return result
}

func main() {
	var sli []int
	fmt.Print("Enter a number: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	sli = parseInput(input)
	bubbleSort(sli)
	fmt.Println("Sorted number:", sli)
}
