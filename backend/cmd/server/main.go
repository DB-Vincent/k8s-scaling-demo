package main

import "github.com/DB-Vincent/k8s-scaling-demo/backend/internal/server"

func main() {
    router := server.SetupRouter()
    router.Run(":8080")
}
