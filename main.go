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

	port := config.ENV.PORT
	if port == "" {
		port = "5432" // Default jika tidak ada
	}

	fmt.Println("Starting server on port:", port)
	r.Run(fmt.Sprintf(":%s", port))
}
