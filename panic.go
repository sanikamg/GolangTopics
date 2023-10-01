// We use the panic statement to immediately end the execution of the program.
// If our program reaches a point where it cannot be recovered due to some major errors, it's best to use panic.
// The lines of code after the panic statement are not executed.
package main

import "fmt"

func main() {

	Division(1, 2)
	Division(6, 3)
	Division(5, 0)
	Division(9, 0)
}

func Division(num1 int, num2 int) {
	if num2 == 0 {
		panic("can't be divided by 0")
	} else {
		fmt.Printf("%v / %v = %v ", num1, num2, num1/num2)
		fmt.Println("")
	}

}
