package token

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"time"
)

type Pasteo_maker struct {
	pasteo *paseto.V2
	key    []byte
}

func NewPasteo_maker(key string) (Maker, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("key is too short")
	}

	return &Pasteo_maker{
		pasteo: paseto.NewV2(),
		key:    []byte(key),
	}, nil
}

func (p *Pasteo_maker) CreateToken(username string, duration time.Duration) (string, error) {
	token, err := p.pasteo.Encrypt(p.key, NewPayLoad(username, duration), nil)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (p *Pasteo_maker) VerifyToken(token string) (*PayLoad, error) {
	var payLoad *PayLoad
	err := p.pasteo.Decrypt(token, p.key, payLoad, nil)
	if err != nil {
		return nil, ErrTokenInvalid
	}
	err = payLoad.Valid()
	if err != nil {
		return nil, ErrTokenExpired
	}
	return payLoad, nil
}
