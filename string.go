package main

import (
	"fmt"
	"strconv"
)

func main() {
	//append string
	var sanika []string
	sanika = append(sanika, "")

	str1 := "sanika mg"
	for i := 0; i < len(str1); i++ {
		if string(str1[i]) != " " {
			sanika[0] += string(str1[i])
		}
	}
	fmt.Println(sanika)
	str2 := " mg"
	str1 += str2
	fmt.Println(str1)
	//slice of string
	fmt.Println(str1[0:])  //sanika mg
	fmt.Println(str1[1:])  //anika mg
	fmt.Println(str1[:3])  //san
	fmt.Println(str1[1:3]) //an
	fmt.Println(str1[:])   //sanika mg

	num := strconv.Itoa(123) //int to string
	fmt.Println(num)
	str, _ := strconv.Atoi("123")

	fmt.Println(str)
	fmt.Println("................................")
	k := 0
	var array []string
	text := "  this   is  a sentence "
	array = append(array, "")

	for i := 0; i < len(text); i++ {
		if string(text[i]) != " " {
			array[k] += string(text[i])
		} else if string(text[i]) == " " {
			if i > 0 && string(text[i-1]) == " " {
				continue
			} else {
				k++
				array = append(array, "")
			}
		}

	}
	fmt.Println(k)

	fmt.Println(array[4])
	fmt.Println(array)
}

//String

//-string is immutable
//-we can't change string after created
