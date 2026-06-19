package main

func Sum_of_nodes() int {
	first := &Node{data: 5}
	second := &Node{data: 10}
	third := &Node{data: 15}
	forth := &Node{data: 20}
	fifth := &Node{data: 25}

	first.next = second
	second.next = third
	third.next = forth
	forth.next = fifth
	head := first

	sum := 0

	for head != nil {
		sum = sum + head.data
		head = head.next
	}
	return sum
}
