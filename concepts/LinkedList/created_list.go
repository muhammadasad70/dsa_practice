package main

type Node struct {
	data int
	next *Node
}

func Create_list() *Node {
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
	return head
}
