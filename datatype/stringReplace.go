// Program using Replace() to replace strings

package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "carrrr"
	fmt.Println("Old String:", text)

	// replace r with t
	replacedText := strings.Replace(text, "r", "t", 2)

	fmt.Println("New String:", replacedText)

	str := "123"
	fmt.Println(string(str[0]))
}
