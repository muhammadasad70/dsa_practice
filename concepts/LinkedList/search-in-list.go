package main

func Search(value int) bool {
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

	for head != nil {
		if head.data == value {
			return true
		}
		head = head.next
	}
	return false
}
