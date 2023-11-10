// Write a Go program that creates two channels, numbers and squares.
//  The program should generate a sequence of numbers from 1 to 10, send them to the numbers channel,
//   and then calculate the square of each number and send it to the squares channel. Finally,
//   the program should print the values received from both channels.

package main

import "fmt"

func main() {
	numbers := make(chan int)
	squares := make(chan int)

	go num(numbers, squares)

	for i := 1; i <= 10; i++ {
		num := <-numbers
		sqr := <-squares
		fmt.Printf("%d * %d = %d\n", num, num, sqr)
	}
}

func num(number, square chan int) {
	for i := 1; i <= 10; i++ {
		number <- i
		square <- i * i
	}
	close(number)
	close(square)
}
