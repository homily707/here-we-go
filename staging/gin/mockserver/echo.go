package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func start() {
	engine := gin.Default()
	engine.POST("/echo", func(context *gin.Context) {
		bytes, _ := ioutil.ReadAll(context.Request.Body)
		log.Println(string(bytes))
		context.JSON(http.StatusOK, gin.H{
			"body": string(bytes),
		})
	})
	engine.GET("/echo", func(context *gin.Context) {
		log.Println(string(context.Request.RequestURI))
		context.JSON(http.StatusOK, gin.H{
			"URI": context.Request.RequestURI,
		})
	})
	engine.Run()
}
