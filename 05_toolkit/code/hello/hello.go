package main

import (
	"fmt"

	"example.com/mathematics"
)


func main() {
	fmt.Println("Hey there.")
	total := mathematics.Add(1,2,3,4,5)
	fmt.Println(total)
}