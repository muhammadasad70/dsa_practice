package main

func sum_of_el() int {
	values := []int{2, 34, 5, 6, 8, 9, 10, 11}
	sum := 0
	for i := 0; i <= len(values)-1; i++ {
		sum = sum + values[i]

	}
	return sum
}
func avg() int {
	values := []int{2, 34, 5, 6, 8, 9, 10, 11}
	sum := 0
	for i := 0; i < len(values); i++ {
		sum = sum + values[i]
	}
	avg := sum / len(values)
	return avg
}
