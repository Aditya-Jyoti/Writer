package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/aditya-jyoti/writer/models"
)

// CreateBlog godoc
// @Summary Create a new blog post
// @Description Creates a blog with title, description, and content. Default `published` to false and `update_date` to current time.
// @Tags Blogs
// @Accept json
// @Produce json
// @Param blog body models.CreateBlogInput true "Blog input"
// @Success 201 {object} models.SuccessfulCreationResponse "Returns created blog ID"
// @Failure 400 {object} object "Error: Invalid input"
// @Failure 500 {object} object "Error: Failed to insert blog"
// @Router /blogs/ [post]
func (h *BlogHandler) CreateBlog(ctx *gin.Context) {
	var input models.CreateBlogInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog := models.NewBlog(input.Title, input.Description, input.Content)

	if _, err := h.Collection.InsertOne(context.Background(), blog); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert blog"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": blog.Id,
	})
}

// GetAllBlogs godoc
// @Summary Get all blog posts
// @Description Retrieves all blog entries from the database
// @Tags Blogs
// @Produce json
// @Success 200 {array} models.Blog "List of all blog posts"
// @Failure 500 {object} object "Error: Failed to retrieve or decode blogs"
// @Router /blogs/ [get]
func (h *BlogHandler) GetAllBlogs(ctx *gin.Context) {
	cursor, err := h.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blogs"})
		return
	}
	defer cursor.Close(context.Background())

	var blogs []models.Blog
	if err := cursor.All(context.Background(), &blogs); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode blogs"})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}
