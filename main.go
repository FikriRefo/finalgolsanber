package main

import (
	"fmt"
	"socmed/config"
	"socmed/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()
	api := r.Group("/api")

	router.AuthRouter(api)
	router.PostRouter(api)
	router.ProfileRouter(api)
	router.CommentRouter(api)
	// Pastikan server berjalan di port yang benar
	port := config.ENV.PORT
	if port == "" {
		port = "8080" // Default port jika tidak ada di .env
	}
	r.Run(fmt.Sprintf(":%s", port))
}
