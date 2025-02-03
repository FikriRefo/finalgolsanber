package router

import (
	"socmed/config"
	"socmed/handler"
	"socmed/middleware"
	"socmed/repository"
	"socmed/service"

	"github.com/gin-gonic/gin"
)

func ProfileRouter(api *gin.RouterGroup) {
	profileRepository := repository.NewProfileRepository(config.DB)
	profileService := service.NewProfileService(profileRepository)
	profileHandler := handler.NewProfileHandler(profileService)

	r := api.Group("/profiles")
	r.Use(middleware.JWTMiddleware())
	r.POST("/", profileHandler.Create)
	r.GET("/", profileHandler.GetByUserID)
	r.PUT("/:id", profileHandler.Update)
	r.DELETE("/:id", profileHandler.Delete)
}
