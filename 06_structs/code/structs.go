package main

import "fmt"

type User struct {
	ID int
	FirstName, LastName, Email string
}

type Group struct {
	role string
	users []User
	newestUser User
	spaceAvailable bool
}

func describeUser (user User) string{
	return fmt.Sprintf("Welcome, %s !", user.Email)
}

func describeGroup(group Group) string{
	if len(group.users) >= 2{
		group.spaceAvailable = false
	}
	description := fmt.Sprintf("Group has %d users. Newest user is %s. Space available: %t \n", len(group.users), group.newestUser.Email, group.spaceAvailable)

	return description
}

func modifyGroup(group Group) Group{
	if len(group.users) >= 2{
		fmt.Println("No more space available.")
		group.spaceAvailable = false
		return group
	}

	return group
}

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}

	u2 := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}

	g := Group{
		role: "Janitors",
		users: []User{u,u2},
		newestUser: u,
		spaceAvailable: true,
	}

	//describeUser(u)
	fmt.Println(describeGroup(g))
	fmt.Println(g)
	g.spaceAvailable = false
	fmt.Println(g)
	//newgroup := modifyGroup(g)
	//fmt.Println(newgroup)
}
