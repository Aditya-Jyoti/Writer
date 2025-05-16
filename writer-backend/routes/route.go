package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type BlogHandler struct {
	Collection *mongo.Collection
}

func (h *BlogHandler) RegisterBlogRoutes(r *gin.Engine) {
	blogs := r.Group("/blogs")
	{
		blogs.POST("/", h.CreateBlog)
		blogs.GET("/", h.GetAllBlogs)
	}
}
