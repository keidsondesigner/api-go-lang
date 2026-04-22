package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("Starting server on :8080")
	server.Run(":8080")
}