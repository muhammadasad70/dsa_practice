package main

func main() {
	// Create_list()
	// fmt.Println("Simple display of the list is ")
	// Display()
	// fmt.Println("Recursive_Display")
	// head := Create_list()
	// Recursive_Display(head)
	// fmt.Println("Display using func")
	// Display_fun(head)
	// count := Count_node()
	// fmt.Println("Total node are :", count)
	// sum := Sum_of_nodes()
	// fmt.Println("Sum of all the nodes :", sum)
	// max := Max_of_nodes()
	// fmt.Println("Max element of the node is :", max)
	// val := 20
	// fmt.Println(Search(val))25

	var head *Node

	// head = IsEmpty(head, 10)
	head = InsertAtFirst(head, 5)
	head = InsertAtLast(head, 10)
	head = InsertAtLast(head, 20)
	head = InsertAtLast(head, 30)
	head = InsertAtSorted(head, 1)

	Display_fun(head)

}
