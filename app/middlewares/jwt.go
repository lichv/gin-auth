package middlewares

import (
	jwt2 "gin-auth/utils/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.DefaultQuery("token","")
		if token == ""{
			token = c.DefaultPostForm("token","")
		}
		if token == ""{
			token, _ = c.Cookie("token")
		}
		if token == "" {
			token = c.GetHeader("X-TOKEN")
		}

		if token == "" {
			code = 400
		} else {
			_, err := jwt2.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = 404
				default:
					code = 405
				}
			}
		}

		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "token失效",
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}