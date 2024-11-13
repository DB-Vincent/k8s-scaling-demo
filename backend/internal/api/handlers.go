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
		replicas, err := k8s.GetRelatedPods(config)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"replicas": replicas,
		})
	}
}
