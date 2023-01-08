package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/moonman/mbank/token"
	"net/http"
	"strings"
)

const (
	AuthHeader        = "Authorization"
	AuthorizationType = "Bearer"
)

var (
	ErrParseHeader = errors.New("Invalid authorization header")
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		val := c.GetHeader(AuthHeader)
		fields := strings.Fields(val)

		if len(fields) != 2 || fields[0] != AuthorizationType {
			c.AbortWithError(http.StatusUnauthorized, ErrParseHeader)
			return
		}

		payLoad, err := token.TokenMaker.VerifyToken(fields[1])
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		if err := payLoad.Valid(); err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Next()
	}
}
