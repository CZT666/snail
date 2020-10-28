package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snail/teacher_backend/common"
	"snail/teacher_backend/utils"
	"strings"
)

const (
	AUTHORIZATION = "Authorization"
	BEARER        = "Bearer"
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

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == BEARER) {
			baseResponse := new(common.BaseResponse)
			baseResponse.Code = common.AuthFormatError
			c.JSON(http.StatusOK, baseResponse)
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := utils.ParseToken(parts[1])
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
