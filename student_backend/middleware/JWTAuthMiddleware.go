package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"snail/student_bakcend/models"
	"snail/student_bakcend/vo"
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
		log.Printf("user -------%v\n", mc)
		c.Set("user", mc)
		c.Next()
	}
}
