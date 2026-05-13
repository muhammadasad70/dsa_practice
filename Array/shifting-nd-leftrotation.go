package main

import "fmt"

func left_shift() {
	values := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("Values befor the left_shift")
	for _, j := range values {
		fmt.Println(j)
	}

	for i := 0; i < len(values)-1; i++ {
		values[i] = values[i+1]
	}
	fmt.Println("Values befor the left_shift")
	for _, j := range values {
		fmt.Println(j)
	}

}
func left_rotation() {
	values := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("Values befor the left_rotation")
	for _, j := range values {
		fmt.Println(j)
	}
	first := values[0]

	for i := 0; i < len(values)-1; i++ {
		values[i] = values[i+1]
	}
	values[len(values)-1] = first
	fmt.Println("Values befor the left_rotation")
	for _, j := range values {
		fmt.Println(j)
	}

}
