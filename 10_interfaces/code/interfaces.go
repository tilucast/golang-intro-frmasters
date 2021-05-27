package main

import "fmt"

type Describer interface {
	describe() string
}

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func Describing(describer Describer) string {
	return describer.describe()
}

func (u *User) describe() string {
	return fmt.Sprintf("Name: %s %s, Email: %s, ID: %d \n", u.FirstName, u.LastName, u.Email, u.ID)
}

func (g *Group) describe() string {
	return fmt.Sprintf("The %s user group has %d users. Newest user:  %s, Accepting New Users:  %t \n", g.role, len(g.users), g.newestUser.FirstName, g.spaceAvailable)
}

func main() {
	u1 := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 1, FirstName: "Humphrey", LastName: "Bogart", Email: "humphrey.bogart@gmail.com"}
	g := Group{role: "admin", users: []User{u1, u2}, newestUser: u2, spaceAvailable: true}

	du := Describing(&u1)
	dg := Describing(&g)

	fmt.Println(du, dg)
}
