package main

import "fmt"

func changeEmail(user *User) {
	user.email = "somethingelse@gmail.com"
}

type User struct {
	id          int
	name, email string
}

func main() {
	user := User{id: 5, name: "Joao", email: "peste@peste.com"}
	changeEmail(&user)
	fmt.Println(user)
}