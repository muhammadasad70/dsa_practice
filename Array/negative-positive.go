package main

import "fmt"

func neg_pos() {
	arr := []int{-6, 3, -8, 10, 5, -7, -9, 12, -4, 2}
	fmt.Println("Values befor swapping")
	for _, j := range arr {
		fmt.Println(j)
	}
	i := 0
	j := len(arr) - 1

	for i < j {
		if arr[i] < 0 {
			i++
		} else if arr[j] >= 0 {
			j--
		} else {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	fmt.Println("Values after swapping")
	for _, j := range arr {
		fmt.Println(j)
	}

}
