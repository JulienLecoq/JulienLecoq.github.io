package user

import (
	"strings"

	"juliano.com/passage-order/stringg"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
}

func (u *User) FullName() string {
	return u.LastName + " " + u.FirstName
}

func (u *User) FullNameFormatted() string {
	return strings.ToUpper(u.LastName) + " " + stringg.CapitalizeEachWord(u.FirstName)
}

func (u User) String() string {
	return u.FullNameFormatted()
}
