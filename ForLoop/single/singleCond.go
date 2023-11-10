package main

import "fmt"

//Write a Go program that finds and prints all the prime numbers between 1 and N (where N is an integer provided by the user).
//Use a single-condition for loop for this task.

func main() {
	var num int
	fmt.Println("Enter any number to print primenumbers till that number : ")
	fmt.Scan(&num)
	i := 0
	for i <= num {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Println(i)
		}
		i++
	}

}
