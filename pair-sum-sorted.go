package main

// bruit force algo

// func pair_sum(traget int) []int {
// 	arr := []int{-5, -2, 3, 4, 6}
// 	for i := 0; i < len(arr); i++ {
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[i]+arr[j] == traget {
// 				return []int{i, j}
// 			}

// 		}

// 	}
// 	return []int{0, 0}

// }

// using thw two pointer
func pair_sum_two_p(target int) []int {
	arr := []int{-1, 2, 3}
	i := 0
	j := len(arr) - 1
	for i < j {
		sum := arr[i] + arr[j]
		if sum == target {
			return []int{i, j}
		} else if sum < target {
			i++

		} else {
			j--
		}
	}
	return []int{}

}

// func pair_sum_sorting(target int){
// 	arr:=[]int{-5,-2,3,4,6}

// }
