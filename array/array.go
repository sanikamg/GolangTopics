package main

import (
	"fmt"
	"sort"
)

func main() {
	// Define a two-dimensional array
	arr := [][]int{
		{1, 5},
		{3, 2},
		{2, 8},
		{4, 1},
	}

	// Define a custom sorting function
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	arr1 := make([][]int, len(arr))
	//to make copy of an array
	copy(arr1[:], arr[:])

	// Print the sorted array
	fmt.Println(arr, arr1)
}
