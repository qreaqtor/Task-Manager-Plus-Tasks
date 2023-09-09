package controllers

import (
	"net/http"
	"task-manager-plus-tasks/models"
	"task-manager-plus-tasks/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	taskService services.TaskService
}

func NewTaskController() TaskController {
	return TaskController{
		taskService: services.NewTaskService(),
	}
}

func (tc *TaskController) createTask(ctx *gin.Context) {
	var taskIn models.TaskInput
	if err := ctx.ShouldBindJSON(&taskIn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	username := ctx.Param("username")
	taskIn.InitTask(username)
	err := tc.taskService.CreateTask(taskIn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (tc *TaskController) getWorkspaceTasks(ctx *gin.Context) {
	workspaceId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	tasks, err := tc.taskService.GetWorkspaceTasks(workspaceId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (tc *TaskController) deleteUserTask(ctx *gin.Context) {
	taskId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = tc.taskService.DeleteTask(taskId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (tc *TaskController) RegisterTasksRoutes(rg *gin.RouterGroup) {
	rg.POST("/create", tc.createTask)
	rg.GET("/get/:id", tc.getWorkspaceTasks)
	rg.DELETE("/delete/:id", tc.deleteUserTask)
}

func (tc *TaskController) DeleteTasksHandler() func(workspaceId primitive.ObjectID) error {
	return tc.taskService.DeleteWorkspaceTasks
}
