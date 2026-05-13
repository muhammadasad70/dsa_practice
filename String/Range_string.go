package main

import (
	"fmt"
)

func Range_String() {
	name := "Muhammad Asad"
	for index, value := range name {
		fmt.Println(index, "and value is", value)

	}
}
