package wechat

import (
	"gin-auth/app/services"
	"gin-auth/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetConfig(c *gin.Context) {
	code := c.DefaultQuery("code","")
	if code == ""{
		code = c.DefaultPostForm("code","")
	}
	m,err := utils.GetMapFromJson(c)
	if err == nil{
		codeTmp,ok := m["code"]
		if ok {
			code = utils.Strval(codeTmp)
		}
	}
	if code == "" {
		c.JSON(http.StatusOK,gin.H{
			"state":3000,
			"message":"缺少参数",
		})
		return
	}
	wc, err := services.GetWechatConfigOne(map[string]interface{}{"code": code}, "code desc")
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"state":3000,
			"message":"error",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"state":2000,
		"message":"success",
		"data":wc,
	})
}
