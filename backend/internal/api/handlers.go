package api

import (
	"net/http"

	"github.com/DB-Vincent/k8s-scaling-demo/backend/internal/k8s"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/rest"
)

func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetPodsHandler(config *rest.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"replicas": k8s.GetRelatedPods(config),
		})
	}
}
