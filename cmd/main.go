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

	server := gin.Default()

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Database connection
	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	defer dbConnection.Close()

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
