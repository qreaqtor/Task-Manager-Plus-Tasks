package main

import (
	"log"
	"task-manager-plus-tasks/controllers"
	"task-manager-plus-tasks/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	tcPath := server.Group("tasks")
	tcPath.Use(utils.JwtAuthMiddleware())
	tc := controllers.NewTaskController()
	tc.RegisterTasksRoutes(tcPath)

	log.Fatal(server.Run(":8081"))
}
