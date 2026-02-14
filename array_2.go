package main

import "fmt"

func arrays_2nd() {
	arr := [5]int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Println(index, value)
	}

}
func myslice() {
	// creating it manually so we use this method
	// slice_1 := []int{10, 20, 30, 40, 50}
	// for index, value := range slice_1 {
	// 	fmt.Println(index, value)
	// }
	// self declear the length and the capacity
	slice_2 := make([]int, 3, 5)
	slice_2 = append(slice_2, 100, 200)
	// for i, j := range slice_2 {
	// 	fmt.Println(i, "and", j)

	// }
	slice_3 := make([]int, 0, 5)
	slice_3 = append(slice_3, 23, 24, 25)
	for l, m := range slice_3 {
		fmt.Println(l, "and", m)
	}

}
func inserting(index int, val int) {
	num := []int{10, 20, 30, 40, 50, 60, 70, 80}
	Index := index
	Value := val
	num = append(num, 0)
	for i := len(num) - 1; i > Index; i-- {
		num[i] = num[i-1]
	}
	num[Index] = Value
	for i, j := range num {
		fmt.Println(i, "and", j)
	}

}
func deleting(index int){
	
}
