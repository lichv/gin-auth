package routers

import (
	"gin-auth/app/controllers/user"
	"gin-auth/app/controllers/wechat"
	"gin-auth/app/middlewares"
	"gin-auth/utils/setting"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
	"path"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(favicon.New(path.Join(setting.AppSetting.PublicPath, "favicon.ico")))

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"state":   2000,
			"message": "success",
		})
	})

	r.Any("/auth/login", user.Login)
	r.GET("/wechat/oauth",wechat.Oauth)
	r.Any("/wechat/callback",wechat.Callback)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middlewares.JWT())
	{
		apiv1.GET("/user/list", user.GetUsers)
		apiv1.Any("/wechat/config",wechat.GetConfig)
		apiv1.Any("/wechat/user",wechat.GetWechatUser)
	}

	return r
}
