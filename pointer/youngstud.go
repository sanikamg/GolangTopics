// You're working on a program that manages student records. Each student has a name and an age.
// You're tasked with creating a function that takes a slice of student structs and finds the youngest student.
// Write a function that accepts a slice of student structs as a parameter and returns a pointer to the youngest student in the slice.
// If multiple students have the same youngest age, return a pointer to the first occurrence in the slice.
// How would you implement this in Go?
package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func main() {

	stud := []Student{{"sanika", 10}, {"sana", 25}, {"sani", 5}, {"vishnu", 15}}
	ynstud := FindYoungest(stud)
	fmt.Println(ynstud.Age, ynstud.Name)

}

func FindYoungest(stud []Student) *Student {
	youngest := &stud[0]

	for i := 1; i < len(stud); i++ {
		if stud[i].Age < youngest.Age {
			youngest = &stud[i]
		}
	}
	return youngest
}
