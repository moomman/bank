package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrTokenExpired = errors.New("token expired")
	ErrTokenInvalid = errors.New("invalid token")
)

type PayLoad struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	IssueAt  time.Time `json:"issue_at"`
	ExpireAt time.Time `json:"expire_at"`
}

func (p *PayLoad) Valid() error {
	if time.Now().After(p.ExpireAt) {
		return ErrTokenExpired
	}
	return nil
}

func NewPayLoad(username string, duration time.Duration) *PayLoad {
	return &PayLoad{
		ID:       uuid.New(),
		Username: username,
		IssueAt:  time.Now(),
		ExpireAt: time.Now().Add(duration),
	}
}
