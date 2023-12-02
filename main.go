package main

import (
	"sebasromero/github.com/rest-api/internal/tasks"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/tasks", tasks.GetTasks)
	router.Run("localhost:9090")
}
