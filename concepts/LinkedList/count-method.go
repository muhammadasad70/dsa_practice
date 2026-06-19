package main

func Count_node() int {
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

	count := 0

	for current != nil {
		count = count + 1
		current = current.next
	}
	return count
}
