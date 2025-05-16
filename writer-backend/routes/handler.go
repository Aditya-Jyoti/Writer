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

// GetBlogById godoc
// @Summary Get a blog post by ID
// @Description Retrieves a blog entry by its ID
// @Tags Blogs
// @Produce json
// @Param id path string true "Blog ID"
// @Success 200 {object} models.Blog "Blog post"
// @Failure 404 {object} object "Error: Blog not found"
// @Router /blogs/{id} [get]
func (h *BlogHandler) GetBlogById(ctx *gin.Context) {
	id := ctx.Param("id")

	var blog models.Blog

	if err := h.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&blog); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

// UpdateBlog godoc
// @Summary Update a blog post
// @Description Updates the title, description, content, and published status of a blog post by its ID
// @Tags Blogs
// @Accept json
// @Produce json
// @Param id path string true "Blog ID"
// @Param blog body models.UpdateBlogInput true "Updated blog fields"
// @Success 200 {object} object "Message: Blog updated successfully"
// @Failure 400 {object} object "Error: Invalid input"
// @Failure 500 {object} object "Error: Failed to update blog"
// @Router /blogs/{id} [put]
func (h *BlogHandler) UpdateBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	var input models.UpdateBlogInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"title":       input.Title,
			"description": input.Description,
			"content":     input.Content,
			"published":   input.Published,
		},
	}

	if _, err := h.Collection.UpdateOne(context.Background(), bson.M{"_id": id}, update); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

// DeleteBlog godoc
// @Summary Delete a blog post
// @Description Deletes a blog post by its ID
// @Tags Blogs
// @Produce json
// @Param id path string true "Blog ID"
// @Success 200 {object} object "Message: Blog deleted successfully"
// @Failure 500 {object} object "Error: Failed to delete blog"
// @Router /blogs/{id} [delete]
func (h *BlogHandler) DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")

	if _, err := h.Collection.DeleteOne(context.Background(), bson.M{"_id": id}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}
