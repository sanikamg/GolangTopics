// We use the defer statement to prevent the execution of a function until all other functions execute.
package main

import "fmt"

func main() {
	defer fmt.Println("a")
	fmt.Println("b")
	fmt.Println("c")
	fmt.Println("......................")
	MultipleDefer()
}

//When we use multiple defer in a program, the order of execution of the defer statements will be LIFO (Last In First Out).

func MultipleDefer() {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}
