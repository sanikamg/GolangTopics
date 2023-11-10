package main

import "fmt"

func divideByZero(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	if b == 0 {
		panic("can't be divided by 0")
	}
}

func main() {

	fmt.Println("Start of the program")

	divideByZero(10, 0)

	fmt.Println("End of the program")
}
