//Suppose you have a program that deals with a large dataset stored in an array.
//You want to create a function that increments each element of this array by a certain value.
//To optimize memory usage, you decide to pass this array by reference (using a pointer) instead of passing it by value.
// How would you implement this function in Go, and how would you use pointers to achieve this without making unnecessary copies of the array?

package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("original array :", data)

	incrementArray(&data, 10)

	fmt.Println("value added array : ", data)

}

func incrementArray(arr *[]int, value int) {
	for i := range *arr {
		(*arr)[i] += value
	}
}
