package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func start() {
	engine := gin.Default()
	engine.POST("/echo", func(c *gin.Context) {
		bytes, _ := ioutil.ReadAll(c.Request.Body)
		c.JSON(http.StatusOK, gin.H{
			"body": string(bytes),
		})
	})
	engine.Run()
}
