package wechat

import (
	"fmt"
	"gin-auth/app/services"
	"gin-auth/utils"
	"gin-auth/utils/jwt"
	mpoauth2 "github.com/chanxuehong/wechat/mp/oauth2"
	"github.com/chanxuehong/wechat/oauth2"
	"github.com/gin-gonic/gin"
	lichv "github.com/lichv/go"
	"net/http"
	"strings"
)

func Oauth(c *gin.Context)  {
	state := c.DefaultQuery("state","")
	if state == ""{
		state = c.DefaultPostForm("state","")
	}
	if state == "" || !strings.Contains(state,"@"){
		c.JSON(http.StatusOK,gin.H{
			"state":3000,
			"message":"缺少参数",
		})
		return
	}
	split := strings.Split(state, "@")
	if len(split) != 2 {
		c.JSON(http.StatusOK,gin.H{
			"state":3000,
			"message":"参数错误",
		})
		return
	}
	code := split[0]
	configOne, _ := services.GetWechatConfigOne(map[string]interface{}{"code": code}, "code desc")
	fmt.Println(configOne)

	url := mpoauth2.AuthCodeURL(configOne.Appid, configOne.AuthRedirectUrl, configOne.Scope, state)
	fmt.Println(url)
	c.Redirect(302,url)
}

func Callback(c *gin.Context) {
	state := c.DefaultQuery("state","")
	if state == ""{
		state = c.DefaultPostForm("state","")
	}
	code := c.DefaultQuery("code","")
	if code == ""{
		code = c.DefaultPostForm("code","")
	}
	if state == "" || !strings.Contains(state,"@") || code == ""{
		c.JSON(http.StatusOK,gin.H{
			"state":3000,
			"message":"缺少参数",
		})
		return
	}
	split := strings.Split(state, "@")
	if len(split) != 2 {
		c.JSON(http.StatusOK,gin.H{
			"state":3000,
			"message":"参数错误",
		})
		return
	}
	configCode := split[0]
	redirectUrl := split[1]
	configOne, _ := services.GetWechatConfigOne(map[string]interface{}{"code": configCode}, "code desc")
	fmt.Println(configOne)

	endpoint := mpoauth2.NewEndpoint(configOne.Appid, configOne.Appsecret)

	oauth2Client := oauth2.Client{
		Endpoint: endpoint,
	}
	atoken, err := oauth2Client.ExchangeToken(code)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"state":3002,
			"message":"获取token失败",
		})
		return
	}
	userinfo, err := mpoauth2.GetUserInfo(atoken.AccessToken, atoken.OpenId, "", nil)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"state":3002,
			"message":"获取用户信息失败",
		})
		return
	}
	fmt.Println(userinfo)
	sex := "unknow"
	if userinfo.Sex == 1{
		sex = "male"
	}else if userinfo.Sex == 2{
		sex = "female"
	}

	wu,err := services.AddWechatUser(map[string]interface{}{"config_code": utils.Strval(configCode), "openid": userinfo.OpenId, "unionid": userinfo.UnionId, "nickname": userinfo.Nickname, "sex": sex, "country": userinfo.Country, "province": userinfo.Province, "city": userinfo.City, "headimage": userinfo.HeadImageURL})
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"state":3002,
			"message":"添加用户信息失败",
		})
		return
	}
	fmt.Println(wu)

	token_validity_time, _ := services.FindSysparamValueByCode("token_validity_time")
	token, err := jwt.GenerateToken(wu.Code, wu.Nickname,lichv.IntVal(token_validity_time))
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"state":3002,
			"message":"生成Token失败",
		})
		return
	}
	fmt.Println(token)
	uri,err := utils.URLAppendParams(redirectUrl,"user_token",token)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"state":3002,
			"message":"生成Token失败",
		})
		return
	}
	fmt.Println(uri)
	c.Redirect(302,uri)

}