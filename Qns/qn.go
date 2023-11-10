package main

import "fmt"

func main() {
	num := 35 // Assuming you have an initial value for num

	for num < 10 {
		sum := 0

		for num == 1 {
			if num%10 != 0 {
				sum += num % 10
			}
			num = sum // You need to update num here
		}
	}

	fmt.Println("Final num:", num)
}
