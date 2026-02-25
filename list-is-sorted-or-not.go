package main

func sorted() int {
	arr := []int{10, 92, 30, 45, 67, 90}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return 1
		}
	}
	return 0

}
