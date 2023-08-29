package main

import (
	"log"
	"task-manager-plus-tasks/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	tcPath := server.Group("tasks")
	tc := controllers.NewTaskController()
	tc.RegisterTasksRoutes(tcPath)

	log.Fatal(server.Run(":8080"))
}
