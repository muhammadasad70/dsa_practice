package main

import "fmt"

func intersection_at_sorted() {
	A := []int{3, 4, 5, 6, 10}
	B := []int{2, 4, 5, 7, 12}
	C := []int{}

	i := 0
	j := 0

	for i < len(A) && j < len(B) {
		if A[i] < B[j] {

			i++
		} else if B[j] < A[i] {

			j++
		} else if A[i] == B[j] {
			C = append(C, A[i]) 
			i++
			j++
		}
	}
	fmt.Println("After intersection")
	for _, k := range C {
		fmt.Println(k)
	}
}
