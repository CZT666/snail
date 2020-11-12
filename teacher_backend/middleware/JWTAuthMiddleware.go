package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snail/teacher_backend/common"
	"snail/teacher_backend/utils"
)

const (
	AUTHORIZATION = "Authorization"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get(AUTHORIZATION)
		if authHeader == "" {
			baseResponse := new(common.BaseResponse)
			baseResponse.Code = common.AuthBlank
			c.JSON(http.StatusOK, baseResponse)
			c.Abort()
			return
		}

		mc, err := utils.ParseToken(authHeader)
		if err != nil {
			baseResponse := new(common.BaseResponse)
			baseResponse.Code = common.InvalidToken
			c.JSON(http.StatusOK, baseResponse)
			c.Abort()
			return
		}
		c.Set("user", mc)
		c.Next()
	}
}
