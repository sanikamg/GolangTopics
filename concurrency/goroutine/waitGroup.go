package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark the task as done when the worker finishes

	fmt.Printf("Worker %d starting\n", id)
	// Perform some work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Add 3 to the WaitGroup (for 3 workers)
	wg.Add(3)

	// Launch 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers done")
}
