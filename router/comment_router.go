package router

import (
	"socmed/config"
	"socmed/handler"
	"socmed/middleware"
	"socmed/repository"
	"socmed/service"

	"github.com/gin-gonic/gin"
)

func CommentRouter(api *gin.RouterGroup) {
	commentRepository := repository.NewCommentRepository(config.DB)
	commentService := service.NewCommentService(commentRepository)
	commentHandler := handler.NewCommentHandler(commentService)

	r := api.Group("/comments")
	r.Use(middleware.JWTMiddleware()) // Middleware to authenticate

	r.POST("/", commentHandler.Create)                // Create a comment
	r.GET("/:post_id", commentHandler.GetAllByPostID) // Get comments by post ID
	r.PUT("/:id", commentHandler.Update)              // Update a comment
	r.DELETE("/:id", commentHandler.Delete)           // Delete a comment
}
