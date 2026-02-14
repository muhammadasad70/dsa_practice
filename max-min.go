package main

func max_value() int {
	values := []int{3, 4, 7, 5, 11, 10, 13, 12, 20, 21, 9}
	max := values[0]
	for i := 1; i < len(values)-1; i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}
func min_value() int {
	values := []int{90, 4, 7, 5, 11, 10, 13, 12, 20, 21, 9}
	min := values[0]
	for i := 1; i < len(values)-1; i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min
}
