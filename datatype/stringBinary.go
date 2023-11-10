package main

import "fmt"

func main() {
	s := "10010100"
	count := minOperations(s)
	fmt.Println(count)
}

func minOperations(s string) int {
	first := s[0]
	count := 0
	if first == '0' {
		for i := 1; i < len(s); i++ {
			if i%2 != 0 && s[i] != '1' {
				count++
			} else if i%2 == 0 && s[i] != '0' {
				count++
			}

		}
	} else {
		for i := 1; i < len(s); i++ {
			if i%2 != 0 && s[i] != '0' {
				count++
			} else if i%2 == 0 && s[i] != '1' {
				count++
			}
		}

	}

	return count
}
