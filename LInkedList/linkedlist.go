package main

import (
	"fmt"
)

type LinkedList struct {
	number int
	next   *LinkedList
}

func (node *LinkedList) insert(num int) {
	var temp = new(LinkedList)
	temp.number = num
	temp.next = nil
	if node == nil {
		node = temp

	} else {
		var p = new(LinkedList)
		p = node
		for p.next != nil {
			p = p.next

		}
		p.next = temp

	}

}

func (node *LinkedList) Display() {
	var p = new(LinkedList)
	p = node.next
	for p != nil {
		fmt.Printf("%d -->", p.number)
		p = p.next
	}
}

func (node *LinkedList) insertBegin(num int) {

}

func main() {
	//initialize linked list
	head := new(LinkedList)
	num := []int{1, 4, 5, 6, 7}
	for i := 0; i < len(num); i++ {
		head.insert(num[i])
	}
	head.Display()

}
