package token

import (
	"github.com/moonman/mbank/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJwtMaker_CreateToken(t *testing.T) {
	maker, err := NewJwtMaker(utils.RandomString(32))
	require.NoError(t, err)
	token, err := maker.CreateToken("moon", 2*time.Second)
	require.NoError(t, err)
	time.Sleep(2 * time.Second)
	_, err = maker.VerifyToken(token)
	require.EqualError(t, err, ErrTokenExpired.Error())
}
