package user

import (
	"gin-auth/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	users,total,errs := services.GetUserPages(map[string]interface{}{},"code desc",0,20)
	if errs != nil {
		c.JSON(http.StatusOK,gin.H{
			"state":3000,
			"message":"error",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"state":2000,
		"message":"success",
		"data":users,
		"total":total,
	})
}

func Login(c *gin.Context)  {
	username := c.DefaultQuery("username","")
	if username == ""{
		username = c.DefaultPostForm("username","")
	}
	password := c.DefaultQuery("password","")
	if password ==""{
		password = c.DefaultPostForm("password","")
	}
	if username == "" || password=="" {
		c.JSON(http.StatusOK,gin.H{
			"state":3001,
			"message":"缺失参数",
			"username":username,
			"password":password,
		})
		c.Abort()
		return
	}
	result,err := services.Auth(username,password)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"state":4001,
			"message":err.Error(),
			"username":username,
			"password":password,
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"state":2000,
		"message":"success",
		"data":result,
	})
}


