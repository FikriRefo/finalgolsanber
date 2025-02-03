package router

import (
	"socmed/config"
	"socmed/handler"
	"socmed/middleware"
	"socmed/repository"
	"socmed/service"

	"github.com/gin-gonic/gin"
)

func PostRouter(api *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(config.DB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	r := api.Group("/tweets")

	r.Use(middleware.JWTMiddleware())
	r.POST("/", postHandler.Create)
	r.GET("/", postHandler.GetAllByUserID) // Get all posts by user ID
	r.PUT("/:id", postHandler.Update)      // Update a post
	r.DELETE("/:id", postHandler.Delete)   // Delete a post
}
