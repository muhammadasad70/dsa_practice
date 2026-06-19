package main

import (
	"fmt"
)

func Loop_String() {
	name := "Muhammad Asad"
	as := " "
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c\n", name[i])
		fmt.Println(name[i])
		fmt.Println(as[0])
	}
}
