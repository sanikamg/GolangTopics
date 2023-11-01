package main

import "fmt"

func findMiddleElement(arr1, arr2 []int) int {
	totalLength := len(arr1) + len(arr2)
	middleIndex := totalLength / 2

	start1, end1 := 0, len(arr1)-1
	start2, end2 := 0, len(arr2)-1

	for {
		if start1 > end1 {
			return arr2[start2+middleIndex-len(arr1)]
		}
		if start2 > end2 {
			return arr1[start1+middleIndex-len(arr2)]
		}

		mid1 := (start1 + end1) / 2
		mid2 := (start2 + end2) / 2

		if mid1+mid2 < middleIndex {
			if arr1[mid1] < arr2[mid2] {
				start1 = mid1 + 1
			} else {
				start2 = mid2 + 1
			}
		} else {
			if arr1[mid1] > arr2[mid2] {
				end1 = mid1 - 1
			} else {
				end2 = mid2 - 1
			}
		}
	}
}

func main() {
	arr1 := []int{1, 5, 9}
	arr2 := []int{6, 7}

	middleElement := findMiddleElement(arr1, arr2)
	fmt.Println("Middle Element:", middleElement)
}
