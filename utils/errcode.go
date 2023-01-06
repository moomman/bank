package utils

import "github.com/gin-gonic/gin"

func ErrS(err error) gin.H {
	return gin.H{"err": err}
}
