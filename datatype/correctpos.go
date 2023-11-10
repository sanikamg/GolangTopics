package main

import "fmt"

func main() {
	a := "a1c1b3z"
	for i := 0; i < len(a); i++ {
		if i%2 == 0 {
			fmt.Printf("%c is the %d th character\n", a[i], int(a[i]-'a')+1)
		} else {
			fmt.Printf("%c's integer value is %d\n", a[i], int(a[i]-'0'))
		}
	}

}
