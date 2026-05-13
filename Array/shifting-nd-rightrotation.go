package main

import "fmt"

func right_shift() {
	values := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("Values befor the right_shift")
	for _, j := range values {
		fmt.Println(j)
	}

	for i := len(values) - 1; i > 0; i-- {
		values[i] = values[i-1]
	}
	fmt.Println("Values befor the right_shift")
	for _, j := range values {
		fmt.Println(j)
	}

}
func right_rotation() {
	values := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("Values befor the right_rotation")
	for _, j := range values {
		fmt.Println(j)
	}
	last := values[len(values)-1]

	for i := len(values) - 1; i > 0; i-- {
		values[i] = values[i-1]
	}
	values[0] = last
	fmt.Println("Values befor the right_rotation")
	for _, j := range values {
		fmt.Println(j)
	}

}
