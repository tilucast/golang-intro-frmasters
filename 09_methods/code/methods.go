package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

type Group struct {
	role string
	users []User
	newestUser User
	spaceAvailable bool
}

func (group *Group) describe(){
	
	group.spaceAvailable = false
	fmt.Println(group)
	
}

func (user User) describe() {
	fmt.Printf("Name: %s %s, Email: %s \n", user.FirstName, user.LastName, user.Email)
}

func main() {
	user := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}

	group := Group{role: "Janitors", users: []User{user}, newestUser: user, spaceAvailable: true}

	user.describe()
	group.describe()
}
