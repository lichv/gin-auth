package main

import (
	"fmt"
	"gin-auth/utils/jwt"
	"gin-auth/utils/setting"
)

func main() {
	setting.Setup()
	//fmt.Println(time.Now().UnixNano())
	//url := "https://baidu.com/search?key=word"
	//a,_ := utils.URLAppendParams(url,"key","new")
	//b,_ := utils.URLAppendParams(a,"name","lala")
	//fmt.Println(a)
	//fmt.Println(b)
	//now := time.Now().UnixNano()/1000
	//fmt.Println(now)
	//code := "wx_"+strconv.FormatUint(uint64(now),36)+strconv.Itoa(rand.New(rand.NewSource(now)).Intn(90)+10)
	//fmt.Println(strconv.FormatUint(uint64(now),36))
	//fmt.Println(code)
	//fmt.Println(utils.EncodeMD5(setting.AppSetting.SecretSalt+"123456"))
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiIiwidXNlcm5hbWUiOiIiLCJleHAiOjE2MjE5Mzc4ODEsImlzcyI6Imdpbi1ibG9nIn0.Ijc3Iqz1LdokiRZuz9Cu478OCmZ6Thb-qO8r-jO3xMg"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiIiwidXNlcm5hbWUiOiIiLCJleHAiOjE2MjE5MzkxMzcsImlzcyI6Imdpbi1ibG9nIn0.s7qQcQGoxPscMxatJ1FNt4zQXlqos0JikcjadVtNAhQ"
	claims, _ := jwt.ParseToken(token)
	fmt.Println(claims)

	token, _ = jwt.GenerateToken("wx_fyxco2hxts83", "。")
	fmt.Println(token)

}
