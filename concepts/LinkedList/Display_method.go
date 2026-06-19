package main

import (
	"fmt"
)

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

func Display_fun(head *Node) {
	if head == nil {
		return
	}

	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}

}
