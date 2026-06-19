package main

import (
	"fmt"
)

func Access() {
	name := "Muhammad Asad"
	fmt.Println(name[8])
	fmt.Println("Character of the string using this ")
	fmt.Printf("%c\n", name[0])
	fmt.Println("Length of the string is")
	fmt.Println(len(name))
}
