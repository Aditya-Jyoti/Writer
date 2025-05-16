package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/aditya-jyoti/writer/lib"
)

type BlogHandler struct {
	Collection *mongo.Collection
}

func (h *BlogHandler) RegisterBlogRoutes(r *gin.Engine) {
	blogs := r.Group("/blogs")
	{
		blogs.GET("/", h.GetAllBlogs)

		blogs.Use(lib.JWTAuthMiddleware())
		{
			blogs.POST("/", h.CreateBlog)
			blogs.GET("/:id", h.GetBlogById)
			blogs.PUT("/:id", h.UpdateBlog)
			blogs.DELETE("/:id", h.DeleteBlog)
		}
	}
}
