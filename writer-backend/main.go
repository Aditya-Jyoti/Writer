package main

import (
	"log"
	"os"

	"github.com/aditya-jyoti/writer/lib"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	uri_string := os.Getenv("MONGO_URI")
	port := "3000"
	db_name := "writer"

	router := gin.Default()

	lib.ConnectDB(uri_string)
	defer lib.DisconnectDB()

	db := lib.GetDB().Database(db_name)
	log.Println("Ready to use DB:", db.Name())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":" + port)
}
