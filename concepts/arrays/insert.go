package main

import "fmt"

// Question insert the value 87 at the index of 4 in the given or taken array

func insert(index int, value int) {
	num := []int{10, 20, 30, 40, 50, 60, 70, 80}
	Index := index
	Value := value
	num = append(num, 0)
	for i := len(num) - 1; i > Index; i-- {
		num[i] = num[i-1]
	}
	num[Index] = Value
	for j, k := range num {
		fmt.Println(j, k)
	}
}
