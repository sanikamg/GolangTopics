//Write a Go program that launches two goroutines.
//One goroutine should print even numbers, and the other should print odd numbers.
//The main function should wait for both goroutines to finish before exiting.

package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go PrintEvenNumbers(&wg)
	go PrintOddNumbers(&wg)

	wg.Wait()

	fmt.Println("completed")

}

func PrintEvenNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func PrintOddNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
