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

	wscPath := server.Group("workspaces")
	wscPath.Use(utils.JwtAuthMiddleware())
	wsc := controllers.NewWorkSpaceController(tc.DeleteWorkspaceTasksHandler())
	wsc.RegisterWorkspaceRoutes(wscPath)

	log.Fatal(server.Run(":8081"))
}
