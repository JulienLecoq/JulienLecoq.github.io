package user

import (
	"juliano.com/passage-order/slice"
	"sort"
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

var Groups = []Group{
	GB1,
	GB2,
	LK1,
	LK2,
}
