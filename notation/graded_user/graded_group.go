package graded_user

import (
	// "fmt"
	"sort"

	"juliano.com/passage-order/slice"
	"juliano.com/passage-order/user"
)

type GradedGroup struct {
	Name  string
	Users []*GradedUser
}

func (g *GradedGroup) RandomizeUsers() {
	slice.Shuffle(g.Users)
}

func (g *GradedGroup) SortAlphabetically() {
	sort.Slice(g.Users, func(i, j int) bool {
		return g.Users[i].LastName < g.Users[j].LastName
	})
}

func (g *GradedGroup) EmailAdresses() string {
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

func (g *GradedGroup) DeleteUser(index int) {
	g.Users = slice.Delete(g.Users, index)
}

func (g GradedGroup) String() string {
	return g.Name
}

func toGradedGroup(group user.Group) GradedGroup {
	users := []*GradedUser{}

	for _, user := range group.Users {
		users = append(users, &GradedUser{
			User: user,
		})
	}

	return GradedGroup{
		Name:  group.Name,
		Users: users,
	}
}

var GB1 = toGradedGroup(user.GB1)
var GB2 = toGradedGroup(user.GB2)
var LK1 = toGradedGroup(user.LK1)
var LK2 = toGradedGroup(user.LK2)

var Groups = []*GradedGroup{
	&GB1,
	&GB2,
	&LK1,
	&LK2,
}
