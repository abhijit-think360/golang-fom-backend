package main

import "github.com/gin-gonic/gin"

func hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "Hello World"})
}

func main() {
	router := gin.Default()
	router.GET("/hello", hello)
	router.Run(":3000")
}
