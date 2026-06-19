package main

import "fmt"

func insertion(val int) {
	values := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("Values befor insertion")
	for _, j := range values {
		fmt.Println(j)
	}
	value := val
	values = append(values, 0)
	i := len(values) - 2

	for i >= 0 && values[i] > value {
		values[i+1] = values[i]
		i--

	}
	values[i+1] = value

	fmt.Println("Values befor insertion")
	for _, j := range values {
		fmt.Println(j)
	}

}
