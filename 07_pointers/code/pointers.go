package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	var name string = "Julio"
// 	var namePointer *string
// 	//fmt.Println("Name *:", namePointer)

// 	namePointer = &name

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println(*namePointer)
// }

// // ******************************************************

// func main() {
// 	var name string = "Beyonce"
// 	var namePointer *string = &name
// 	var nameValue = *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name Value:", nameValue)

// }

// // ******************************************************

func changeName(n *string) {
	// returning a new thing also accomplishes the same thing, which is a modified version of whatever you are working with
	*n = strings.ToUpper(*n)
}

func main() {
	// name := "Elvis"
	// changeName(&name)
	// fmt.Println(name)

	type Coordinates struct{
		x, y float64
	}

	c := Coordinates{x: 5.5, y: 9.8}

	fmt.Println(c);

	c.x = 9.9;

	fmt.Println(c)
}

// // ******************************************************
