package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	_ "github.com/aditya-jyoti/writer/docs"
	"github.com/aditya-jyoti/writer/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Writer API
// @version 1.0
// @description Simple CRUD API for a blog.

// @contact.name Aditya Jyoti
// @contact.url https://github.com/aditya-jyoti
// @contact.email reach@adityajyoti.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide your JWT token like: "Bearer {token}"

// @host localhost:3000
// @BasePath /
func main() {
	client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var blogCollection *mongo.Collection = client.Database("writer").Collection("blogs")

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	blogHandler := &routes.BlogHandler{Collection: blogCollection}
	blogHandler.RegisterBlogRoutes(router)

	port := "3000"
	router.Run(":" + port)
}
