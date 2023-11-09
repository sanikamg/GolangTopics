//Write a Go program that calculates the sum of all even numbers from 1 to 100 and prints the result.

package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i = i + 1 {

		sum += i

	}
	fmt.Println("sum of even numbers between 0 and 100 : ", sum)
	fmt.Println("sum of even numbers between 0 and 100 : ", 100*(100+1)/2)
}
