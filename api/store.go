package api

//对外访问接口
type group struct {
	Account account
}

var Group = new(group)
