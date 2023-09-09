package main

import (
	"log"
	"task-manager-plus-tasks/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	tcPath := server.Group("users/tasks")
	tc := controllers.NewTaskController()
	tc.RegisterTasksRoutes(tcPath)

	wscPath := server.Group("users/workspaces")
	wsc := controllers.NewWorkSpaceController(tc.DeleteTasksHandler())
	wsc.RegisterWorkspaceRoutes(wscPath)

	log.Fatal(server.Run(":8081"))
}
