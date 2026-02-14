package main

import "fmt"

func main() {
	// arrays()
	// slice()
	// fmt.Println("Please enter the index at which you want to insert the value")
	// index := 0
	// fmt.Scan(&index)
	// fmt.Println("Please enter the values which you want to int the array")
	// value := 0
	// fmt.Scan(&value)
	// insert(index, value)
	// fmt.Println("Please enter the index of the value you want to delete ")
	// val := 0
	// fmt.Scan(&val)
	// Val := deleteing(val)
	// fmt.Println("Delete value is ", Val)
	// fmt.Println("Please enter the key value you want to search")
	// Key := 0
	// fmt.Scan(&Key)
	// val := searching(Key)
	// if val != -1 {
	// 	fmt.Println("Successfully search at the index :", val)
	// } else {
	// 	fmt.Println("Not found the given key :(")
	// }
	// arrays_2nd()
	myslice()
	fmt.Println("Please enter the index at which you want to insert the value")
	index := 0
	fmt.Scan(&index)
	fmt.Println("Please enter the values which you want to int the array")
	value := 0
	fmt.Scan(&value)
	inserting(index, value)

}
