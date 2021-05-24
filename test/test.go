package main

import (
	"fmt"
	"gin-auth/utils"
)

func main() {
	//setting.Setup()
	//fmt.Println(time.Now().UnixNano())
	url := "https://baidu.com/search?key=word"
	a,_ := utils.URLAppendParams(url,"key","new")
	b,_ := utils.URLAppendParams(a,"name","lala")
	fmt.Println(a)
	fmt.Println(b)
	//now := time.Now().UnixNano()/1000
	//fmt.Println(now)
	//code := "wx_"+strconv.FormatUint(uint64(now),36)+strconv.Itoa(rand.New(rand.NewSource(now)).Intn(90)+10)
	//fmt.Println(strconv.FormatUint(uint64(now),36))
	//fmt.Println(code)
	//fmt.Println(utils.EncodeMD5(setting.AppSetting.SecretSalt+"123456"))
}
