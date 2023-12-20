package main

import (
	"fmt"
	"strconv"
	"sync"
)

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			// swap arr
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// swap arr
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}

func quicksort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quicksort(arr, low, pi-1)
		quicksort(arr, pi+1, high)
	}
}

func merge(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))
	idx1, idx2, idx := 0, 0, 0
	for idx1 < len(arr1) && idx2 < len(arr2) {
		if arr1[idx1] < arr2[idx2] {
			result[idx] = arr1[idx1]
			idx1++
		} else {
			result[idx] = arr2[idx2]
			idx2++
		}
		idx++
	}
	for idx1 < len(arr1) {
		result[idx] = arr1[idx1]
		idx1++
		idx++
	}
	for idx2 < len(arr2) {
		result[idx] = arr2[idx2]
		idx2++
		idx++
	}
	return result
}

func callSort(arr []int, wg *sync.WaitGroup) {
	quicksort(arr, 0, len(arr)-1)
	wg.Done()
}

func main() {
	var size int
	var wg sync.WaitGroup

	wg.Add(4)
	fmt.Print("Size of array: ")
	fmt.Scan(&size)
	fmt.Printf("Input %d data: ", size)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	arr1 := arr[:size/4]
	arr2 := arr[size/4 : size/2]
	arr3 := arr[size/2 : 3*size/4]
	arr4 := arr[3*size/4:]

	go callSort(arr1, &wg)
	go callSort(arr2, &wg)
	go callSort(arr3, &wg)
	go callSort(arr4, &wg)
	wg.Wait()

	fmt.Println("Array 1:", arr1)
	fmt.Println("Array 2:", arr2)
	fmt.Println("Array 3:", arr3)
	fmt.Println("Array 4:", arr4)
	xw

	result1 := merge(arr1, arr2)
	result2 := merge(arr3, arr4)
	fmt.Print("Sorted result: ")
	for _, e := range merge(result2, result1) {
		fmt.Print(strconv.Itoa(e) + " ")
	}
	fmt.Println()

}
