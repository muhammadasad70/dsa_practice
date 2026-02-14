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

func searching(value int) int {
	num := []int{10, 20, 30, 40, 50, 60, 70, 80}
	Value := value
	for i := 0; i < len(num)-1; i++ {
		if Value == num[i] {
			return i
		}
	}
	return -1
}
