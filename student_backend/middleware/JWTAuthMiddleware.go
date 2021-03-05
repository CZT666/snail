package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"student_bakcend/models"
	"student_bakcend/vo"
)

const (
	AUTHORIZATION = "Authorization"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get(AUTHORIZATION)
		if authHeader == "" {
			baseResponse := new(vo.BaseResponse)
			baseResponse.Code = vo.AuthBlank
			c.JSON(http.StatusOK, baseResponse)
			c.Abort()
			return
		}

		mc, err := models.ParseToken(authHeader)
		if err != nil {
			baseResponse := new(vo.BaseResponse)
			baseResponse.Code = vo.InvalidToken
			c.JSON(http.StatusOK, baseResponse)

			c.Abort()
			return
		}
		c.Set("user", mc)
		c.Next()
	}
}
