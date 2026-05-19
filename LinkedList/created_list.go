package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func Create_list() *Node {
	first := &Node{data: 10}
	second := &Node{data: 20}
	third := &Node{data: 30}

	first.next = second
	second.next = third
	third.next = nil

	head := first
	return head
}
func Display() {
	first := &Node{data: 10}
	second := &Node{data: 20}
	third := &Node{data: 30}
	forth := &Node{data: 40}
	fifth := &Node{data: 50}

	first.next = second
	second.next = third
	third.next = forth
	forth.next = fifth
	head := first

	current := head

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}

}
func Recursive_Display(head *Node) {
	Create_list()
	if head == nil {
		return
	}

	Recursive_Display(head.next)
	fmt.Println(head.data)

}
