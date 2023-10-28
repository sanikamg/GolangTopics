package main

import "fmt"

func main() {
	fmt.Println("Main function")

	go countNumbers(20)

	fmt.Println("End main function")
}

func countNumbers(limit int) {
	num := 0
	for i := 1; i < limit; i++ {
		num += i
	}
	fmt.Println("Num: ", num)
}
