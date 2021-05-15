package main

import "fmt"

func main() {
	//var sentence = "I am a sentence boi."
	sentence := "I am a sentence boi."

	for index, letter := range sentence {
		// "_" like javascript, indicates a variable that is not going to be used.
		if index % 2 != 0 {
			fmt.Println(string(letter))
		}
	}
}