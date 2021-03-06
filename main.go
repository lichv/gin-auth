package main

import (
	"fmt"
	"gin-auth/app/crontabs"
	"gin-auth/app/models"
	"gin-auth/routers"
	"gin-auth/utils/setting"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	models.Setup()
	crontabs.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}

