package main

import "fmt"

func getting(index int) int {
	values := []int{2, 3, 4, 6, 8, 10, 12, 23, 19, 20, 31}
	if index >= 0 && index < len(values) {
		return values[index]
	}
	return 0
}

func setting(index int, val int) {
	values := []int{3, 4, 5, 7, 9, 11, 13, 24, 20, 21, 32}
	fmt.Println("The array or slice before setting the values")
	for i, j := range values {
		fmt.Println("index is :", i, "value is :", j)
	}
	if index >= 0 && index < len(values) {
		values[index] = val
	}
	fmt.Println("The array or slice after setting the values")
	for i, j := range values {
		fmt.Println("index is :", i, "value is :", j)
	}
}
