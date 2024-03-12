package user

import (
	"sort"

	"juliano.com/passage-order/slice"
)

type Group struct {
	Name  string
	Users []User
}

func (g *Group) RandomizeUsers() {
	slice.Shuffle(g.Users)
}

func (g *Group) SortAlphabetically() {
	sort.Slice(g.Users, func(i, j int) bool {
		return g.Users[i].LastName < g.Users[j].LastName
	})
}

func (g *Group) EmailAdresses() string {
	emails := ""

	for i, user := range g.Users {
		if i < len(g.Users)-1 {
			emails += user.Email + ", "
		} else {
			// For the last element, don't add ", "
			emails += user.Email
		}
	}

	return emails
}

func (g *Group) DeleteUser(index int) {
	g.Users = slice.Delete(g.Users, index)
}

func (g Group) String() string {
	return g.Name
}

var LK1 = Group{
	Name:  "LK1",
	Users: LK1Users,
}

var LK2 = Group{
	Name:  "LK2",
	Users: LK2Users,
}

var GB1 = Group{
	Name:  "GB1",
	Users: GB1Users,
}

var GB2 = Group{
	Name:  "GB2",
	Users: GB2Users,
}

var Groups = []Group{
	GB1,
	GB2,
	LK1,
	LK2,
}
