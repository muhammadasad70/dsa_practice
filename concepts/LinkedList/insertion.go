package main

func IsEmpty(head *Node, value int) *Node {
	if head == nil {
		newNode := &Node{data: value}
		head = newNode

	}
	return head
}

func InsertAtFirst(head *Node, value int) *Node {
	newNode := &Node{data: value}
	newNode.next = head
	head = newNode

	return head
}
func InsertAtLast(head *Node, value int) *Node {
	newNode := &Node{data: value}
	newNode.next = nil

	current := head

	for current.next != nil {
		current = current.next
	}
	current.next = newNode

	return head

}
func InsertionAtGiven(head *Node, position int, value int) *Node {
	newNode := &Node{data: value}
	current := head

	for i := 1; i < position-1; i++ {
		current = current.next
	}
	newNode.next = current.next
	current.next = newNode

	return head

}

func InsertAtSorted(head *Node, value int) *Node {
	newNode := &Node{data: value}

	if head == nil || value < head.data {
		newNode.next = head
		return newNode
	}
	current := head

	for current.next != nil && current.next.data < value {
		current = current.next
	}
	newNode.next = current.next
	current.next = newNode

	return head
}
