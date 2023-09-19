package main

import (
	"fmt"
	"strconv"
)

func main() {
	//append string
	str1 := "sanika"
	str2 := " mg"
	str1 += str2
	fmt.Println(str1)
	//slice of string
	fmt.Println(str1[0:])    //sanika mg
	fmt.Println(str1[1:])    //anika mg
	fmt.Println(str1[:3])    //san
	fmt.Println(str1[1:3])   //an
	fmt.Println(str1[:])     //sanika mg
	num := strconv.Itoa(123) //int to string
	fmt.Println(num)
	str, _ := strconv.Atoi("123")
	fmt.Println(str)
}

//String

//-string is immutable
//-we can't change string after created
