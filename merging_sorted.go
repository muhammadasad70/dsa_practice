package main

import "fmt"

func merging() {
	array_1 := []int{3, 8, 16, 20, 25}
	array_2 := []int{4, 10, 12, 22, 23}
	array_3 := []int{}

	i := 0
	j := 0
	for i < len(array_1) && j < len(array_2) {
		if array_1[i] < array_2[j] {
			array_3 = append(array_3, array_1[i])
			i++
		} else {
			array_3 = append(array_3, array_2[j])
			j++
		}
	}
	for ; i < len(array_1); i++ {
		array_3 = append(array_3, array_1[i])
	}
	for ; j < len(array_2); j++ {
		array_3 = append(array_3, array_2[j])
	}
	fmt.Println("after merging sorted array is :")
	for _, l := range array_3 {
		fmt.Println(l)
	}

}
