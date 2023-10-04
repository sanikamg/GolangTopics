package main

import "fmt"

func main() {
	fmt.Println("one")
	defer fmt.Println("three")
	panic("a panic happened")
	fmt.Println("four")
}
