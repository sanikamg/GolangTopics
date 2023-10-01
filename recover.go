//The panic statement immediately terminates the program.
//However, sometimes it might be important for a program to complete its execution and get some required results.

//In such cases, we use the recover statement to handle panic in Go.
//The recover statement prevents the termination of the program and recovers the program from panic.

package main

import "fmt"

func main() {

	Div(1, 2)
	Div(6, 3)
	Div(5, 0)
	Div(9, 0)
}

func HandlePanic() {
	a := recover()
	if a != nil {
		fmt.Println("Recover", a)
	}
}
func Div(num1 int, num2 int) {
	defer HandlePanic()

	if num2 == 0 {
		panic("can't be divided by 0")
	} else {
		fmt.Printf("%v / %v = %v ", num1, num2, num1/num2)
		fmt.Println("")
	}

}
