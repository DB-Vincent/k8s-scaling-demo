package main

import "github.com/DB-Vincent/k8s-scaling-demo/backend/internal/api"

func main() {
    router := api.SetupRouter()
    router.Run(":8080")
}
