package router

type group struct {
	Account account
	User    user
}

var Group = new(group)
