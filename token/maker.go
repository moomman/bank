package token

import (
	"time"
)

//管理token

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)

	VerifyToken(token string) (*PayLoad, error)
}

var TokenMaker Maker
