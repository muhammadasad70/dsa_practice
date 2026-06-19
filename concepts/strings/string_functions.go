package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Pre_function() {
	name := "Muhammad"
	fmt.Println(strings.Contains(name, "mu"))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.ToLower(name))
	text := "apple,banana,mango"
	parts := strings.Split(text, ",")
	fmt.Println(parts)
	arr := []string{"Go", "java", "python"}
	result := strings.Join(arr, "-")
	fmt.Println(result)
	text_1 := "apple,banana,mango"
	res := strings.ReplaceAll(text_1, "apple", "lokat")
	fmt.Println(res)
	num, _ := strconv.Atoi("")
	fmt.Println(num)
	str := strconv.Itoa(123)
	fmt.Println(str)

}
