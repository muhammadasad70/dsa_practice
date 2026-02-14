package main

import "fmt"

// Question delete the given index value from the

func deleteing(index int) int {
	num := []int{10, 20, 30, 40, 50, 60, 70, 80}
	del_val := num[index]
	fmt.Println("Array before deleting")
	for i, j := range num {
		fmt.Println(i, j)
	}
	Index := index
	for i := Index; i < len(num)-1; i++ {
		num[i] = num[i+1]

	}
	num = num[:len(num)-1]
	fmt.Println("Array after deleting")
	for i, j := range num {
		fmt.Println(i, j)
	}
	return del_val

}
