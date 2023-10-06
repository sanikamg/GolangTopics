package main

import "fmt"

//You are given a list of integers.
//Write a Go function filter that takes a list of integers and a predicate function as arguments.
//The function should return a new list containing only the elements that satisfy the given predicate.

func main() {

	numbers := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}

	even := filter(numbers, func(num int) bool {
		return num%2 == 0
	})
	fmt.Println(even)
}

func filter(arr []int, predicate func(int) bool) []int {

	result := []int{}
	for _, num := range arr {

		if predicate(num) {
			result = append(result, num)
		}
	}
	return result

}
