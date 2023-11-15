package main

import "fmt"

func main() {
	var p *string
	str := "sanika"
	p = &str
	fmt.Println(*p)
	fmt.Println(p)

}
