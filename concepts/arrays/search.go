package main

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
