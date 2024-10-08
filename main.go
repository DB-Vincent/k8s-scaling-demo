package main

import (
	"context"
	"net/http"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"mime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	config := setupConfig()

	env := os.Getenv("GIN_MODE")

	if env != "release" {
		// Custom CORS configuration for development
		corsConfig := cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			AllowCredentials: true,
		}

		router.Use(cors.New(corsConfig))
	}

	api := router.Group("/api")

	// Ping endpoint
	api.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	// Pods endpoint
	api.GET("/pods", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"replicas": getRelatedPods(config),
		})
	})

	// Route for static angular content
	router.Static("/static", "./frontend/dist/frontend/browser")
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/frontend/browser/index.html")
	})

	// Quick fix for javascript files
	mime.AddExtensionType(".js", "application/javascript")

	router.Use(cors.Default())

	return router
}

func setupConfig() *rest.Config {
	var config *rest.Config
	var err error

	if _, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount/token"); err == nil {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	}

	if err != nil {
		panic(err.Error())
	}

	return config
}

type Replica struct {
	Name      string `json:"name"`
	Current   bool   `json:"current"`
	NodeName  string `json:"nodeName"`
	Status    string `json:"status"`
	StartTime string `json:"startTime"`
}

func getRelatedPods(config *rest.Config) []Replica {
	var pods []Replica

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	podNamespace := os.Getenv("POD_NAMESPACE")
	podName := os.Getenv("POD_NAME")

	pod, err := clientset.CoreV1().Pods(podNamespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, ownerRef := range pod.OwnerReferences {
		if *ownerRef.Controller {
			labelSelector := labels.Set(pod.Labels).AsSelector().String()
			podList, err := clientset.CoreV1().Pods(podNamespace).List(context.TODO(), metav1.ListOptions{
				LabelSelector: labelSelector,
			})
			if err != nil {
				panic(err.Error())
			}

			for _, p := range podList.Items {
				replica := Replica{}

				replica.Name = p.Name
				replica.NodeName = p.Spec.NodeName
				replica.Status = string(p.Status.Phase)
				replica.StartTime = p.Status.StartTime.String()

				if p.Name == podName {
					replica.Current = true
				} else {
					replica.Current = false
				}

				pods = append(pods, replica)
			}
		}
	}

	return pods
}

func main() {
	router := setupRouter()

	// Start Gin
	router.Run(":8080")
}
