package main

import (
	"fmt"
	"log"

	"api-go-lang/db"
	"api-go-lang/repository"
	"api-go-lang/usecase"
	"api-go-lang/controller"

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

	// Repositories
	ProductRepository := repository.NewProductRepository(dbConnection)
	// Usecases
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	// Controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/product", ProductController.GetProducts)

	log.Println("Starting server on :8080")
	server.Run(":8080")
}
