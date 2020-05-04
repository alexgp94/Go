package models

type User struct {
	Id 		 	int
	Firstname 	string
	Lastname  	string
}

var (
	user   []*User
	nextId = 1
)
