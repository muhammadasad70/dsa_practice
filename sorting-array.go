package main

import (
	"fmt"
)

func sorting() {
	arr := []int{-1, 2, -2, 1, -1, 2}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}

	}
	fmt.Println("After sorting array is ")
	for _, k := range arr {
		fmt.Println(k)
	}

}
