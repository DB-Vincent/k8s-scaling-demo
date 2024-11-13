package api

import (
	"mime"
	"os"
	"path/filepath"

	"github.com/DB-Vincent/k8s-scaling-demo/backend/internal/k8s"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	config := k8s.SetupConfig()

	// Get the path to the frontend files
	frontendPath := "../frontend/dist/frontend/browser"
	if os.Getenv("GIN_MODE") == "release" {
		frontendPath = "./frontend/dist/frontend/browser"
	}

	// Ensure the frontend path exists
	if _, err := os.Stat(frontendPath); os.IsNotExist(err) {
		panic("Frontend path not found: " + frontendPath)
	}

	if os.Getenv("GIN_MODE") != "release" {
		corsConfig := cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			AllowCredentials: true,
		}
		router.Use(cors.New(corsConfig))
	}

	// API routes
	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/ping", PingHandler)
		apiGroup.GET("/pods", GetPodsHandler(config))
	}

	// Static files for frontend
	router.Static("/static", frontendPath)
	router.NoRoute(func(c *gin.Context) {
		indexPath := filepath.Join(frontendPath, "index.html")
		c.File(indexPath)
	})

	// Quick fix for javascript files
	mime.AddExtensionType(".js", "application/javascript")

	return router
}
