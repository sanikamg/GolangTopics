package main

import "fmt"

func main() {
	str := "I am Sanika M G"
	str1 := `I am Sanika 
	M G`
	fmt.Println(str)
	fmt.Println(str1)
	k := len(str)
	fmt.Println(str[2:k])
}
