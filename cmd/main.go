package main

import (
	"GoStream/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    server.RegisterRoutes(router)
    router.Run(":8080")
}
