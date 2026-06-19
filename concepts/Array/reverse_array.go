package main

import "fmt"

func reverse_1() {
	values := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("Values befor the reverse")
	for i, j := range values {
		fmt.Println(i, "and", j)
	}
	values_2 := make([]int, len(values))
	j := 0

	for i := len(values) - 1; i >= 0; i-- {
		values_2[j] = values[i]
		j++
	}
	fmt.Println("Values after the reverse")
	for l, m := range values_2 {
		fmt.Println(l, "and", m)
	}

}
func reverse_2() {
	values := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("Values befor the reverse")
	for i, j := range values {
		fmt.Println(i, "and", j)
	}
	i := 0

	for j := len(values) - 1; j > i; j-- {
		temp := values[i]
		values[i] = values[j]
		values[j] = temp
		i++
	}
	fmt.Println("Values after the reverse")
	for l, m := range values {
		fmt.Println(l, "and", m)
	}

}

// in this i have use the two pointer concept
func reverse_3() {
	values := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("Values befor the reverse")
	for i, j := range values {
		fmt.Println(i, "and", j)
	}
	i := 0
	j := len(values) - 1

	for i < j {
		values[j], values[i] = values[i], values[j]
		i++
		j--
	}
	fmt.Println("Values after the reverse")
	for l, m := range values {
		fmt.Println(l, "and", m)
	}

}
