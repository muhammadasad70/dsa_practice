package main

func Max_of_nodes() int {
	first := &Node{data: 5}
	second := &Node{data: 10}
	third := &Node{data: 25}
	forth := &Node{data: 20}
	fifth := &Node{data: 15}

	first.next = second
	second.next = third
	third.next = forth
	forth.next = fifth
	head := first

	max := head.data
	for head != nil {
		if head.data > max {
			max = head.data

		}
		head = head.next
	}
	return max
}
