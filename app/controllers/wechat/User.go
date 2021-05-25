package wechat

import (
	"fmt"
	"gin-auth/app/services"
	"gin-auth/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSelf(c *gin.Context) {
	token := c.DefaultQuery("token", "")
	if token == "" {
		token = c.DefaultPostForm("token", "")
	}
	if token == "" {
		token, _ = c.Cookie("token")
	}
	if token == "" {
		token = c.GetHeader("X-TOKEN")
	}
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"state":   3000,
			"message": "缺少参数",
		})
		return
	}
	claims, err := jwt.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"state":   3001,
			"message": err.Error(),
		})
		return
	}
	fmt.Println(claims)
	code := claims.Code
	if code != "" &&  len(code) > 2 &&  code[:2] == "wx" {
		one, err := services.GetWechatUserOne(map[string]interface{}{"code": code}, "code asc")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"state":   3001,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"state":   2000,
			"message": "success",
			"data":    one,
		})
		return
	} else {
		one, err := services.GetWechatUserOne(map[string]interface{}{"code": code}, "code asc")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"state":   3001,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"state":   2000,
			"message": "success",
			"data":    one,
		})
		return
	}

}

func GetWechatUser(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		code = c.DefaultPostForm("code", "")
	}
	if code == "" {
		c.JSON(http.StatusOK, gin.H{
			"state":   3000,
			"message": "缺少参数",
		})
		return
	}
	wc, err := services.GetWechatUserOne(map[string]interface{}{"code": code}, "code desc")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"state":   3000,
			"message": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state":   2000,
		"message": "success",
		"data":    wc,
	})
}

func GetUsers(c *gin.Context) {
	users, total, errs := services.GetUserPages(map[string]interface{}{}, "code desc", 0, 20)
	if errs != nil {
		c.JSON(http.StatusOK, gin.H{
			"state":   3000,
			"message": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state":   2000,
		"message": "success",
		"data":    users,
		"total":   total,
	})
}

func Login(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	if username == "" {
		username = c.DefaultPostForm("username", "")
	}
	password := c.DefaultQuery("password", "")
	if password == "" {
		password = c.DefaultPostForm("password", "")
	}
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"state":    3001,
			"message":  "缺失参数",
			"username": username,
			"password": password,
		})
		c.Abort()
		return
	}
	result, err := services.Auth(username, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"state":    4001,
			"message":  err.Error(),
			"username": username,
			"password": password,
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state":   2000,
		"message": "success",
		"data":    result,
	})
}
