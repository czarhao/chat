package main

import (
	"chat/models"
	"chat/routes"
	"chat/system"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	defer totalDefer()
	serverConf := system.ReadServerIni()
	gin.SetMode(serverConf.RunMode)
	server := &http.Server{
		Addr:           serverConf.HttpPort,
		Handler:        routes.InitRouter(),
		ReadTimeout:    serverConf.ReadTimeout,
		WriteTimeout:   serverConf.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("[info] start http server listening %s", serverConf.HttpPort)
	if err := server.ListenAndServe(); err != nil {
		system.Save.ServerPanic("the http service has made some errors : ", err)
	}
}

func totalDefer() {
	system.Save.DeferFile()
	models.DeferDb()
}
