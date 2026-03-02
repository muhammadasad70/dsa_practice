package main

import "fmt"

func intersection_nonsorted() {
	A := []int{3, 5, 5, 4, 6}
	B := []int{12, 4, 7, 4, 5}
	C := []int{}

	for i := 0; i < len(A); i++ {
		if search_2(B, A[i]) == true && search_2(C, A[i]) == false {
			C = append(C, A[i])
		} else {
			continue
		}

	}

	fmt.Println("After intersection")
	for _, k := range C {
		fmt.Println(k)
	}
}
func search_2(arr_1 []int, value int) bool {
	for i := 0; i < len(arr_1); i++ {
		if arr_1[i] == value {
			return true

		}
	}
	return false
}
