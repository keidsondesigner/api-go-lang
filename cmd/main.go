package main

import (
	"fmt"
	"log"
	"api-go-lang/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()
	defer db.Pool.Close()

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
