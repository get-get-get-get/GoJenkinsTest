package main

import "github.com/gin-gonic/gin"

var router *gin.Engine

func init() {
	router = gin.Default()
}

func setRoutes() {
	router.GET("/add/:strX/:strY", AddNumbers)
	router.GET("/subtract/:strX/:strY", SubtractNumbers)
}
