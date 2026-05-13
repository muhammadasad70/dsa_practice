package main

import "fmt"

func arrays() {
	arr := [5]int{1, 2, 3, 4, 5}
	for index, values := range arr {
		fmt.Println(index, values)
	}

}
func slice() {
	myslice := []int{10, 20, 30, 40, 50}
	fmt.Println(myslice)
	myslice2 := make([]int, 0, 3)
	myslice2 = append(myslice2, 23, 34, 56)
	fmt.Println(myslice2)
}
