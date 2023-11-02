package main

import "fmt"

func findMiddleElement(arr1, arr2 []int) int {
	var mid int
	totalLength := len(arr1) + len(arr2)
	mid = (totalLength / 2) + 1
	if mid < len(arr1) {
		return arr1[mid-1]
	}

	if mid > len(arr2) {
		return arr2[0]
	}

	return arr2[len(arr2)-mid]

}

func main() {
	arr1 := []int{1, 5, 9}
	arr2 := []int{2, 7, 9}

	middleElement := findMiddleElement(arr1, arr2)
	fmt.Println("Middle Element:", middleElement)
}
