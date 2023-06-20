package models

type UserHistory struct {
	FirstName string
	LastName  string
	Products  map[string][]int
}
