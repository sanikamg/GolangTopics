package main

import "fmt"

func main() {

	defer func() {
		a := recover()
		if a != nil {
			fmt.Println("recovered ", a)
		}
		fmt.Println("four") // This line will be reached after the panic is recovered.
	}()
	fmt.Println("one")
	defer fmt.Println("three")
	panic("a panic happened")
}
