package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestComparePassword(t *testing.T) {
	hashPassword, err := HashPassword(RandomString(6))
	require.NoError(t, err)

	err = ComparePassword(hashPassword, RandomString(6))
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
