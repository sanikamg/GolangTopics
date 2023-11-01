// Write a Go program that creates a channel, spawns a goroutine,
// and sends a value through the channel. The main program should
// receive and print the value.
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go Sample("hii", c)

	for {
		value := <-c
		fmt.Printf("goroutine value : %s\n", value)
	}

}

func Sample(msg string, ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Second)
	}
}
