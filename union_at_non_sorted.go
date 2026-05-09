package main

import "fmt"

func union_nonsorted() {
	A := []int{3, 5, 10, 4, 6}
	B := []int{12, 4, 7, 2, 5}
	C := []int{}

	for i := 0; i < len(A); i++ {
		C = append(C, A[i])
	}
	for i := 0; i <len(B); i++ {
		if search(C, B[i]) == true {
			continue
		}
		C = append(C, B[i])

	}
	fmt.Println("After union this is the array union")

	for _, k := range C {
		fmt.Println(k)

	}
}
func search(arr_1 []int, value int) bool {
	for i := 0; i < len(arr_1); i++ {
		if arr_1[i] == value {
			return true

		}
	}
	return false
}
