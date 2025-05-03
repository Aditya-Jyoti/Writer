package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/aditya-jyoti/writer/models"
)

var blogCollection *mongo.Collection

func RegisterBlogRoutes(r *gin.Engine, collection *mongo.Collection) {
	blogs := r.Group("/blogs")
	{
		blogs.POST("/", CreateBlog)
		blogs.GET("/", GetAllBlogs)
	}

	blogCollection = collection
}

func CreateBlog(ctx *gin.Context) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Content     string `json:"content"`
	}

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog := models.NewBlog(input.Title, input.Description, input.Content)

	_, err := blogCollection.InsertOne(context.Background(), blog)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert blog"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Blog entry created", "id": blog.Id})
}

func GetAllBlogs(ctx *gin.Context) {
	cursor, err := blogCollection.Find(context.Background(), bson.M{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blogs"})
		return
	}
	defer cursor.Close(context.TODO())

	var blogs []models.Blog
	if err := cursor.All(context.TODO(), &blogs); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode blogs"})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}
