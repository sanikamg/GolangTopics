// Write a Go program that continuously prompts the user for a number and prints whether the number is even or odd.
// The program should keep running until the user enters the number 0. Use an infinite loop for this task.
package main

import "fmt"

func main() {
	var num int
	for {
		fmt.Println("Enter any number : ")
		fmt.Scan(&num)
		if num%2 == 0 {
			fmt.Println("Entered Number is even")
		} else {
			fmt.Println("Entered Number is odd")
		}
		if num == 0 {
			break
		}
	}
}
