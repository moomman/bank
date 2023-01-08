package api

//对外访问接口
type group struct {
	Account account
	User    user
}

var Group = new(group)
